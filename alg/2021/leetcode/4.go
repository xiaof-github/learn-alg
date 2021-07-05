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
    fmt.Println(findMedianSortedArrays([]int{1,23,3}, []int{4,5,6}))
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
        k1 = (len1 + len2)/2 - 1
        k2 = (len1 + len2)/2
    }
    var mid1,mid2,left1,left2,right1,right2 int
    right1 = len1-1
    right2 = len2-1    
    for mid1,mid2 = (left1 + right1)/2,(left2 + right2)/2 ;(mid1+mid2) != k1; {                
        if (nums1[mid1] < nums2[mid2]) {
            mid1 = (left1 + right1)/2            
            if (mid1+mid2)/2 > k1 {
                right1 = mid1
            } else {
                left1 = mid1
            }            
            
        }else {
            mid2 = (left2 + right2)/2
            if (mid1+mid2)/2 > k1 {
                right2 = mid2
            } else {
                left2 = mid2
            }            
        }

        
        
    }

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