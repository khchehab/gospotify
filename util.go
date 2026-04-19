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
