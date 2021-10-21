package main

func bubbleSort(arr []int) []int {
	if arr == nil || len(arr) < 2 {
		return arr
	}
	for i := 0; i < len(arr)-1; i++ {
		for j := i; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				//swap(&arr[j], &arr[j+1])
			}
		}
	}
	return arr
}

func bubbleSort1(arr []int) []int {
	if arr == nil || len(arr) < 2 {
		return arr
	}
	for i := 0; i < len(arr); i++ {
		for j := i; j < len(arr)-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}
