package main

import "fmt"

func Add(x, y int) {
	z := x + y
	fmt.Println(z)
}

func main() {
    for i := 0; i < 10; i++ {
    	go Add(i, i)
    }
  	
}

//在这个例子中main函数执行了add但未等Add返回main函数已经执行完毕结束了主程序，导致无任何输出内容