package main

/*
#include <stdio.h>
void hello() {
     printf("Hello, Cgo! --  From C world.\n")
}
*/
import (
	"C"
)

func hello() int {
	return int(C.hello())
}

/*
这个块注释里直接写了个C函数，它使用C标准库里的printf()打印了一句话
还有另外一个问题，那就是如果这里的C代码需哟啊依赖一个非C标准库的第三方库，
Cgo提供了#cgo这样的伪C文法
*/

