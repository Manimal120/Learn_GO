package main

import "fmt"

func main() {
	var y string

	fmt.Scan(&y)

	switch y {
	case "yes", "y", "Y":
		fmt.Println(10)
	}
}
