package main

import (
	"log"
	"reflect"
	"fmt"
)

//关于数组，长度也是数组类型的一部分，[3]int和[4]int是不同的类型
func arryTest() {
	arry1 := [...]int{3, 4 ,5} //[...]可以省略数组长度，由go语言自动获取
	log.Println(arry1)
	doubleArray := [2][4]int{[4]int{1, 2, 3, 4}, [4]int{5, 6, 7, 8}} //支持多维数组
	easyArray := [2][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}} //简化版
	log.Println(doubleArray[0][3])
	log.Println(easyArray[1][2])
}
//nil 切片和长度为0的切片是不相等的
func sliceTest()  {
	// 声明一个含有10个元素元素类型为byte的数组
	var ar = [10]byte {'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j'}
	// 声明两个含有byte的slice
	var a, b []byte
	// a指向数组的第3个元素开始，并到第五个元素结束，
	a = ar[2:5]
	//现在a含有的元素: ar[2]、ar[3]和ar[4]
	// b是数组ar的另一个slice
	b = ar[3:5]
	// b的元素是：ar[3]和ar[4]
	log.Println(a,b)
	var s1 []int
	var s2 = []int{}
	aa  := true
	if  reflect.DeepEqual(s1, s2){
		aa = true
	}else {
		aa = false
	}
	log.Println(aa,cap(s1),cap(s2))
}
func main() {
	//sliceTest()
	s1 := []int{0, 1, 2, 3, 6: 100}
	fmt.Println(s1, len(s1), cap(s1))
}