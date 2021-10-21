package main

import "fmt"

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		ch1 <- 1
	}()

	go func() {
		ch2 <- 2
	}()

	select {
	case e := <-ch1:
		fmt.Println(e)
	case e := <-ch2:
		fmt.Println(e)
	default:
		fmt.Println("default")
	}
}
