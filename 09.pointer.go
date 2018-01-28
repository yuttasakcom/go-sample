package main

import "fmt"

func main() {
	a := 10
	fmt.Println(a)

	ptrA := &a
	fmt.Println(ptrA)
	fmt.Println(*ptrA)

	*ptrA = 11
	fmt.Println(a)
}
