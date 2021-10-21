package main

import "fmt"

func main() {
	// 作用域 定义标识符可以作用的范围
	// use {}
	// 里可以用外，外不能用里
	outer := 1
	{
		inner := 2
		fmt.Println(outer)
		fmt.Println(inner)
	}

}
