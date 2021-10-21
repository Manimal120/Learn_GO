package main

import "fmt"

func main() {
	var scores map[string]float64

	fmt.Printf("%T, %#v\n", scores, scores)

	scores = map[string]float64{}
	fmt.Printf("%T, %#v\n", scores, scores)

	scores = make(map[string]float64) // == map[string]float64{}
	scores["22"] = 70.5

	fmt.Println(scores["22"])
	fmt.Println(scores["33"]) // 0 值

	// So --- >> Judge Existence
	v, ok := scores["33"]
	fmt.Println(v, ok)

	// Change
	scores["22"] = 123
	// Create
	scores["33"] = 100
	// Delete if not exist, return nothing
	delete(scores, "33")
	delete(scores, "44")

	// Traverse
	for k, v := range scores {
		fmt.Println(k, v)
	}
}
