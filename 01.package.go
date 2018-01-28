package main

import (
	"fmt"

	"github.com/yuttasakcom/go-sample/utils"
)

func main() {
	fmt.Println(utils.Hello())
	fmt.Println("Sum 1 + 2 = ", utils.Add(1, 2))
}
