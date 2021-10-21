package main

import (
	"fmt"
	"strconv"
	"unicode/utf8"
)

func main() {
	ascii := "abc我爱到底"
	fmt.Println([]byte(ascii)) // bin
	fmt.Println([]rune(ascii)) // unicode

	fmt.Println(len(ascii))
	fmt.Println(len([]rune(ascii)))

	fmt.Println(utf8.RuneCountInString(ascii))
	fmt.Println(string([]byte(ascii)))
	fmt.Println(string([]rune(ascii)))

	fmt.Println(strconv.Itoa('a'))
	fmt.Println(strconv.Atoi("97"))
	fmt.Println(strconv.FormatFloat(3.1314, 'f', 6, 64))
	fmt.Println(strconv.ParseFloat("3.1313131", 64))
	fmt.Println(strconv.FormatBool(true))
	fmt.Println(strconv.ParseBool("true"))

	fmt.Println(strconv.ParseInt("6", 10, 8))
	fmt.Println(strconv.FormatInt(16, 2))
}
