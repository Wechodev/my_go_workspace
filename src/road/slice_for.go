package main

import "fmt"

func main() {
	var myArray [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	/*for i := 0; i < len(myArray); i++ {
		fmt.Println("myArray[",i,"]=", myArray[i])
	}*/

	for i, v := range myArray {
		
        fmt.Println("myArray[",i,"]", v)
	}
}