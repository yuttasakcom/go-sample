package utils

import "testing"

func TestReverse(t *testing.T) {
	msg := Reverse("Hello World!")

	if msg != "!dlroW olleH" {
		t.Error("อยากได้ !dlroW olleH แต่ได้ ", msg)
	}
}
