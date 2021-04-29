package main

import (
	"fmt"
)

func main() {
	var array = []int{1,2,3,4,5,6,7,8,10,14}
	// var array = [8]int{6,7,8,1,2,3,4,5}
	// 拆分为两个数组，使用二分查找方法


	r := binarySearch(array,-1)
	fmt.Printf("result : %d", r)
}

func binarySearch(arr []int, dst int) int {
	left := 0
	right := len(arr)-1
	mid := (left + right)/2
	for left <= right {
		v := arr[mid]
		if (v == dst) {
			return mid
		} else if (v < dst) {
			left = mid + 1			
		} else {
			right = mid - 1
		}
		mid = (left+right)/2
	}
	return -1
}