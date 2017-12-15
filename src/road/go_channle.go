package main

import "fmt"

func sum(a []int, c chan int) {
	sum := 0
	for _, v := range a {
		sum += v
	}
	c <- sum
}

func main()  {
	a := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)
	x, y := <-c, <-c
	fmt.Println(x, y, x + y)
}

//默认情况下，channel接收和发送数据都是阻塞的，除非另一端已经准备好。这样就使goroutine同步变得更加的简单，而不需要显示的loc看，所谓阻塞
//也就是如果读取(value := <-ch)它将会被阻塞，直到有数据接收，其次，任何发送ch<-5将会被阻塞，直到数据被读出，无缓存channel是在多个goroutine之间
//同步很棒的工具
