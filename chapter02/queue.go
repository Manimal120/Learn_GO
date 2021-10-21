package main

import "fmt"

// queue  FIFO

func main() {
	var queue []string

	// push - append
	queue = append(queue, "a", "b")
	queue = append(queue, "c")

	// pop
	x := queue[0]
	queue = queue[1:]
	fmt.Println(x)

	x = queue[0]
	queue = queue[1:]
	fmt.Println(x, "1")

	x = queue[0]
	queue = queue[1:]
	fmt.Println(x, "2")

}
