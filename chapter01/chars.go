package main

import "fmt"

func main() {
	letters := "abcdefg"
	for i := 0; i < len(letters); i++ {
		fmt.Printf("%c\n", letters[i])
	}

	s := "我爱中国"
	for _, v := range s {
		fmt.Printf("%q", v)
	}
}
