package main

import (
	"fmt"
	"os"
)

func main() {
	os.Mkdir("ark", 0777)
	os.MkdirAll("ark/test1/test2", 0777)
	err := os.Remove("ark")
	if err != nil {
		fmt.Println(err)
	}
	os.RemoveAll("ark")
}
