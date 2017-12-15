package main

import (
	"fmt"
	"os"
)

func main() {
	userFile := "test1.txt"
	fout, err := os.Create(userFile)
	if err != nil {
		fmt.Println(userFile, err)
		return
	}
	defer fout.Close()
	for i := 0;i < 10; i++ {
		fout.WriteString("Just a  good test!\r\n")
		fout.Write([]byte("Just a bad test!\r\n"))
	}
}
