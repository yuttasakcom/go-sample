package utils

import (
	"testing"
)

func TestReverseString(t *testing.T) {
	msg := reverse()

	if msg != "!dlroW olleH" {
		t.Error("Expected !dlroW olleH", msg)
	}
}
