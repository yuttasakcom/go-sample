package utils

import "testing"

func TestAddTwoNumber(t *testing.T) {

	res := Add(1, 2)

	if res != 3 {
		t.Error("คาดหวังว่าจะได้ 3 ผลลัพธ์คือ ", res)
	}
}
