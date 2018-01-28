package main

import "fmt"

func main() {
	a := make([]int, 5)

	a[1] = 1
	a[3] = 3

	fmt.Println(a)

	a = append(a, 5)

	fmt.Println(a)

	fmt.Println("-------------")

	b := []int{1, 2, 3, 4, 5}

	fmt.Println(b[1:])  // ตั้งแต่ index = 1 , เป็นต้นไป
	fmt.Println(b[:3])  // ตั้งแต่ index = 0 , ถึง index = 3 แต่ไม่เอา index = 3
	fmt.Println(b[1:3]) // ตั้งแต่ index = 1 , ถึง index = 3 แต่ไม่เอา index = 3
	fmt.Println(b[:])   // ทั้งหมด
}
