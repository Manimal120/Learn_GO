package main

import (
	"fmt"
	"strconv"
)

func main() {
	var (
		//intVal     = 65
		float64Val = 2.2
		stringVal  = "3.3"
	)

	vv, err := strconv.ParseFloat(stringVal, 64)
	fmt.Println(err, vv)
	fmt.Println(strconv.FormatFloat(float64Val, 'f', 10, 64))
}
