package gospotify

import (
	"fmt"
	"math/rand/v2"
	"os/exec"
	"runtime"
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

// OpenBrowser opens the given URL in the user's default browser.'
func OpenBrowser(url string) error {
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "windows":
		cmd = "rundll32"
		args = []string{"url.dll,FileProtocolHandler", url}
	case "darwin":
		cmd = "open"
		args = []string{url}
	case "linux":
		cmd = "xdg-open"
		args = []string{url}
	default:
		return fmt.Errorf("unsupported platform: %s", runtime.GOOS)
	}

	return exec.Command(cmd, args...).Start()
}
