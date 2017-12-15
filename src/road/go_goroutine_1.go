package main

import "fmt"
import "sync"
import "runtime"

var counter int = 0

func Count(lock *sync.Mutex) {
	lock.Lock()
	counter++
	fmt.Println(counter)
	lock.Unlock()
}

func main() {
	lock := &sync.Mutex{}

	for i := 0; i < 10; i++ {
		go Count(lock)
	}

	for {
		lock.Lock()

		c := counter

		lock.Unlock()

		runtime.Gosched()
		if c >= 10 {
			break
		}
	}
}

//例子中在10个goroutine中共享了变量counter，每个goroutine执行完成后，将counter的值增加1，因为10个goroutine是并发
//执行的，所以我们还引入了锁，也就是代码中lock变量，每次对N的操作，都要先将锁锁住，操作完成后，再将锁打开
//在主函数中，使用for循环来不断检查counter的值，同样需要加锁，当值达到10时，说明所有goroutine都执行完毕了，这时主函数返回，程序退出了

/*
事情好像糟糕了，实现如此简单的功能，却写出如此臃肿而难以理解的代码。想象一下，在一个大的系统中具有无数
的锁，无数变量共享，无数的业务逻辑与错误处理分支，那将是一场噩梦，这噩梦就是C/C++开发者所经历的，其实
JAVA和C#也好不到哪里去
*/
