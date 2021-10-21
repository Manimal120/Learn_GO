package main

import (
	"fmt"
)

func main() {

	chan1 := make(chan int)

	fmt.Println("Before")
	for i := 0; i < 2; i++ {
		go func(prefix int) {
			for ch := 'A'; ch <= 'Z'; ch++ {
				fmt.Printf("%d: %c\n", prefix, ch)
			}
			chan1 <- prefix // End and write into the channel
		}(i)
	}
	// 管道中读取3次数据证明三个例程执行结束
	for i := 0; i < 3; i++ {
		fmt.Println("Over: ", <-chan1)
	}
	fmt.Println("Done")
}
