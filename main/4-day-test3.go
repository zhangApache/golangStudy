package main


/**
  空interface(interface{})不包含任何的method，
  正因为如此，所有的类型都实现了空interface。
  空interface对于描述起不到任何的作用(因为它不包含任何的method），
  但是空interface在我们需要存储任意类型的数值的时候相当有用，因为它可以存储任意类型的数值。
  它有点类似于C语言的void*类型。
 */

type a interface {

}
var i int = 5

func test111(A a)  {

}

func main() {
	
}
