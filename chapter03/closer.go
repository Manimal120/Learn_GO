package main

import "fmt"

// 闭包： 方法内部的匿名函数引用了匿名函数之外的变量，导致GC无法回收

func addBase(base int) func(int) int {
	return func(i int) int {
		return i + base
	}
}

func main() {
	add1 := addBase(1)
	fmt.Println(add1(5))

}
