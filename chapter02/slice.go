package main

import "fmt"

func main() {

	//struct {
	//	*array
	//	length
	//	cap
	//}

	// 1 未初始化的空切片
	var names []string
	fmt.Printf("%T\n", names)
	fmt.Printf("%#v\n", names)

	// 2 初始化的空切片
	names1 := []string{}
	fmt.Printf("%#v\n", names1)

	names1 = []string{"Amy", "Bob"}
	fmt.Printf("%#v\n", names1)
	fmt.Printf("%q\n", names1)

	// 3 make
	names2 := make([]string, 5)
	fmt.Printf("%#v\n", names2)
	fmt.Printf("%q\n", names2)
	fmt.Println(cap(names2))

	// 4 traverse
	//for range
	//for i

	// 5 copy
	// a b 同长， 不同长时只替换掉能换掉的
	aSlice := []int{1, 2, 3}
	bSlice := []int{2, 2, 3}
	fmt.Println("~~~~~~~~~~~~~~~~~~~")
	copy(aSlice, bSlice) // (des, src)
	fmt.Printf("%#v,%#v\n", aSlice, bSlice)

	// 6 sub_slice
	nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	//start <= end <= cap  new_cap = cap - start
	numChild := nums[3:4]
	numChild = append(numChild, 100) // change Index 4
	fmt.Printf("%#v, %#v\n", nums, numChild)

	numChild2 := nums[9:10]
	numChild2 = append(numChild2, 100)        // change Index 4, change on nums Index 4
	fmt.Printf("%#v, %#v\n", nums, numChild2) // change Index 10, no change on nums

	fmt.Printf("%T, %#v\n", numChild, numChild)
	fmt.Println(cap(numChild))

	// 7 sub_slice2
	// start <= end <= max <= cap, new_cap = max - start
	nums1 := []int{0, 1, 2, 3, 4}
	numChild3 := nums1[3:4:4]
	numChild3 = append(numChild3, 100)
	fmt.Printf("%#v, %#v\n", nums1, numChild3) // no change on nums, cuz cap is doubled to a new array

	// 8 slice_ptr
	nums3 := make([]int, 3, 3)
	nums3 = []int{0, 1, 2}

	fmt.Printf("%p\n", &nums3)
	fmt.Println(cap(nums3))
	nums3 = append(nums3, 3)
	fmt.Printf("%p\n", &nums3)
	fmt.Println(cap(nums3))

	// 9 变量也有地址
	months := []string{1: "January"}
	fmt.Printf("%p,\n", &months)    //0xc0000ae040,
	fmt.Printf("%p,\n", months)     //0xc0000ae060,
	fmt.Printf("%p,\n", &months[0]) //0xc0000ae060,

	// 10 扩容新的切片地址不变？ 连续内存占用会变，否则不变

	// 11 移除中间元素

	nums4 := []int{2, 3, 4}
	// 移除3
	copy(nums4[1:], nums4[2:])
	nums4 = nums4[:len(nums4)-1]
	fmt.Println(nums4)
}
