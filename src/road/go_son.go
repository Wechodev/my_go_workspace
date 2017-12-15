package main

import "fmt"

type Human struct {
	name string
	age int
	phone string
}

type Student struct {
	Human //匿名字段
	school string
}

type Employee struct {
	Human //匿名字段
	company string
}

//在human上定义一个method

func (h *Human) SayHi() {
	fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}

func main()  {
	mark := Student{Human{"Mark", 25, "456825212"}, "MTH"}
	sam := Employee{Human{"Sam", 45, "455834545"}, "Golang"}
	mark.SayHi()
	sam.SayHi()
}
