package main

import (
	"errors"
	"fmt"
	"strconv"
)

func div(n1, n2 int) (int, error) {
	if n2 == 0 {
		return -1, errors.New("wrong")
	}
	return n1 / n2, nil
}

func main() {
	value, err := strconv.Atoi("XXX")
	fmt.Println(value, err)

	e := fmt.Errorf("DIY Error")
	fmt.Println(e)

	e2 := errors.New("DIY Error 2")
	fmt.Println(e2)

	if v, err := div(0, 0); err == nil {
		fmt.Println(v)
	} else {
		fmt.Println(err)
	}

}
