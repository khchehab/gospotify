package gospotify

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"sync"
	"time"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
	"golang.org/x/oauth2/spotify"
)

// authResult is the result of the authorization code flow.
type authResult struct {
	code string
	err  error
}

// authCodeURLFunc is a function that returns the authorization code URL.
type authCodeURLFunc func(state string) string

// exchangeFunc is a function that exchanges an authorization code for an access token.
type exchangeFunc func(ctx context.Context, code string) (*oauth2.Token, error)

// AuthorizationCode performs the authorization code flow and returns the initial token.
func AuthorizationCode(clientID, clientSecret, redirectURL string, scopes []string, serverPort int) (oauth2.TokenSource, error) {
	oauthConfig := oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  redirectURL,
		Scopes:       scopes,
		Endpoint:     spotify.Endpoint,
	}

	return authorizationCode(&oauthConfig, serverPort, func(state string) string {
		return oauthConfig.AuthCodeURL(state)
	}, func(ctx context.Context, code string) (*oauth2.Token, error) {
		return oauthConfig.Exchange(ctx, code)
	})
}

// AuthorizationCodeWithPKCE performs the authorization code flow with PKCE and returns the initial token.
func AuthorizationCodeWithPKCE(clientID, redirectURL string, scopes []string, serverPort int) (oauth2.TokenSource, error) {
	verifier := oauth2.GenerateVerifier()
	challenge := oauth2.S256ChallengeOption(verifier)

	oauthConfig := oauth2.Config{
		ClientID:    clientID,
		RedirectURL: redirectURL,
		Scopes:      scopes,
		Endpoint:    spotify.Endpoint,
	}

	return authorizationCode(&oauthConfig, serverPort, func(state string) string {
		return oauthConfig.AuthCodeURL(state, challenge)
	}, func(ctx context.Context, code string) (*oauth2.Token, error) {
		return oauthConfig.Exchange(ctx, code, oauth2.VerifierOption(verifier))
	})
}

// ClientCredentials performs the client credentials flow and returns the initial token.
func ClientCredentials(clientID, clientSecret string, scopes []string) oauth2.TokenSource {
	oauthConfig := clientcredentials.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Scopes:       scopes,
		TokenURL:     spotify.Endpoint.TokenURL,
	}

	return oauthConfig.TokenSource(context.Background())
}

// authorizationCode performs the authorization code flow and returns the initial token.
// This function is used internally by both AuthorizationCode and AuthorizationCodeWithPKCE.
func authorizationCode(oauthConfig *oauth2.Config, serverPort int, authCodeURL authCodeURLFunc, exchange exchangeFunc) (oauth2.TokenSource, error) {
	state := RandomString(16)
	authCh := make(chan authResult, 1)

	server := startLocalAuthServer(serverPort, authCh, state)
	defer shutdownServer(server)

	url := authCodeURL(state)
	if err := OpenBrowser(url); err != nil {
		fmt.Println("error opening browser: ", err) // TODO change this into a logger.Error or logger.Debug after implementing the logger
		fmt.Println("Open this url in your browser:", url)
	}

	code, err := waitForAuthResult(authCh, 2*time.Minute)
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
func startLocalAuthServer(serverPort int, authCh chan authResult, state string) *http.Server {
	var once sync.Once

	mux := http.NewServeMux()
	server := &http.Server{
		Addr:    fmt.Sprintf("127.0.0.1:%d", serverPort),
		Handler: mux,
	}
	mux.HandleFunc("/callback", newCallbackHandler(state, authCh, &once))
	go func(srv *http.Server) {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			fmt.Println("error starting server:", err)
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
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	token, err := exchange(ctx, code)
	if err != nil {
		return nil, fmt.Errorf("error exchanging code for token: %w", err)
	}

	tokenSource := oauthConfig.TokenSource(context.Background(), token)

	return tokenSource, nil
}

// shutdownServer shuts down the local HTTP server.
func shutdownServer(server *http.Server) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		fmt.Println("error shutting down server: ", err)
	}

}
