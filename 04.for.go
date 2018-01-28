package main

import "fmt"

func main() {
	var n [5]int
	n[1] = 1
	n[3] = 3

	for i := 0; i < len(n); i++ {
		fmt.Println(n[i])
	}

	fmt.Println("--------------")

	for _, v := range n {
		fmt.Println(v)
	}
}
