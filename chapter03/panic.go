package main

import (
	"fmt"
)

func test1() (err error) {
	// recover 必须在defer内
	defer func() {
		fmt.Println("defer")
		//fmt.Println(recover())

		if e := recover(); e != nil {
			err = fmt.Errorf("%s", e)
		}

	}()

	fmt.Println("Before")
	panic("DIY")
	fmt.Println("After")

	return

}

func main() {
	err := test1()
	fmt.Println(err)
	fmt.Println(123)
}
