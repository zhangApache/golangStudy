package main

import "log"

/*
  在Go中函数也是一种变量，我们可以通过type来定义它，
  它的类型就是所有拥有相同的参数，相同的返回值的一种类型
 */
type testInt func(int) bool // 声明了一个函数类型

func isOdd(integer int) bool {
	if integer%2 == 0 {
		return false
	}
	return true
}

func isEven(integer int) bool {
	if integer%2 == 0{
		return true
	}
	return false
}

func filter(slice []int, f testInt) []int {
	var result []int
	for value := range slice {
		if f(value) {
			result = append(result, value)
		}
	}
	return result
}

func main() {
	slice := []int{1,2,3,4,5,7}
	log.Println("slice = ", slice)
	odd := filter(slice, isOdd)
	log.Println("Odd elements of slice are: ", odd)
	even := filter(slice, isEven)
	log.Println("Even elements of slice are: ", even)
}