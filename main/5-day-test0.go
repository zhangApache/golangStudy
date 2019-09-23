package main

import (
	"strconv"
	"fmt"
)

/**
  Go语言里面有一个语法，可以直接判断是否是该类型的变量： value, ok = element.(T)，
  这里value就是变量的值，ok是一个bool类型，element是interface变量，T是断言的类型。
  如果element里面确实存储了T类型的数值，那么ok返回true，否则返回false。
 */

type Element interface{}
type List [] Element

type Person struct {
	name string
	age int
}

func (p Person) String() string {
	return "(name: " + p.name + "- age: "+ strconv.Itoa(p.age)+ " years)"
}

func main() {
	list := make(List, 3)
	list[0] = 1
	list[1] = "Hello"
	list[2] = Person{"Dennis", 27}

	for index, element := range list {
		if value, ok := element.(int); ok {
			fmt.Printf("list[%d] is an int and its value is %d\n", index, value)
		} else if value, ok := element.(string); ok {
			fmt.Printf("list[%d] is a string and its value is %s\n", index, value)
		} else if value, ok := element.(Person); ok {
			fmt.Printf("list[%d] is a Person and its value is %s\n", index, value)
		} else {
			fmt.Printf("list[%d] is of a different type\n", index)
		}
	}
}
