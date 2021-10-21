package main

import (
	"fmt"
	"strings"
)

func main() {
	// 空白符分割的切片统计
	fmt.Printf("%#v\n", strings.Fields("aasd b \t c \t c \f vb"))

	// 开头结尾
	fmt.Println(strings.HasPrefix("abc", "a"))
	fmt.Println(strings.HasPrefix("Abc", "b"))

	// 字符串出现的位置
	fmt.Println(strings.Index("acsadasa", "sa"))
	fmt.Println(strings.Index("acsadasa", "x"))

	// Link Split
	fmt.Println(strings.Join([]string{"ab", "ac", "ed"}, "--"))
	fmt.Println(strings.Split("ab--ac--ed", "-"))
	fmt.Println(strings.SplitN("ab--ac--ed", "--", 2))

	// Repeat
	fmt.Println(strings.Repeat("8", 5))

}
