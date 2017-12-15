package main

import (
	"fmt"
	"strconv"
)

type Element interface {}
type List[] Element

type Person struct {
	name string
	age int
}

func (p Person) String() string {
	return "(name: "+ p.name +" - age: "+strconv.Itoa(p.age)+" years)"
}

func main()  {
	list := make(List, 3)
	list[0] = 1
	list[1] = "Hello"
	list[2] = Person{"TOM", 70}

	for index, element := range list {
		switch value := element.(type) {
		    case int:
		    	fmt.Printf("list[%d] is an int and its value is %d\n", index, value)
		    case string:
				fmt.Printf("list[%d] is an string and its value is %s\n", index, value)
		    case Person:
				fmt.Printf("list[%d] is an interface and its value is %s\n", index, value)
		    default:
		    	fmt.Println("list[%d] is of a different type", index)
		}
	}
}
//只有在switch下才能element.(type)语法
//断言语法