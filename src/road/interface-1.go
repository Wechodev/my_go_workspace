package main

import "fmt"

type ISpeaker interface {
	Speak()
}

type SimpleSpeaker struct {
	Message string
}

func (speaker *SimpleSpeaker) Speak() {
	fmt.Println("I am speaking?", speaker.Message)
}

func main() {
	var speaker ISpeaker
	    speaker = &SimpleSpeaker{"Hello"}
	    speaker.Speak()
}
//从机器的角度如何判断一个SimpleSpeaker类型实现了ISpeaker接口的方法
//一个简单的逻辑就是需要获取这个类型的所有方法的集合A，并获取该接口包含的所有方法的集合B，然后判断列表B是否为列表A的子集，则意味着SimpleSpeaker
//类型实现了ISpeaker接口