package main

import "fmt"

func example(x int) int {//避免这样使用。避免在if中做返回
	if a:=3;x < a {
		return 5
	} else {
		return x
	}
}


func flowSwitch(i int) {

	switch i {
		case 0:
		   fmt.Println("0")
	    case 1:
           fmt.Println("1")
        case 2:
            fallthrough
        case 3:
            fmt.Println("3")
        case 4, 5, 6:
             fmt.Println("4, 5, 6")
        default:
             fmt.Println("default")                  
	}	
}

func flowFor() {
	/*sum := 0
	for i:=0; i < 10; i++ {
		sum += i
	}*/
	/*sum := 0
	for {
		sum++
		if sum > 100 {
			break
		}
	}*/

	/*a := []int{1, 2, 3, 4, 5, 6}
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}

    return a*/

     JLoop:for j := 0; j < 5; j++ {
    	for i :=0; i < 10 ; i++ {
    		if i > 5 {

    			break JLoop
    		}
    		fmt.Println(i)
    	}
    }
}

func myfunc() {
	i := 0
	HERE:
	fmt.Println(i)
	i++
	if i < 10 {
		goto HERE
	}//此函数自身形成了类型for的循环
}

func myfunc1(args ...int) {
	for _, arg := range args {
		fmt.Println(arg)
	}
	
}

func myfunc2(args ...interface{}) {
	 for _, v := range args {
	 	fmt.Println(v)
	 }
}

func MyPrintf(args ...interface{}) {
	for _, arg := range args {
		switch arg.(type) {
			case int:
				fmt.Println(arg, "is an int value")
			case string:
			    fmt.Println(arg, "is a string value")
			case int64:
			    fmt.Println(arg, "is an int64 value")
			default:
			     fmt.Println(arg, "is an unknown type")    
		}
	}
}

func main() {
    //a := 2
    //b := example(a)
    //fmt.Println(b)
    //flowSwitch(a)
    //flowFor()
    //fmt.Println(b)
    //myfunc()
    //myfunc1(1, 2, 3, 4, 5)
    //var b int64 = 234
    //MyPrintf(1, b, "hello", 1.35)
    f := func (x, y int) int {
    	 return x+y 
    }

    fmt.Println(f(1,6))
}









