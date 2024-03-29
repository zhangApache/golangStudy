package main

import (
	"log"
)
/*  map是无序的，每次打印出来的map都会不一样，它不能通过index获取，而必须通过key获取
	map的长度是不固定的，也就是和slice一样，也是一种引用类型
	内置的len函数同样适用于map，返回map拥有的key的数量
	map的值可以很方便的修改，通过numbers["one"]=11可以很容易的把key为one的字典值改为11
	map和其他基本型别不同，它不是thread-safe，在多个go-routine存取时，必须使用mutex lock机制*/
func main() {
	testmap := make(map[int]string)
	testmap[1] = "one"
	vale, ok := testmap[2]
	if ok {
		log.Println(vale)
	}else {
		log.Println("no value found")
	}
	maplen := len(testmap)
	log.Println(maplen)
}
