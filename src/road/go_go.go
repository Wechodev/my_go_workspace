package main

import (
	"fmt"
	"runtime"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		runtime.Gosched()//表示把CPU让给别人，下次某些时候继续恢复执行goroutine
		fmt.Println(s)
	}
}

func main()  {
	go say("world")//开启一个新的Goroutines执行
	say("hello")//当前Goroutines执行
}

//默认情况下调度器仅适用单线程，也就说只实现了并发，想要发挥多核处理器的并行，需在我们程序中显示调用，runtime.GOMAXPROCS(n)告诉调度器
//同时使用多个线程，GOMAXPROCS设置了同时运行逻辑代码的系统线程的最大数量，并返回之前的设置，如果n < 1.不会改变当前设置，以后Go的新版这将被移除



