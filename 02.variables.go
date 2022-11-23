package main

import (
	"fmt"
	"reflect"
	"strconv"
	"unicode/utf8"
)

var Pi = 3.14

func main() {

	var name string
	name = "Gopher"
	fmt.Print(name)
	fmt.Printf(" type is %T\n", name)

	var age = 35
	fmt.Print(age)
	fmt.Printf(" type is %T\n", age)
	fmt.Printf("Convert %d to string = %q\n", age, strconv.Itoa(age))

	grade := 3.47
	fmt.Print(grade)
	fmt.Printf(" type is %T\n", grade)

	// Convert
	sum := float64(age) * Pi
	fmt.Println("sum: ", sum)
	fmt.Println("type of sum =", reflect.TypeOf(sum))

	greeting := "สวัสดี"
	fmt.Println("Len: ", len(greeting))
	fmt.Println("RuneCount: ", utf8.RuneCountInString(greeting))

	// ## Output
	// Gopher type is string
	// 35 type is int
	// Convert 35 to string = "35"
	// 3.47 type is float64
	// sum:  109.9
	// type of sum = float64
	// Len:  18
	// RuneCount:  6
}
