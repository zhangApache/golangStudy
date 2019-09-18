package main

import (
	"fmt"
	"log"
	"sync"
	"reflect"
)
/*
  关于Go语言中new和make是内建的两个函数，主要用来创建分配类型内存
 */

 //对于引用类型的变量，我们不光要声明它，还要为它分配内容空间，否则我们的值放在哪里去呢？
func test1() {
	var i *int
	*i = 10
	fmt.Println(*i)
}
func test2()  {
	//var i *int
	i := new(int)
	*i = 10
	//fmt.Println()
	log.Println(*i)
}

type user struct {
	lock sync.Mutex
	name string
	age int
}
/*
示例中的user类型中的lock字段我不用初始化，直接可以拿来用，不会有无效内存引用异常，因为它已经被零值了。
这就是new，它返回的永远是类型的指针，指向分配类型的内存地址。
 */
func test3()  {
	u := new(user)
	//u.lock.Lock()
	u.name = "zhangsan"
	u.age = 19
	//u.lock.Unlock()
	fmt.Println(reflect.ValueOf(u))
	fmt.Println(reflect.TypeOf(u))
}

/*
	make也是用于内存分配的，但是和new不同，
	它只用于chan、map以及切片的内存创建，
	而且它返回的类型就是这三个类型本身，
	而不是他们的指针类型，
	因为这三种类型就是引用类型，
	所以就没有必要返回他们的指针了。
 */
func test4()  {
	u := user{}
	//u.lock.Lock()
	u.name = "zhangsan"
	u.age = 19
	//u.lock.Unlock()
	fmt.Println(reflect.ValueOf(u))
	fmt.Println(reflect.TypeOf(u))
}
func main() {
	test4()
}



