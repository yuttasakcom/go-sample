package main

import "fmt"

type person struct {
	Name     string
	NickName string
}

func mutatePerson(p person) {
	p.Name = "Hacker"
	fmt.Println(p)
}

func mutatePersonPointer(p *person) {
	p.Name = "Cracker"
}

func main() {
	p := person{
		Name:     "Facker",
		NickName: "Gopher",
	}

	fmt.Println(p)

	fmt.Println("---------")

	mutatePerson(p)

	fmt.Println("----------")
	mutatePersonPointer(&p)
	fmt.Println(p)
}
