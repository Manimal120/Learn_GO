package main

import (
	"fmt"
	"time"
)

func main() {
	chan1 := make(chan struct{})

	go func() {
		time.Sleep(2 * time.Second)
		close(chan1) // 关闭管道，不再写入
		//chan1 <- struct{}{}
	}()
	fmt.Println("B4")
	//fmt.Println(<-chan1) // 当管道关闭，读取管道返回数据，不再进行阻塞
	_, ok := <-chan1 // ok is FALSE, means that it didn't load the chan
	fmt.Println(ok)

	fmt.Println("After")
}
