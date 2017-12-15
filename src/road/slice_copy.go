package main

import "fmt"

func main() {
	
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{5, 4 ,3}

	copy(slice1, slice2)
	copy(slice2, slice1)
}