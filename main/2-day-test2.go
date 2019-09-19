package main

import "fmt"

func testdefer()  {
	for i:=0; i < 5; i++ {
		fmt.Print(i)
		defer fmt.Print("a")
	}
}
func main() {
	testdefer()
}