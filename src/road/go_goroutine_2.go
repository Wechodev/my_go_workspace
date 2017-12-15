package main

import "fmt"
var a int = 0
func Count(ch chan int)  {
	ch <- a
	a++
	fmt.Println(a)
}

func main()  {
	chs := make([]chan int, 10)
	for i := 0;i < 10 ; i++ {
		chs[i] = make(chan int)
		go Count(chs[i])
	}
	for _, ch := range(chs) {
		<-ch
	}
	fmt.Println(len(chs))
}

/*
每次循环将一个数字写进了channel造成阻塞，在所有的goroutine启动完成后，通过<-ch语句从10个channel依次读数据
在对应的channel写入数据前，这个操作也是阻塞的，这样我就用channel实现了类似锁的功能，进而保证了所有goroutine完成后主函数才返回。
*/
