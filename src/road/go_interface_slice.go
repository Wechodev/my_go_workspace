package main

import "fmt"

type Human struct {
	name string
	age int
	phone string
}

type Student struct {
	Human
	school string
	loan float32
}

type Employee struct {
	Human
	company string
	money float32
}

func (h Human) SayHi() {
	fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}

func (h Human) Sing(lyrics string) {
	fmt.Println("lalalallalalalal...", lyrics)
}

func (e Employee) SayHi() {
	fmt.Printf("Hi, I am %s, I work at %s. Call me on %s\n", e.name, e.company, e.phone)
}

type Men interface {
	SayHi()
	Sing(lyrics string)
}

func main()  {
	mike := Student{Human{"Mike", 25, "12569865"}, "MIT", 0.00}
	paul := Student{Human{"Paul", 26, "56457454"}, "Harvard", 100}
	sam := Employee{Human{"Sam", 36, "63645454"}, "Go", 1000}
	Tom := Employee{Human{"Sam", 36, "4545455687"}, "PHP", 5000}

	var i Men
	i = mike
	fmt.Println("This is Mike, a Student:")
	i.SayHi()
	i.Sing("o wuwuwu")

	i = Tom
	fmt.Println("This is Tom, an Employee:")
	i.SayHi()
	i.Sing("give money")

	fmt.Println("Let's use a slice of Men and see what happens")
	x := make([]Men, 3)
	x[0], x[1], x[2] = paul, sam, mike

	for _, value := range x {
		value.SayHi()
	}
}

//interface就是一组抽象方法的集合，必须由其他非interface类型实现，而不能自我实现，go通过interface实现了duck-typing：即"当看到一只鸟走起来
// 像鸭子，游泳起来像鸭子，叫起来也像鸭子，那么我们可以称为鸭子"