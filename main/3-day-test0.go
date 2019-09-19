package main

import "log"

type person struct {
	name string
	age int
}

func Older(p1, p2 person) (person, int) {
	if p1.age > p2.age {
		return p1, p1.age - p2.age
	}
	return p2, p2.age - p1.age
}

func main()  {
	//赋值初始化
	var tom person
	tom.name, tom.age = "Tom", 18
	//两个字段都写清楚的初始化
	bob := person{name:"Bob", age: 20}
	//按照struct定义顺序初始化值
	paul := person{"Paul", 25}
	tb_Older, tb_diff := Older(tom, bob)
	log.Println(tb_Older, tb_diff)
	tp_Older, tp_diff := Older(bob, paul)
	log.Println(tp_Older, tp_diff)
}