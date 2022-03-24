package main

import "fmt"

func findBiggestNumber(nums []int) int {
	max := nums[0]
	for _, num := range nums {
		if num > max {
			max = num
		}
	}

	return max
}

func test1() {
	nums1 := []int{1, 2, 3, 4, 5, 6, 7}
	nums2 := []int{10, 22, 33, 41, 544, 6123, 7123}
	nums3 := []int{12313123, 221313, 3342, 454, 512, 6123, 732}

	max1 := nums1[0]
	for _, num := range nums1 {
		if num > max1 {
			max1 = num
		}
	}
	fmt.Println(max1)

	max2 := nums2[0]
	for _, num := range nums2 {
		if num > max2 {
			max2 = num
		}
	}
	fmt.Println(max2)

	max3 := nums1[0]
	for _, num := range nums3 {
		if num > max3 {
			max3 = num
		}
	}
	fmt.Println(max3)
}

func test2() {
	nums1 := []int{1, 2, 3, 4, 5, 6, 7}
	nums2 := []int{10, 22, 33, 41, 544, 6123, 7123}
	nums3 := []int{12313123, 221313, 3342, 454, 512, 6123, 732}

	max1 := findBiggestNumber(nums1)
	max2 := findBiggestNumber(nums2)
	max3 := findBiggestNumber(nums3)

	fmt.Println(max1)
	fmt.Println(max2)
	fmt.Println(max3)
}
