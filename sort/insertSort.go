package main

// 打扑克，新拿的牌和已有的牌对比，放到合适的地方

func insertSort(arr []int) []int {
	if arr == nil || len(arr) < 2 {
		return arr
	}
	for i := 1; i < len(arr); i++ { // 0 - i 范围有序
		for j := i - 1; j >= 0 && arr[j] > arr[j+1]; j-- { // 从i前的数往前走，看是否合适
			arr[j], arr[j+1] = arr[j+1], arr[j]
		}
	}
	return arr
}

func insertSort1(arr []int) []int {
	if arr == nil || len(arr) < 2 {
		return arr
	}

	for i := 1; i < len(arr); i++ {
		for j := i - 1; j >= 0 && arr[j] > arr[j+1]; j-- {
			arr[j], arr[j+1] = arr[j+1], arr[j]
		}
	}
	return arr
}
