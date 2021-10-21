package main

import "fmt"

func main() {
	arr := []int{12312341, 3, 5, 45674575, 10, 12, 12, 1231, 12312341, 3, 5, 45674575, 10, 12, 12, 1231, 12312341, 3, 5, 45674575, 10, 12, 12, 1231}
	//selectSort1(arr)
	//fmt.Println(arr)
	//
	//bubbleSort1(arr)
	//fmt.Println(arr)
	//
	//insertSort1(arr)
	//fmt.Println(arr)

	fmt.Println(quickSort1(arr))
}
