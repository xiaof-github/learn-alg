/**
 * 解题思路：
 * 1.循环判断条件，mid1+mid2 = (len1+len2)/2
 * 2.每轮循环，取mid1的值和mid2的值，比较大小，小的一方，修改left,right值，计算新的mid，如果mid等于right，退出循环，否则继续循环
 * 3.循环结束后，如果是奇数，返回mid1和mid2中较大的一个数 nums[mid];如果是偶数，计算返回(mid1+mid2)/2
**/

package main

import (
	"fmt"
)

func main() {
    fmt.Println(findMedianSortedArrays([]int{1,3}, []int{4,5,6}))
}


func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    var k1,k2,len1,len2 int

    len1 = len(nums1)
    len2 = len(nums2)
    
    // 奇数返回一个数，偶数返回两个数的平均数
    if (len1 + len2) % 2 != 0 {
        k1 = (len1+len2)/2
        k2 = (len1+len2)/2        
    } else {
        k1 = (len1 + len2)/2
        k2 = (len1 + len2)/2+1
    }
    var mid1,mid2,left1,left2,right1,right2,num int
    right1 = len1-1
    right2 = len2-1    

    for mid1,mid2 = (left1 + right1)/2,(left2 + right2)/2 ;num != k1; {          
        
        // 滑动中轴，使mid1+mid2逼近k1
        if (mid1 + 1 + mid2 + 1) > k1 && nums1[mid1] < nums2[mid2] {
            right2 = mid2
            mid2 = (left2+right2)/2
        } else if (mid1 + 1 + mid2 + 1) <= k1 && nums1[mid1] < nums2[mid2] {
            left1 = mid1 + 1
            mid1 = (left1+right1)/2
        } else if (mid1 + 1 + mid2 + 1) > k1 && nums1[mid1] > nums2[mid2] {
            right1 = mid1
            mid1 = (left1+right1)/2
        } else if (mid1 + 1 + mid2 + 1) <= k1 && nums1[mid1] > nums2[mid2] {
            left2 = mid2 + 1
            mid2 = (left2+right2)/2
        }

        // 取当前的判断条件，num等于中位数
        num = mid1+1+mid2+1
        // fmt.Println("num:", num)

    }
    // fmt.Println("num:", num, "mid1：", mid1, "mid2：", mid2)
    if k1 == k2 {
        if nums1[mid1] > nums2[mid2] {
            return float64(nums1[mid1])
        } else {
            return float64(nums2[mid2])
        }        
    } else {
        return ( float64(nums1[mid1]) + float64(nums2[mid2]) )/2
    }
}