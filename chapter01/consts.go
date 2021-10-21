package main

import "fmt"

func main() {
	const NAME string = "AKA"
	const (
		C3 = "licence"
		C4 = "dog"
		C5 // "dog"
	)
	const (
		E4 int = iota
		E5
		E6
	)
	fmt.Println(NAME)
}
