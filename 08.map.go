package main

import "fmt"

func main() {
	a := make(map[string]string)

	a["name"] = "Gopher"
	a["age"] = "35"

	fmt.Println(a)

	x, ok := a["x"] // ok = false, x =

	fmt.Println(ok)
	fmt.Println(x)

	for k, v := range a {
		fmt.Println(k, ":", v)
	}

	fmt.Println("-----------")

	b := map[string]string{
		"name": "Gopher",
		"age":  "35",
	}

	fmt.Println(b)

	for k, v := range b {
		fmt.Println(k, ":", v)
	}
}
