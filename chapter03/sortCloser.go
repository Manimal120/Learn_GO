package main

import (
	"fmt"
	"sort"
)

func main() {
	stats := [][]int{{'A', 3}, {'B', 2}, {'C', 3}, {'D', 3}}

	sort.Slice(stats, func(i, j int) bool {
		return stats[i][1] < stats[j][1]
	})

	fmt.Println(stats)

	//sort.SliceStable() 次数相同时，按照出现的先后顺序开始

}
