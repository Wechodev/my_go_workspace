package main

import "fmt"

func main() {
	var myArray [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	//基于数组创建一个数组切片
	var mySlice []int = myArray[:5]

	fmt.Println("Element of myArray: ")
	for _, v := range myArray {
		fmt.Println(v, " ")
	}
	fmt.Println("Element of mySlice: ")
	for _, v := range mySlice {
		fmt.Println(v, " ")
	}
	var myNewSlice []int = myArray[:]
	fmt.Println("Element of mySlice(All): ")
	for _, v := range myNewSlice {
		fmt.Println(v, " ")
	}

	fmt.Println("Element of mySlice(5): ")
	var myFiveSlice []int = myArray[5:]
	for _, v := range myFiveSlice {
		fmt.Println(v, " ")
	}
	fmt.Println("make()1")
	mySlice1 := make([]int, 5)
	for _, v := range mySlice1 {
		fmt.Println(v, " ")
	}
	fmt.Println("make()2")
	mySlice2 := make([]int, 5, 10)
	for _, v := range mySlice2 {
		fmt.Println(v, " ")
	}
	fmt.Println("直接创建")
	mySlice3 := []int{1, 2, 3, 4, 5}
	for _, v := range mySlice3 {
		fmt.Println(v, " ")
	}
	fmt.Println()
}
