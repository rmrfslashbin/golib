package golib

import (
	"path"
	"testing"
)

// TestBrowserOpen tests the golib.BrowserOpen function.
func TestBrowserOpen(t *testing.T) {
	if err := BrowserOpen("https://github.com/rmrfslashbin/golib"); err != nil {
		t.Error("BrowserOpen returned error: " + err.Error())
	}
}

// TestCheckSha256 checks the golib.CheckSha256 function.
func TestCheckSha256(t *testing.T) {
	expectedHash := "aefdc17c3a9e2cadf498ab1efbec5b1da6fe88a323a74d46e0bd2e2502d9a952"
	if ok, err := CheckSha256(path.Clean("sha_256_test_file"), expectedHash); err != nil {
		t.Error("BrowserOpen returned error: " + err.Error())
	} else if !ok {
		t.Error("CheckSha256 returned false for a valid hash")
	}
}

// TestNumberFormat calls golib.TestNumberFormat function.
func TestNumberFormat(t *testing.T) {
	number := NumberFormat(1234567890)
	if number != "1.23B" {
		t.Error("NumberFormat returned an invalid string: " + number)
	}
}

// TestRandString calls golib.TestRandString function.
func TestRandString(t *testing.T) {
	randString := RandString(10)
	if len(randString) != 10 || randString == "" {
		t.Error("RandString returned an invalid string")
	}
}
