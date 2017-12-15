package main

import (
	"fmt"
	"time"
	"strconv"
)

func main() {
	t1 := time.Now()
	fmt.Println(t1)
	fmt.Println(t1.Unix())
	timestamp := strconv.FormatInt(t1.UTC().UnixNano(), 10)
	fmt.Println(timestamp)
	timestamp = timestamp[:10]
	fmt.Println(timestamp)
	t := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	fmt.Printf("Go launched at %s\n", t.Local())
}
