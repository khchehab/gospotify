package gospotify

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"sync"
	"time"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
	"golang.org/x/oauth2/spotify"
)

// Authenticator represents an authenticator object containing the OAuth2 credentials.
type Authenticator struct {
	clientID     string
	clientSecret string
	redirectURL  string
	serverPort   int

	logger *slog.Logger
}

// NewAuthenticator creates a new authenticator object.
func NewAuthenticator(clientID, clientSecret, redirectURL string, logger *slog.Logger) *Authenticator {
	return NewAuthenticatorWithServerPort(clientID, clientSecret, redirectURL, 8080, logger)
}

// NewAuthenticatorWithServerPort creates a new authenticator object with a specified server port.
func NewAuthenticatorWithServerPort(clientID, clientSecret, redirectURL string, serverPort int, logger *slog.Logger) *Authenticator {
	if logger == nil {
		logger = slog.New(slog.NewTextHandler(io.Discard, nil))
	}
	return &Authenticator{
		clientID:     clientID,
		clientSecret: clientSecret,
		redirectURL:  redirectURL,
		serverPort:   serverPort,
		logger:       logger,
	}
}

// AuthorizationCode performs the authorization code flow and returns the initial token.
func (a *Authenticator) AuthorizationCode(scopes []Scope) (oauth2.TokenSource, error) {
	if err := requireNonEmpty("client id", a.clientID); err != nil {
		return nil, err
	}
	if err := requireNonEmpty("client secret", a.clientSecret); err != nil {
		return nil, err
	}
	if err := requireNonEmpty("redirect URL", a.redirectURL); err != nil {
		return nil, err
	}

	oauthConfig := oauth2.Config{
		ClientID:     a.clientID,
		ClientSecret: a.clientSecret,
		RedirectURL:  a.redirectURL,
		Scopes:       scopesToString(scopes),
		Endpoint:     spotify.Endpoint,
	}

	return a.authorizationCode(&oauthConfig, func(state string) string {
		return oauthConfig.AuthCodeURL(state)
	}, func(ctx context.Context, code string) (*oauth2.Token, error) {
		return oauthConfig.Exchange(ctx, code)
	})
}

// AuthorizationCodeWithPKCE performs the authorization code flow with PKCE and returns the initial token.
func (a *Authenticator) AuthorizationCodeWithPKCE(scopes []Scope) (oauth2.TokenSource, error) {
	if err := requireNonEmpty("client id", a.clientID); err != nil {
		return nil, err
	}
	if err := requireNonEmpty("redirect URL", a.redirectURL); err != nil {
		return nil, err
	}

	verifier := oauth2.GenerateVerifier()
	challenge := oauth2.S256ChallengeOption(verifier)

	oauthConfig := oauth2.Config{
		ClientID:    a.clientID,
		RedirectURL: a.redirectURL,
		Scopes:      scopesToString(scopes),
		Endpoint:    spotify.Endpoint,
	}

	return a.authorizationCode(&oauthConfig, func(state string) string {
		return oauthConfig.AuthCodeURL(state, challenge)
	}, func(ctx context.Context, code string) (*oauth2.Token, error) {
		return oauthConfig.Exchange(ctx, code, oauth2.VerifierOption(verifier))
	})
}

// ClientCredentials performs the client credentials flow and returns the initial token.
func (a *Authenticator) ClientCredentials(scopes []Scope) (oauth2.TokenSource, error) {
	if err := requireNonEmpty("client id", a.clientID); err != nil {
		return nil, err
	}
	if err := requireNonEmpty("client secret", a.clientSecret); err != nil {
		return nil, err
	}

	oauthConfig := clientcredentials.Config{
		ClientID:     a.clientID,
		ClientSecret: a.clientSecret,
		Scopes:       scopesToString(scopes),
		TokenURL:     spotify.Endpoint.TokenURL,
	}

	return oauthConfig.TokenSource(context.Background()), nil
}

// authResult is the result of the authorization code flow.
type authResult struct {
	code string
	err  error
}

// authCodeURLFunc is a function that returns the authorization code URL.
type authCodeURLFunc func(state string) string

// exchangeFunc is a function that exchanges an authorization code for an access token.
type exchangeFunc func(ctx context.Context, code string) (*oauth2.Token, error)

