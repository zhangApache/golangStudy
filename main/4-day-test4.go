package main

import (
	"strconv"
	"fmt"
)

/**
  interface的变量可以持有任意实现该interface类型的对象，
  这给我们编写函数(包括method)提供了一些额外的思考，
  我们是不是可以通过定义interface参数，让函数接受各种类型的参数。
 */

type Human4 struct {
	name string
	age int
	phone string
}

func (h Human4) String() string{
	return "❰"+h.name+" - "+strconv.Itoa(h.age)+" years -  ✆ " +h.phone+"❱"
}
func main() {
	Bob := Human4{"Bob", 24,"sadsads"}
	fmt.Println(Bob)
}
