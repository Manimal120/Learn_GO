package main

import "fmt"

func main() {
	var nilSlice []int
	var emptySlice = []int{}

	fmt.Printf("%T, %#v\n", nilSlice, nilSlice)
	fmt.Printf("%T, %#v\n", emptySlice, emptySlice)
}