// authorizationCode performs the authorization code flow and returns the initial token.
// This function is used internally by both [AuthorizationCode] and [AuthorizationCodeWithPKCE].
func (a *Authenticator) authorizationCode(oauthConfig *oauth2.Config, authCodeURL authCodeURLFunc, exchange exchangeFunc) (oauth2.TokenSource, error) {
	state, err := RandomString(authStateLength)
	if err != nil {
		return nil, err
	}

	authCh := make(chan authResult, 1)

	server := a.startLocalAuthServer(authCh, state)
	defer a.shutdownServer(server)

	url := authCodeURL(state)
	if err := OpenBrowser(url); err != nil {
		a.logger.Error("error opening browser", "error", err)
		a.logger.Info("Open this URL in your browser", "url", url)
	}

	code, err := waitForAuthResult(authCh, authWaitTimeout)
	if err != nil {
		return nil, err
	}

	return exchangeCode(oauthConfig, code, exchange)
}

// newCallbackHandler returns the HTTP handler for the OAuth2 callback endpoint.
// It sends exactly one authResult to authCh (guarded by once) and then returns.
func newCallbackHandler(state string, authCh chan authResult, once *sync.Once) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res := authResult{code: "", err: errors.New("unknown error")}

		defer once.Do(func() {
			authCh <- res
		})

		if r.URL.Query().Get("state") != state {
			res.err = errors.New("invalid state parameter")
			http.Error(w, "invalid state parameter", http.StatusBadRequest)
			return
		}

		if errorCode := r.URL.Query().Get("error"); errorCode != "" {
			errMsg := fmt.Sprintf("oauth error: %s", errorCode)
			res.err = errors.New(errMsg)
			http.Error(w, errMsg, http.StatusForbidden)
			return
		}

		code := r.URL.Query().Get("code")
		if code == "" {
			res.err = errors.New("code not found in query parameters")
			http.Error(w, "code not found in query parameters", http.StatusBadRequest)
			return
		}

		_, _ = fmt.Fprintf(w, "Authorization successful! You can close this window")

		res.code = code
		res.err = nil
	}
}

// startLocalAuthServer starts a local HTTP server that listens for the authorization code.
func (a *Authenticator) startLocalAuthServer(authCh chan authResult, state string) *http.Server {
	var once sync.Once

	mux := http.NewServeMux()
	server := &http.Server{
		Addr:    fmt.Sprintf("127.0.0.1:%d", a.serverPort),
		Handler: mux,
	}
	mux.HandleFunc("/callback", newCallbackHandler(state, authCh, &once))
	go func(srv *http.Server) {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			a.logger.Error("error starting http server", "error", err)
		}
	}(server)

	return server
}

// waitForAuthResult waits for the authorization code to be received.
func waitForAuthResult(authCh chan authResult, timeout time.Duration) (string, error) {
	timer := time.NewTimer(timeout)
	defer timer.Stop()

	select {
	case res := <-authCh:
		if res.err != nil {
			return "", res.err
		}

		return res.code, nil
	case <-timer.C:
		return "", errors.New("authorization timed out")
	}
}

// exchangeCode exchanges the authorization code for an access token.
func exchangeCode(oauthConfig *oauth2.Config, code string, exchange exchangeFunc) (oauth2.TokenSource, error) {
	ctx, cancel := context.WithTimeout(context.Background(), authCodeExchangeTimeout)
	defer cancel()

	token, err := exchange(ctx, code)
	if err != nil {
		return nil, fmt.Errorf("error exchanging code for token: %w", err)
	}

	tokenSource := oauthConfig.TokenSource(context.Background(), token)

	return tokenSource, nil
}

// shutdownServer shuts down the local HTTP server.
func (a *Authenticator) shutdownServer(server *http.Server) {
	ctx, cancel := context.WithTimeout(context.Background(), authServerShutdownTimeout)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		a.logger.Error("error shutting down server", "error", err)
	}

}

// scopesToStrings converts an array of [Scope] to an array of strings.
func scopesToString(scopes []Scope) []string {
	result := make([]string, len(scopes))
	for i, scope := range scopes {
		result[i] = string(scope)
	}
	return result
}
