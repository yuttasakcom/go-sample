package main

import (
	"fmt"
)

func main() {
	fmt.Println("Input score: ")

	var score int

	fmt.Scan(&score)

	if score > 50 {
		fmt.Println("score มีค่ามากกว่า 50")
	} else {
		fmt.Println("score มีค่าน้อยกว่า หรือเท่ากับ 50")
	}
}
