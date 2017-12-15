package main

import "fmt"
import "os"
import "syscall"

type PathError struct {
	Op   string
	Path string
	Err  error
}

func (e *PathError) Error() string {
	 
	 return e.Op + " " + e.Path + ": " + e.Err.Error()
}

func Stat(name string) (fi os.FileInfo, err error) {
	var Stat syscall.Stat_t
	
	err = syscall.Stat_t

    if err != nil {
     	
     	return nil, &PathError("Stat", name, err)
     } 

     return fileInfoFromStat(&Stat, name), nil
}

func main() {
	n ,err := Stat("jack")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(n)
	}
}
//此处没有看懂