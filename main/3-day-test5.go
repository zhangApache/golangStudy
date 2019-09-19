package main

import "fmt"

/**
  method 继承
 */
type Human1 struct {
	name string
	age int
	phone string
}

type Student1 struct {
	Human1
	school string
}

type Employee struct {
	Human1
	company string
}

func (h *Human1) SayHi()  {
	fmt.Printf("Hi, I am %s you can call me on %\n", h.name, h.phone)
}
func (e *Employee) SayHi()  {
	fmt.Printf("Hi, I am %s you can call me on %\n", e.name, e.phone)
}
func main() {
	mark := Student1{Human1{"Mark", 25, "18940932065"}, "QingHua"}
	sam := Employee{Human1{"Sam", 45, "110456"}, "Google"}

	mark.SayHi()
	sam.Human1.SayHi()
}
