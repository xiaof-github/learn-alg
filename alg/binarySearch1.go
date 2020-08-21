package main

import (
	"fmt"
)

func main() {
	var array = [8]int{1,2,3,4,5,6,7,8}
	// var array = [8]int{6,7,8,1,2,3,4,5}
	// 拆分为两个数组，使用二分查找方法


	r := binarySearch(array, 8, 6)
	fmt.Printf("result : %d", r)
}

func binarySearch(array [8]int, size int, target int) int {
	var left int 
	var right int  
	right = size - 1 
	var middle int 
	middle = (left + right)/2 

    for (left < right) {
		middle = (right + left) / 2;		
        if (array[middle] == target) {
            return middle
        } else if (array[middle] < target) {
			left = middle + 1
        } else if (array[middle] > target) {
            right = middle - 1
        }
    }
    return -1;
}