// github.com/rmrfslashbin/golib is a collection of commonly used functions.
package golib

import (
	"crypto/rand"
	"encoding/hex"
	"math"
	"os/exec"
	"runtime"
)

// BrowserOpen opens a local browser.
func BrowserOpen(url string) error {
	// From: https://gist.github.com/hyg/9c4afcd91fe24316cbf0
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "windows":
		cmd = "cmd"
		args = []string{"/c", "start"}
	case "darwin":
		cmd = "open"
	default: // "linux", "freebsd", "openbsd", "netbsd"
		cmd = "xdg-open"
	}
	args = append(args, url)
	return exec.Command(cmd, args...).Start()
}

// RandString generages a random string of n chars
func RandString(n int) string {
	// Derived from https://stackoverflow.com/a/55860599
	buff := make([]byte, int(math.Ceil(float64(n)/2)))
	rand.Read(buff)
	str := hex.EncodeToString(buff)
	return str[:n]
}
