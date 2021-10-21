package main

import "fmt"

// 匿名函数

//func main() {
//	a := func(a, b int) int {
//		return a + b
//	}
//
//	fmt.Println(a(1, 2))
//}

func main() {
	func() {
		fmt.Println(120)
	}()

	// ==
	a := func() {
		fmt.Println(110)
	}
	a()

	//
	func(i int) {
		fmt.Println(i)
	}(100)

	// 就近原则
	name, desc := 1, 2
	func(name int) {
		fmt.Println(name, desc)
		name, desc = 3, 4
	}(6)
	fmt.Println(name, desc)
}
