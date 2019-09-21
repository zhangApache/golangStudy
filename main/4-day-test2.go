package main

import "fmt"

/**
  那么interface里面到底能存什么值呢？
  如果我们定义了一个interface的变量，
  那么这个变量里面可以存实现这个interface的任意类型的对象。
 */

type Human3 struct {
	name string
	age int
	phone string
}

type Student3 struct {
	Human3
	school string
	loan float32
}
type Employee3 struct {
	Human3
	company string
	money float32
}

//Human实现SayHi方法
func (h Human3) SayHi() {
	fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}

//Human实现Sing方法
func (h Human3) Sing(lyrics string) {
	fmt.Println("La la la la...", lyrics)
}

//Employee重载Human的SayHi方法
func (e Employee3) SayHi() {
	fmt.Printf("Hi, I am %s, I work at %s. Call me on %s\n", e.name,
		e.company, e.phone)
}

// Interface Men被Human,Student和Employee实现
// 因为这三个类型都实现了这两个方法
type Men1 interface {
	SayHi()
	Sing(lyrics string)
}
/**
  通过上面的代码，你会发现interface就是一组抽象方法的集合，
  它必须由其他非interface类型实现，而不能自我实现，
  Go通过interface实现了duck-typing:即"当看到一只鸟走起来像鸭子、
  游泳起来像鸭子、叫起来也像鸭子，那么这只鸟就可以被称为鸭子"。
 */
func main() {
	mike := Student3{Human3{"Mike", 25, "431564"},
	"qinghua", 0.00}
	paul := Student3{Human3{"Paul", 18,"4248621"},
	"beida",5.12}
	sam := Employee3{Human3{"Sam",28,"123456"},
	"zhongda",312}
	tom := Employee3{Human3{"Tom", 31, "52146"},
	"qq", 456}

	var i Men1
	i = mike
	fmt.Println("This is Mike, a Student:")
	i.SayHi()
	i.Sing("November rain")

	i = tom
	fmt.Println("This is Mike, an Employee:")
	i.SayHi()
	i.Sing("Born to be wild")

	fmt.Println("Let's use a slice of Men and see what happens")
	x := make([]Men1, 3)
	x[0], x[1], x[2] = paul, sam, mike

	for _, value := range x{
		value.SayHi()
	}
}
