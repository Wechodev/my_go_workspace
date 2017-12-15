package main

import "fmt"

type Human struct {
	name string
	age int
	phone string
}

type Student struct {
	Human
	school string
}

type Employee struct {
	Human
	company string
}

func (h *Human) SayHi() {
	fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}

func (e *Employee) SayHi() {
	fmt.Printf("Hi, I am %s, I work at %s. Call me on %s\n", e.name, e.company, e.phone)
}

func main()  {
	mark := Student{Human{"Mark", 25, "45775345242"}, "MIT"}
	sam := Employee{Human{"Sam", 45, "4351278785"}, "Golang"}
	mark.SayHi()
	sam.SayHi()
}
