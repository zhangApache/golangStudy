package main

import (
	"runtime"
	"fmt"
)

/**
  goroutine是Go并行设计的核心。
  goroutine说到底其实就是协程，但是它比线程更小，
  十几个goroutine可能体现在底层就是五六个线程，
  Go语言内部帮你实现了这些goroutine之间的内存共享。
  执行goroutine只需极少的栈内存(大概是4~5KB)，当然会根据相应的数据伸缩。
  也正因为如此，可同时运行成千上万个并发任务。
  goroutine比thread更易用、更高效、更轻便。

  goroutine是通过Go的runtime管理的一个线程管理器。
  goroutine通过go关键字实现了，其实就是一个普通的函数。
 */
func say(s string)  {
	for  i := 0; i < 5; i++ {
		runtime.Gosched()//runtime.Gosched()表示让CPU把时间片让给别人,下次某个时候继续恢复执行该goroutine。
		fmt.Println(s)
	}
}
/**
  设计上我们要遵循：不要通过共享来通信，而要通过通信来共享。

  goroutine运行在相同的地址空间，因此访问共享内存必须做好同步。
 */
func go_t()  {
	go say("word")
	say("Hello")
}

func main() {
	go_t()
}
