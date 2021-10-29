package main

import (
	"fmt"
)

func main() {
	arr := []int{-4, -2, 0, 2, 3, 221, 23, 45, 63, 38, 2, 32, 5, 6}

	fmt.Println(threeSum(arr))
}

/**
 * 排序。双指针移动。
 */
func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return make([][]int, 0)
	}
	rst := make([][]int, 0)

	fastSort(nums, 0, len(nums)-1)

	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			break
		}
		if i >= 1 && nums[i] == nums[i-1] {
			continue
		}
		left := i + 1
		right := len(nums) - 1
		lastLeft := left
		lastRight := right
		for left < right {
			if left != lastLeft && nums[left] == nums[lastLeft] {
				left++
				continue
			}
			if right != lastRight && nums[right] == nums[lastRight] {
				right--
				continue
			}

			// 满足条件时，left和right同时滑动
			if nums[i]+nums[left]+nums[right] == 0 {
				add := make([]int, 0)
				add = append(add, nums[i], nums[left], nums[right])
				rst = append(rst, add)
				lastLeft = left
				lastRight = right
				left++
				right--
			} else if nums[i]+nums[left]+nums[right] > 0 {
				// 右滑
				lastRight = right
				right--

			} else {
				// 左滑
				lastLeft = left
				left++
			}
		}
	}
	return rst
}

func fastSort(arr []int, left int, right int) {
	if left >= right {
		return
	}
	l1 := left
	r1 := right
	middle := arr[left]

	for left != right {

		for right > left && arr[right] >= middle {
			right--
		}
		arr[left] = arr[right]
		for left < right && arr[left] <= middle {
			left++
		}
		arr[right] = arr[left]
	}
	arr[left] = middle
	fastSort(arr, l1, left-1)
	fastSort(arr, left+1, r1)

}
