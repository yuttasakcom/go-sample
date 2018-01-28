package main

import (
	"fmt"
)

func main() {
	fmt.Println("Input score: ")

	var score int

	fmt.Scanln(&score)

	switch {
	case score > 50:
		fmt.Println("score มีค่ามากกว่า 50")
	default:
		fmt.Println("score มีค่าน้อยกว่า หรือเท่ากับ 50")
	}
}
