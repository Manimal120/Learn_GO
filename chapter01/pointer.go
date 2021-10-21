package main

import "fmt"

func main() {
	var (
		pointerInt    *int
		pointerString *string
	)
	fmt.Printf("%T %#v\n", pointerInt, pointerInt)
	fmt.Printf("%T %#v\n", pointerString, pointerString)

	age := 32

	fmt.Printf("%T %#v\n", &age, &age)



	//new
	pointerString = new(string)
	fmt.Println(pointerString)

	// pointer's pointer
	pp := &pointerString
	fmt.Printf("%T %#v\n", pp, **pp)




}
