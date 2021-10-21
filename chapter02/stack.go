package main

func main() {
	// stack 堆栈 FILO

	var stack []string

	// push
	stack = append(stack, "a")
	stack = append(stack, "b")
	stack = append(stack, "c")

	// pop
	stack = stack[:len(stack)-1]

}
