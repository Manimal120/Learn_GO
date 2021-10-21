package main

import (
	"fmt"
	"time"
)

func main() {
	// 主例程完全打印， 工作例程不一定
	go chars("goroutine")
	chars("main")
	//time.Sleep(time.Second*3)
}

func chars(prefix string) {
	for c := 'A'; c <= 'Z'; c++ {
		fmt.Printf("%s: %c\n", prefix, c)
		time.Sleep(time.Microsecond)
	}
}
