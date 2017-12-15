package main

import (
	"fmt"
	"reflect"
)

func main()  {
	var x float64 = 3.4
	/*fmt.Println("type:", reflect.TypeOf(x))
	v := reflect.ValueOf(x)
	fmt.Println("type:", v.Type())
	fmt.Println("kind is float64:", v.Kind() == reflect.Float64)
	fmt.Println("value:", v.Float())*/
	p := reflect.ValueOf(&x)

	fmt.Println("type of p:", p.Type())
	fmt.Println("settability of p:", p.CanSet())

	v := p.Elem()
	fmt.Println("settability of v:", v.CanSet())

	v.SetFloat(7.1)
	fmt.Println(v.Interface())
	fmt.Println(x)

}

/*
Type和Value都包含了大量的方法，其中第一个有用的方法应该是Kind,这个方法返回该值的类型的具体信息:Uint, Float64等，Value类型还包含了一系列
类型方法，比如Int(),用于返回对应的值
*/
