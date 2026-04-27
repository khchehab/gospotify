package gospotify

import (
	"fmt"
	"math/rand/v2"
	"net/url"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"time"
)

// allowedChars is the set of characters allowed in a random string.
const allowedChars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// r is a random number generator.
var r = rand.New(rand.NewPCG(uint64(time.Now().UnixMicro()), uint64(time.Now().UnixMicro())))

// RandomString returns a random string of the given length.
func RandomString(len int) string {
	b := make([]byte, len)
	for i := range b {
		b[i] = allowedChars[r.IntN(62)]
	}
	return string(b)
}

// browserCommand returns the command and arguments needed to open a URL on the given OS.
func browserCommand(goos, url string) (cmd string, args []string, err error) {
	switch goos {
	case "windows":
		return "rundll32", []string{"url.dll,FileProtocolHandler", url}, nil
	case "darwin":
		return "open", []string{url}, nil
	case "linux":
		return "xdg-open", []string{url}, nil
	default:
		return "", nil, fmt.Errorf("unsupported platform: %s", goos)
	}
}

// OpenBrowser opens the given URL in the user's default browser.
func OpenBrowser(url string) error {
	cmd, args, err := browserCommand(runtime.GOOS, url)
	if err != nil {
		return err
	}
	return exec.Command(cmd, args...).Start()
}

// requireNonEmpty makes sure a value is not empty, and if its empty, return an error.
func requireNonEmpty(field, value string) error {
	if value == "" {
		return fmt.Errorf("%s cannot be empty", field)
	}
	return nil
}

// requireNonEmptyArray makes sure an array is not empty, and if its empty, return an error.
func requireNonEmptyArray[T any](field string, value []T) error {
	if len(value) == 0 {
		return fmt.Errorf("%s cannot be empty", field)
	}
	return nil
}

// requiredQueryParam is a structure for holding the key/value pair for a required query parameter.
type requiredQueryParam struct {
	// key is the query parameter key.
	key string
	// value is the query parameter value.
	value any
}

// String is the string representation of a required query parameter.
func (q requiredQueryParam) String() string {
	var value string
	switch t := q.value.(type) {
	case string:
		value = t
	case []string:
		value = strings.Join(t, ",")
	case bool:
		value = strconv.FormatBool(t)
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		value = fmt.Sprintf("%d", t)
	default:
		value = fmt.Sprintf("%v", q.value)
	}

	return fmt.Sprintf("%s=%s", q.key, url.QueryEscape(value))
}

// appendQueryParams appends the given query parameters to the endpoint. This is used for required query parameters.
// If params is empty string will be returned.
func appendQueryParams(endpoint string, params ...requiredQueryParam) string {
	if len(params) == 0 {
		return endpoint
	}
	sb := strings.Builder{}
	sb.WriteString(endpoint)
	sb.WriteByte('?')
	for i, p := range params {
		if i > 0 {
			sb.WriteByte('&')
		}
		sb.WriteString(p.String())
	}
	return sb.String()
}
