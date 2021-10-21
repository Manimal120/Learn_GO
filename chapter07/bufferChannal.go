package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var chan1 chan int
	chan1 = make(chan int, 3)

	//fmt.Printf("%T, %#v, %d\n", chan1, chan1, cap(chan1))

	//var wg sync.WaitGroup
	//wg.Add(2)

	// situation 1
	//go func() {
	//	for i := 0; i < 3; i++ {
	//		time.Sleep(time.Second)
	//		fmt.Println("<-", i)
	//		chan1 <- i
	//	}
	//	wg.Done()
	//}()
	//
	//go func() {
	//	for i := 0; i < 3; i++ {
	//		fmt.Println(<-chan1)
	//	}
	//	wg.Done()
	//}()

	// situation 2
	go func() {
		for i := 0; i < 3; i++ {
			time.Sleep(time.Second)
			fmt.Println("<-", i)
			chan1 <- i
		}
		close(chan1)
	}()

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		for ch := range chan1 { // for range 需要 close 管道
			fmt.Println(ch)
		}
		wg.Done()
	}()

	wg.Wait()

}
