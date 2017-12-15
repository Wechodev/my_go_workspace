package main

import "fmt"

type Base struct {
	FirstName, LastName string
    Age float32
}

func (base *Base) HasFeet() {
	fmt.Println(base.FirstName + base.LastName + "has feet! Base")
}

func (base *Base) Flying() {
	//fmt.Println("Base Can flying!")
}

type Sub struct {
	Base
	Area string
}

func (sub *Sub) Flying() {
	sub.Base.Flying()
	fmt.Println("Sub flying")
}

func main() {
	chk := new(Sub)
	chk.Flying()
	chk2 := &Sub{Base{"Bob", "Smith", 25}, "list"}
	fmt.Println(chk2.Area)
}