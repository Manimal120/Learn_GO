package main

import (
	"fmt"
	"strings"
)

func main() {
	var text string

	fmt.Println("Pls Input: ")
	fmt.Scan(&text)

	for _, todo := range todos {
		if strings.Contains(todo[name], text) {
			fmt.Println(true)
		}
	}

	//methods := map[string]func(){
	//	"add":   add,
	//	"query": query,
	//}
	//
	//for {
	//	t := input("Pls Input.....")
	//	if test == "exit" {
	//		break
	//	}
	//	method, ok := methods[t]
	//	if ok {
	//		method()
	//	} else {
	//		fmt.Println("Wrong Order!")
	//	}
	//}
}
