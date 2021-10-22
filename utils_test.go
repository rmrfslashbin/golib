package golib

import (
	"testing"
)

// TestRandString calls golib.TestRandString checking
// for a valid return value.
func TestRandString(t *testing.T) {
	randString := RandString(10)
	if len(randString) != 10 || randString == "" {
		t.Error("RandString returned an invalid string")
	}
}

func TestNumberFormat(t *testing.T) {
	number := NumberFormat(1234567890)
	if number != "1.23B" {
		t.Error("NumberFormat returned an invalid string: " + number)
	}
}

func TestBrowserOpen(t *testing.T) {
	if err := BrowserOpen("https://github.com/rmrfslashbin/golib"); err != nil {
		t.Error("BrowserOpen returned error: " + err.Error())
	}
}
