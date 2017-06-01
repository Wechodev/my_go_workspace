package main

import "fmt"

func main() {
	for n:=0; n<=8; n++{
		if n%2!=0{
		  continue
		}
		fmt.Println(n)
	}

}
