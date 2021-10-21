package main

import (
	"fmt"
	"time"
)

func main() {
	var channel chan int
	fmt.Printf("%T, %#v\n", channel, channel)

	// 初始化 - 无缓冲管道 - 必须有写和读！
	channel = make(chan int)
	fmt.Printf("%T, %#v\n", channel, channel)

	// Write
	go func() {
		time.Sleep(time.Second * 5)
		channel <- 1
	}()

	fmt.Println("Before")
	// Read
	num := <-channel // need another goroutine to write, or 阻塞
	fmt.Println(num)
}
