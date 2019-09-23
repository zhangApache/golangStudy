package main

import (
	"reflect"
	"log"
)

/**
  反射
  要去反射是一个类型的值(这些值都实现了空interface)，
  首先需要把它转化成reflect对象(reflect.Type或者reflect.Value，根据不同的情况调用不同的函数)。
 */

type testRe struct {
	name string
	age int
}
func main() {
	i := 5

	t := reflect.TypeOf(i)
	v := reflect.ValueOf(i)

	log.Println(t, v)

	r := testRe{"zhangsan", 34}

	t1 := reflect.TypeOf(r)
	v1 := reflect.ValueOf(r)

	tag := t1.Elem().Field(0).Tag
	name := v1.Elem().Field(0).String()
	 log.Println(tag, name)

}
