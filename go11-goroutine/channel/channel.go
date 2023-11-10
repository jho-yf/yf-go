package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("No cache channel demo >>>>>>>>>>>>>>")
	noCacheChannel()

	fmt.Println("With cache channel demo >>>>>>>>>>>>>>")
	withCacheChannel()

	fmt.Println("Close cache channel demo >>>>>>>>>>>>>>")
	closeChannel()

	fmt.Println("range channel demo >>>>>>>>>>>>>>")
	rangeChannel()

	fmt.Println("select channel demo >>>>>>>>>>>>>>")
	selectChannel()
}

func noCacheChannel() {
	// 定义一个channel
	c := make(chan int)

	go func() {
		defer fmt.Println("sub goroutine end...")

		fmt.Println("sub goroutine running...")

		// 将666写进channel
		c <- 666
	}()

	fmt.Println("main goroutine running...")

	// 从channel中阻塞取数据赋值给num
	num := <-c
	fmt.Println("num =", num)
	fmt.Println("main goroutine end...")
}

func withCacheChannel() {
	// 带缓存的channel
	c := make(chan int, 3)

	fmt.Println("[main] len(c) = ", len(c), ", cap(c)", cap(c))

	go func() {
		defer fmt.Println("[sub] end...")

		for i := 0; i < 3; i++ {
			// 当channel已经满了，再写数据就会阻塞
			c <- i
			fmt.Println("[sub] element =", i, "len(c) =", len(c), "cap(c) =", cap(c))
		}
	}()

	time.Sleep(3 * time.Second)
	for i := 0; i < 3; i++ {
		// 当channel为空，从里面取数据就会阻塞
		num := <-c
		fmt.Println("[main] num = ", num)
	}

	fmt.Println("[main] end...")
}

func closeChannel() {
	c := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}

		// 关闭channel，关闭channel之后再写数据会报panic的错误
		close(c)
	}()

	for {
		if data, ok := <-c; ok {
			fmt.Println(data)
		} else {
			break
		}
	}

	fmt.Println("[main] end...")
}

func rangeChannel() {
	c := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}

		// 关闭channel，关闭channel之后再写数据会报panic的错误
		close(c)
	}()

	for data := range c {
		fmt.Println(data)
	}

	fmt.Println("[main] end...")
}

func selectChannel() {
	// 使用select可以监控多个channel的状态

	c := make(chan int)
	quit := make(chan int)

	go func() {
		for i := 0; i < 6; i++ {
			// 阻塞打印值
			fmt.Println(<-c)
		}

		quit <- 0
	}()

	fibonacii(c, quit)
}

func fibonacii(c, quit chan int) {
	x := 1
	y := 1

	for {
		select {
		case c <- x:
			x = y
			y = x + y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}
