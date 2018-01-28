package utils

import "testing"

func TestHello(t *testing.T) {
	msg := Hello()

	if msg != "Hello!" {
		t.Error("คาดหวังว่าจะแสดง Hello! ผลลัพธ์คือ ", msg)
	}
}
