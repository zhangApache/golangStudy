package main

import "fmt"

/*
  简单的说，interface是一组method签名的组合，
  我们通过interface来定义对象的一组行为。
 */
type Human2 struct {
	name string
	age int
	phone string
}
type Student2 struct {
	Human2
	school string
	loan float32
}
type Employee2 struct {
	Human2
	company string
	money float32
}

func (h *Human2) SayHi()  {
	fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}

func (h *Human2) Sing(lyrics string){
	fmt.Println("La la, la la la, la la la la la...", lyrics)
}

func (h *Human2)Guzzle(beerStein string)  {
	fmt.Println("Guzzle Guzzle Guzzle...", beerStein)
}

//Employee2 重载Human2的SayHi方法
func (e *Employee2) SayHi()  {
	fmt.Printf("Hi, I am %s, I work at %s. Call me on %s\n", e.name,
		e.company, e.phone) //此句可以分成多行
}
//student 实现借钱方法
func (s *Student2) BorrowMoney(amount float32)  {
	s.loan += amount
}

//employee2 实现加班工资方法
func (e *Employee2) SpendSalary(amount float32)  {
	e.money -= amount
}

type Men interface {
	SayHi()
	Sing(lyrics string)
	Guzzle(beerStein string)
}
type YoungChap interface {
	SayHi()
	Sing(song string)
	BorrowMoney(amount float32)
}

type ElderlyGent interface {
	SayHi()
	Sing(song string)
	SpendSalary(amount float32)
}

/**
  interface可以被任意的对象实现,
  任意的类型都实现了空interface(我们这样定义：interface{})，
  也就是包含0个method的interface。
 */
func main() {
	
}
