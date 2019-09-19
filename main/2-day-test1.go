package main

import (
	"fmt"
	"log"
)

func max(a, b int) int {
	if a > b{
		return a
	}else {
		return b
	}
}

func myfunc(arg ...int)  {
	for _, n := range arg {
		fmt.Printf("And the number is: %d\n", n)
	}

}
/*
  当我们传一个参数值到被调用函数里面时，
  实际上是传了这个值的一份copy，
  当在被调用函数中修改参数值的时候，
  调用函数中相应实参不会发生任何变化，因为数值变化只作用在copy上。
 */
func add1(a int) int{
	a = a + 1
	log.Println(&a)
	return a
}

/*
  指针类型，才能在函数中修改x变量的值。
  此时参数仍然是按copy传递的，只是copy的是一个指针。
 */
func add2(a *int) int {
	*a = *a + 1
	log.Println(a)
	return *a
}
/*
  指针传值的优点
  1.传指针使得多个函数能操作同一个对象。
  2.传指针比较轻量级 (8bytes),只是传内存地址，
    我们可以用指针传递体积大的结构体。
    如果用参数值传递的话, 在每次copy上面就会花费相对较多的系统开销（内存和时间）。
    所以当你要传递大的结构体的时候，用指针是一个明智的选择。
  3.Go语言中channel，slice，map这三种类型的实现机制类似指针，所以可以直接传递，而不用取地址后传递指针。
   （注：若函数需改变slice的长度，则仍需要取地址传递指针）
 */
func main() {
	x := 3
	log.Println(&x)
	//x1 := add1(x)
	x1 := add2(&x)
	log.Println(x1)
	log.Println(x)
}