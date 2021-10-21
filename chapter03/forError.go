package main

import "fmt"

func main() {
	// 直接放入for循环会导致文件一直打开，不关闭
	// 解决方法: for内加入匿名函数，函数结束的时候，defer延迟执行正常
	for i := 0; i < 3; i++ {
		fmt.Println("for before, ", i)

		//defer func() {
		//	fmt.Println("defer", i)
		//}()

		defer func(i int) {
			fmt.Println("defer", i)
		}(i)

		fmt.Println("for after", i)
	}
	fmt.Println("main")
}
