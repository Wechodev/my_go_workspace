package main

import "fmt"

func main()  {
	staticDir := "./public"
	name := "/assets/"
	result := staticDir + name
	fmt.Println(result)
}
