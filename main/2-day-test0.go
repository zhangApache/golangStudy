package main

import (
	"fmt"
	"log"
)

func test20() {
	// 计算获取值x,然后根据x返回的大小，判断是否大于10。
	if x := 5; x <= 10 {
		fmt.Println("x is less than 10")
	} else {
		fmt.Println("x is greater than 10")
	}
}

/*
  Go有goto语句——请明智地使用它。用goto跳转到必须在当前函数内定义的标签。
 */
func test21()  {
	i := 0
	Home:
		println(i)
		i++
		if i < 100 {
			goto Home
		}else {
			println("over")
		}
}

func test22()  {
	sum := 0;
	for index:=0; index < 10 ; index++ {
		sum += index
		fmt.Println("sum: ", sum)
	}
	fmt.Println("sum is equal to ", sum)
}
// 等价于while
func test23()  {
	sum := 1
	for sum < 10 {
		sum ++
	}
	fmt.Println(sum)
}

func test24()  {
	for index := 10;index > 0 ; index --  {
		if index == 5 {
			break
		}
		fmt.Println(index)
	}
}
/*
  for配合range可以用于读取slice和map的数据
 */
func test25()  {
	flag1 := make(map[string]string)
	flag1["test1"] = "1"
	flag1["test2"] = "2"
	flag1["test3"] = "3"
	for k , v := range flag1{
		log.Println(k,v)
	}
}
/*
Go里面switch默认相当于每个case最后带有break，
匹配成功后不会自动向下执行其他case，
而是跳出整个switch, 但是可以使用fallthrough强制执行后面的case代码。
 */
func test26()  {
	integer := 6
	switch integer {
	case 4:
		fmt.Println("The integer was <= 4")
		fallthrough
	case 5:
		fmt.Println("The integer was <= 5")
		fallthrough
	case 6:
		fmt.Println("The integer was <= 6")
		fallthrough
	case 7:
		fmt.Println("The integer was <= 7")
		//fallthrough
	case 8:
		fmt.Println("The integer was <= 8")
		//fallthrough
	/*case 3:
		fmt.Println("The integer was <= 3")
		fallthrough*/
	default:
		fmt.Println("default case")
	}
}
func main()  {
	test26()
}

