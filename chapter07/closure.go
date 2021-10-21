package main

import (
	"fmt"
	"sync"
)

func main() {

	// 001. 10
	//for i := 0; i < 10; i++ {
	//	go func() {
	//		fmt.Println(i)
	//	}()
	//}
	//time.Sleep(time.Second)

	// 002. 0-9
	//for i := 0; i < 10; i++ {
	//	go func(i int) {
	//		fmt.Println(i)
	//	}(i)
	//}
	//time.Sleep(time.Second)

	// 003. WaitGroup
	var wg sync.WaitGroup

	// wg.Add(10)
	for i := 0; i < 10; i++ {
		wg.Add(1)        // Add加在启动例程同级
		go func(i int) { // 匿名函数不需要传递指针，因为子块可以访问父块
			fmt.Println(i)
			wg.Done() // Done放在协程中
		}(i)

	}
	fmt.Println("Wait")
	wg.Wait()
	fmt.Println("Over")
}
