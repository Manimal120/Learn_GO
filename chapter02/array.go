package main

import "fmt"

func main() {
	var names [3]string
	fmt.Printf("%#v\n", names)

	names = [...]string{"a", "b", "c"}

	a2 := [3][2]int{}
	fmt.Printf("%#v", a2)

	a2 = [3][2]int{{1, 2}, {1, 2}, {1, 2}}
	fmt.Printf("%#v\n", a2)

}
