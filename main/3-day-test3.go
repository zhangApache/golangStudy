package main

import (
	"math"
	"log"
)

/*
  虽然method的名字一模一样，但是如果接收者不一样，那么method就不一样
  method里面可以访问接收者的字段
  调用method通过.访问，就像struct里面访问字段一样
 */
type Rectangle struct {
	width, height float64
}
type Circle struct {
	radius float64
}

func (r Rectangle) area() float64 {
	return r.width * r.height
}
func (c Circle) area() float64 {
	return c.radius * c.radius * math.Pi
}

func main() {
 r1 := Rectangle{12, 2}
 r2 := Rectangle{8,2}
 c1 := Circle{5}
 c2 := Circle{6}

 log.Println(r1.area())
 log.Println(r2.area())
 log.Println(c1.area())
 log.Println(c2.area())
}
