package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func main() {
	var a, b = 10000, 10000
	var wg sync.WaitGroup
	var locker sync.Mutex // 互斥锁
	count := 1000

	wg.Add(2)

	// a->b
	go func() {
		for i := 0; i < count; i++ {
			money := rand.Intn(100)
			if a > money {
				//locker.Lock()
				//a -= money
				//time.Sleep(time.Millisecond)
				//b += money
				//locker.Unlock()

				// defer, 延迟执行必须在函数内！！！
				func() {
					locker.Lock()
					defer locker.Unlock()
					a -= money
					time.Sleep(time.Millisecond)
					b += money
				}()

			}
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < count; i++ {
			money := rand.Intn(100)
			if b > money {
				locker.Lock()
				b -= money
				time.Sleep(time.Microsecond)
				a += money
				locker.Unlock()
			}
		}
		wg.Done()
	}()

	wg.Wait()

	fmt.Printf("a : %d , b : %d\n", a, b)
	fmt.Printf("total: %d", a+b)
}
