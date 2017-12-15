package main

import (
	"reflect"
	"fmt"
)

type T struct {
	 A int
	 B string
}

func main()  {
	t := T{203, "mh203"}
	s := reflect.ValueOf(&t).Elem()
	typeOfT := s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("%d: %s %s = %v\n", i,
			typeOfT.Field(i).Name, f.Type(), f.Interface())
	}
}
/*
可以看得出，对于结构的反射操作并没有根本上的不同，只是用了Field()方法来按索引获取对应的成员，当然在试图修改成员的值时，也需要注意可赋值属性
*/