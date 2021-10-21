package main

import "fmt"

func main() {
	var nilMap map[string]string
	fmt.Println(len(nilMap))

	fmt.Println(nilMap["kk"])

	// panic: assignment to entry in nil map
	nilMap["kk"] = "xx"

	// In conclusion, don't use nilSlice or nilMap
	// Use emptyMap instead
	// var emptyMap = map[string]string{}

}
