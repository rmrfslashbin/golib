// Package golib is a collection of commonly used functions.
package golib

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"math"
	"os"
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

// CheckSha256 check the sha-256 sum of a file against the expected value.
func CheckSha256(filename string, sha256sum string) (bool, error) {
	// Open the file
	f, err := os.Open(filename)
	if err != nil {
		return false, err
	}
	defer f.Close()

	// New hash object
	hash := sha256.New()

	// Digest the file
	if _, err := io.Copy(hash, f); err != nil {
		return false, err
	}

	// Compare the hashes
	if ok := sha256sum == hex.EncodeToString(hash.Sum(nil)); ok {
		// Match!
		return true, nil
	}
	// Fail!
	return false, nil
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

// RandString generages a random string of n chars
func RandString(n int) string {
	// Derived from https://stackoverflow.com/a/55860599
	buff := make([]byte, int(math.Ceil(float64(n)/2)))
	rand.Read(buff)
	str := hex.EncodeToString(buff)
	return str[:n]
}
