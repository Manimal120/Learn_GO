package main

import "fmt"

func test() (rt string) {
	// defer 函数调用
	// stack ORDER
	defer func() {
		fmt.Println("defer1")
		rt = "qzd" // 在延迟执行中不要修改返回值
	}()

	defer func() {
		fmt.Println("defer2")
	}()

	defer func() {
		fmt.Println("defer3")
	}()

	fmt.Println("main")
	rt = "test"
	return
}

func test2(n1, n2 int) {
	defer func() {
		// 函数体内不管是否发生错误都会执行
		fmt.Println("New defer")
	}()
	fmt.Println("Before")
	fmt.Println(n1 / n2)
	fmt.Println("After")

}

func main() {
	//fmt.Println(test())
	test2(1, 0)
}
