package main

import (
	"log"
	"fmt"
)

/**
  goroutine运行在相同的地址空间，因此访问共享内存必须做好同步。
  那么goroutine之间如何进行数据的通信呢，Go提供了一个很好的通信机制channel。
  channel可以与Unix shell 中的双向管道做类比：可以通过它发送或者接收值。
  这些值只能是特定的类型：channel类型。定义一个channel时，也需要定义发送到channel的值的类型。注意，必须使用make 创建channel：
 */

func sum(a []int, s string, c chan int)  {
	total := 0
	for _, v := range a{
		log.Println(v)
		total += v
	}
	log.Println(s)
	c <- total
}

/**
  默认情况下，channel接收和发送数据都是阻塞的，
  除非另一端已经准备好，这样就使得Goroutines同步变的更加的简单，
  而不需要显式的lock。所谓阻塞，也就是如果读取（value := <-ch）它将会被阻塞，
  直到有数据接收。其次，任何发送（ch<-5）将会被阻塞，直到数据被读出。
  无缓冲channel是在多个goroutine之间同步很棒的工具。
 */
func chan_t()  {
	a := []int{2, 3, 6, 4, -2}
	c := make(chan int)
	go sum(a[:len(a)/2],"one", c)
	go sum(a[len(a)/2:], "two", c)
	x, y := <-c, <-c
	log.Println(x, y, x+y)
}

/**
  上面我们介绍了默认的非缓存类型的channel，
  不过Go也允许指定channel的缓冲大小，很简单，就是channel可以存储多少元素。
  ch:= make(chan bool, 4)，创建了可以存储4个元素的bool 型channel。
  在这个channel 中，前4个元素可以无阻塞的写入。
  当写入第5个元素时，代码将会阻塞，直到其他goroutine从channel 中读取一些元素，腾出空间。
 */
func chan_b()  {
	c := make(chan int, 2)
	c <- 1
	c <- 2
	log.Println(<-c)
	log.Println(<-c)
}

/**
  上面这个例子中，我们需要读取两次c，
  这样不是很方便，Go考虑到了这一点，所以也可以通过range，
  像操作slice或者map一样操作缓存类型的channel，请看下面的例子
 */
func fibonacci(n int, c chan int)  {
	x, y := 1, 1
	for i := 0; i < n ; i++  {
		c <- x
		x, y = y, x + y
	}
	close(c)
}
/**
  for i := range c能够不断的读取channel里面的数据，
  直到该channel被显式的关闭。上面代码我们看到可以显式的关闭channel，
  生产者通过内置函数close关闭channel。
  关闭channel之后就无法再发送任何数据了，
  在消费方可以通过语法v, ok := <-ch测试channel是否被关闭。
  如果ok返回false，那么说明channel已经没有任何数据并且已经被关闭。
 */
 /**
   记住应该在生产者的地方关闭channel，
   而不是消费的地方去关闭它，这样容易引起panic
   另外记住一点的就是channel不像文件之类的，不需要经常去关闭，
   只有当你确实没有任何发送数据了，或者你想显式的结束range循环之类的
  */
func chan_r()  {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	for i := range c{
		fmt.Println(i)
	}
}
func main() {
	chan_r()
}
