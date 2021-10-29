package main

import (
	"fmt"
)

func main() {
    fmt.Println(findMedianSortedArrays([]int{2, 2, 4, 4}, []int{2, 2, 4, 4}))
}

/**
 * 解题思路：
 * 
 * 1.每轮循环，正常情况下排除掉k/2个数。
 * 2.单独处理k为1的情况，比较两个数组中当前位置的数。
 * 3.如果遇到数组1的边界，则数组1加上k/2会越界，只能取到边界的数比较。
**/

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    var plen, len1, len2 int

    len1 = len(nums1)
    len2 = len(nums2)

    if len1 == 0 && len2 == 0 {
        return float64(0)
    }

    if len1 == 0 {
        if len2%2 != 0 {
            return float64(nums2[len2/2])
        } else {
            return (float64(nums2[len2/2]) + float64(nums2[len2/2-1])) / 2
        }
    }

    if len2 == 0 {
        if len1%2 != 0 {
            return float64(nums1[len1/2])
        } else {
            return (float64(nums1[len1/2]) + float64(nums1[len1/2-1])) / 2
        }
    }

    // 奇数返回一个数，偶数返回两个数的平均数
    if (len1+len2)%2 != 0 {
        plen = (len1+len2)/2 + 1
        return float64(getKelement(nums1, nums2, plen))
    } else {
        plen = (len1+len2)/2
        return (float64(getKelement(nums1, nums2, plen)) + 
        float64(getKelement(nums1, nums2, plen+1)))/2.0
    }
}
/**
 * 找到第k个最小的数
 */
func getKelement(nums1 []int, nums2 []int, k int) int {   
    var newStart1, newStart2, start1, start2 int
    len1 := len(nums1)
    len2 := len(nums2)
    
        
    for {
        // 防止数组越界
        if newStart1 == len1 {
            
            return nums2[newStart2 + k - 1]
        }

        if newStart2 == len2 {
            
            return nums1[newStart1 + k -1]
        }

        if k == 1 {
            if nums1[newStart1] <= nums2[newStart2] {
                
                return nums1[newStart1]
            } else {
                
                return nums2[newStart2]
            }            
        }        
        start1 = min(newStart1 + k/2, len1)-1
        
        start2 = min(newStart2+ k/2, len2)-1
        // fmt.Println("start1: ",start1)
        // fmt.Println("start2: ",start2)
        if nums1[start1] <= nums2[start2] {                        
            k -= (start1-newStart1 + 1)            
            newStart1 = start1+1
        } else {
            k -= (start2-newStart2 + 1)            
            newStart2 = start2+1
        }
        
    }
}

func min(a int, b int) int {
    if a < b {
        return a
    }
    return b
}
