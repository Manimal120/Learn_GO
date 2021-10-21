package main

func quickSort(arr []int) []int {
	if arr == nil || len(arr) < 2 {
		return arr
	}

	pivot := arr[0]
	var left, right []int

	for _, ele := range arr[1:] {
		if ele <= pivot {
			left = append(left, ele)
		} else {
			right = append(right, ele)
		}
	}
	return append(quickSort(left), append([]int{pivot}, quickSort(right)...)...)
}

func quickSort1(arr []int) []int {
	if arr == nil || len(arr) < 2 {
		return arr
	}

	pivot := arr[0]
	var left, right []int

	for _, element := range arr[1:] {
		if element >= pivot {
			right = append(right, element)
		} else {
			left = append(left, element)
		}
	}

	return append(quickSort1(left), append([]int{pivot}, quickSort1(right)...)...)
}
