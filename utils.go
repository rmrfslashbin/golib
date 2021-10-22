// Package golib is a collection of commonly used functions.
package golib

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
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

// NumberFormat formats a number with magnitude.
func NumberFormat(num float64) string {
	// Inspired by https://stackoverflow.com/a/66105942
	units := [5]string{"k", "M", "B", "T", "Q"}

	unit := len(fmt.Sprintf("%.0f", math.Floor(num/1.0e+1)))

	x := math.Abs(num) / math.Pow(10, float64(unit-(unit%3)))
	i := int(math.Floor(float64(unit)/3.0) - 1.0)

	if i < 0 || i > len(units)-1 {
		return fmt.Sprintf("%.0f", num)
	}
	return fmt.Sprintf("%.2f%s", x, units[i])
}
