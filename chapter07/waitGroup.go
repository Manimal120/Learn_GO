package main

import (
	"fmt"
	"sync"
	"time"
)

func chars1(prefix string, wg *sync.WaitGroup) {
	for c := 'A'; c <= 'Z'; c++ {
		fmt.Printf("%s: %c\n", prefix, c)
		time.Sleep(time.Microsecond)
	}
	wg.Done()
}

func chars2(prefix string) {
	for c := 'A'; c <= 'Z'; c++ {
		fmt.Printf("%s: %c\n", prefix, c)
		time.Sleep(time.Microsecond)
	}
}

func main() {
	var wg sync.WaitGroup // 结构体是值类型，在结构体之间传递需要指针
	// Add(n)  add n signals
	// Done()  done n signals
	// Wait()  wait counter to 0, all signals were done

	wg.Add(1)
	go chars1("Goroutine", &wg)

	chars2("main")

	fmt.Println("Wait")
	wg.Wait()
	fmt.Println("Over")
}
