package main

import (
	"fmt"
	"sort"
)

func main() {
	// Int sort
	nums := []int{12312341, 3, 5, 9, 10, 12, 12, 1231}
	//sort.Ints(nums)
	//nums = selectSort(nums)
	//nums = bubbleSort(nums)
	nums = insertSort(nums)
	fmt.Println(nums)
	// Desc
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	fmt.Println(nums)

	// String sort
	names := []string{"abc", "xyz", "mn"}
	sort.Strings(names)
	fmt.Println(names)

	// Sort Search 二分查找
	nums = []int{1, 3, 5, 2, 12, 1231, 1231234, 9, 10, 1}
	fmt.Println(nums[sort.SearchInts(nums, 5)] == 8)
	fmt.Println(nums[sort.SearchInts(nums, 8)] == 8)

}

func selectSort(arr []int) []int {
	if arr == nil || len(arr) < 2 {
		return arr
	}
	for i := 0; i < len(arr)-1; i++ {
		minIndex := i
		for j := i + 1; j < len(arr); j++ {
			if arr[minIndex] > arr[j] {
				minIndex = j
			}
		}
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
	return arr
}

func bubbleSort(arr []int) []int {
	if arr == nil || len(arr) < 2 {
		return arr
	}
	for i := 0; i < len(arr)-1; i++ {
		for j := i; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				//arr[j], arr[j+1] = arr[j+1], arr[j]
				swap(&arr[j], &arr[j+1])
			}
		}
	}
	return arr
}

func insertSort(arr []int) []int {
	if arr == nil || len(arr) < 2 {
		return arr
	}
	for i := 1; i < len(arr); i++ {
		for j := i - 1; j >= 0 && arr[j] > arr[j+1]; j-- {
			if arr[j] > arr[j+1] {
				swap(&arr[j], &arr[j+1])
			}
		}
	}
	return arr
}

func swap(a, b *int) {
	*a = *a ^ *b
	*b = *a ^ *b
	*a = *a ^ *b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
