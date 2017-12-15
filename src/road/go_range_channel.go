package main

import "fmt"

func fibonacci(n int, c chan int)  {
	x, y := 1, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x + y
	}
	close(c)
}
func main() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}

//for能都不断读取channel里面的数据，直到channel被显式关闭，，上面代码我们看到可以显式关闭channel，生产者通过关键字close函数关闭channel，
//关闭了channel之后就无法再发送任何数据了，在消费方可以通过语法v,ok := <-ch测试是否被关闭，如果ok返回FALSE，那么说明channel已经没有
//任何数据了并且已经被关闭

