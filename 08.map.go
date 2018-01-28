package main

import "fmt"

func main() {
	a := make(map[string]string)

	a["name"] = "Gopher"
	a["age"] = "35"

	fmt.Println(a)

	x, ok := a["x"]

	fmt.Println(ok)
	fmt.Println(x)
}
