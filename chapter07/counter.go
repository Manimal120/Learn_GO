package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	var counter int64
	var wg sync.WaitGroup
	var locker sync.Mutex

	for i := 0; i < 5; i++ {
		wg.Add(2)
		go func() {
			for i := 0; i < 100; i++ {
				//locker.Lock()
				//counter++
				//locker.Unlock()

				// 原子操作
				atomic.AddInt64(&counter, 1)
				time.Sleep(time.Microsecond)
			}
			wg.Done()
		}()
		go func() {
			for i := 0; i < 100; i++ {
				locker.Lock()
				counter--
				locker.Unlock()
				time.Sleep(time.Microsecond)
			}
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(counter)
}
