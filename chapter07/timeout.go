package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func main() {
	res := make(chan interface{})
	timeout := make(chan struct{})

	go func() {
		time.Sleep(time.Second * 5)
		timeout <- struct{}{}
	}()

	go func() {
		r := rand.Intn(10)
		fmt.Println(r)
		time.Sleep(time.Second * time.Duration(r))
		res <- r
	}()

	select {
	case <-timeout:
		fmt.Println("Over time...")
	case r := <-res:
		fmt.Println("Success, ", r)

	}
}
