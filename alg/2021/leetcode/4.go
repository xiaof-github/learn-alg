package main

import (
	"fmt"
)

func main() {
    fmt.Println(findMedianSortedArrays([]int{2,7}, []int{4,5}))
}

/**
 * 解题思路：
 * 1.循环判断条件，mid1+mid2 = (len1+len2)/2
 * 2.每轮循环，取mid1的值和mid2的值，比较大小，小的一方，修改left,right值，计算新的mid，如果mid等于right，退出循环，否则继续循环
 * 3.循环结束后，如果是奇数，返回mid1和mid2中较大的一个数 nums[mid];如果是偶数，计算返回(mid1+mid2)/2
**/

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    var plen,len1,len2 int

    len1 = len(nums1)
    len2 = len(nums2)
    
    // 奇数返回一个数，偶数返回两个数的平均数
    if (len1 + len2) % 2 != 0 {
        plen = (len1+len2)/2 + 1        
    } else {
        plen = (len1 + len2)/2 + 1        
    }
    var mid1,mid2,left1,left2,right1,right2,num int
    var pos1, pos2 int
    pos1 = -1
    pos2 = -1
    right1 = len1-1
    right2 = len2-1    

    for mid1,mid2 = (left1 + right1)/2,(left2 + right2)/2 ;num != plen; {          

        if mid1 + mid2 + 2 > plen {
            if nums1[mid1] < nums2[mid2] {
                right2  = mid2
                mid2 = (left2 + right2)/2
            } else {
                right1 = mid1
                mid1 = (left1 + right1)/2
            }
            continue
        }
        
        // 取出两个数组中值较小的数组的左边数据
        if nums1[mid1] < nums2[mid2] {
            pos1 = mid1
            left1 = mid1+1
            // 超过数组长度
            if left1 == len1 {
                pos2 = plen - mid1 - 2
                break
            }
            mid1 = (left1 + right1)/2
            // 设置mid2值为待去掉长度的1/2            
            mid2 = mid2 - (mid1-pos1)
            if mid2 < 0 {
                mid2 = 0
            }
            right2 = mid2*2-left2
        } else if nums1[mid1] > nums2[mid2] {
            pos2 = mid2
            left2 = mid2+1
            // 超过数组长度
            if left2 == len2 {
                pos1 = plen - mid2 -2
                break
            }
            mid2 = (left2 + right2) /2
            mid1 = mid1 - (mid2-pos2)
            if mid1 < 0 {
                mid1 = 0
            }
            right1 = mid1*2 - left1
        } else {
            if mid1 + mid2 + 2 == plen {
                pos1 = mid1
                pos2 = mid2
                break
            } else {
                pos1 = mid1
                pos2 = mid2
                left1 = mid1+1
                if left1 == len1 {
                    pos2 = plen - mid1 - 2
                    break
                }
                left2 = mid2+1
                if left2 == len2 {
                    pos1 = plen - mid2 -2
                    break
                }
            }
            // 重新设置left1,right1,left2,right2,更新num值，进入下一轮计算,mid1<len1-1,mid2<len2-1
            mid1 = (left1 + right1)/2
            mid2 = (left2 + right2)/2
        }
        
        num = pos1+1+pos2+1
    }
    fmt.Println("num:", num, "mid1：", mid1, "mid2：", mid2, "pos1：", pos1, "pos2：", pos2)
    if pos1 == -1 {
        if (len1 + len2) % 2 != 0 {
            return float64(nums2[pos2])
        } else {
            return ( float64(nums2[pos2]) + float64(nums2[pos2-1]) )/2
        }
    }
    if pos2 == -1 {
        if (len1 + len2) % 2 != 0 {
            return float64(nums1[pos1])
        } else {
            return ( float64(nums1[pos1]) + float64(nums1[pos1-1]) )/2
        }
    }
    if (len1 + len2) % 2 != 0 {
        if nums1[pos1] > nums2[pos2] {
            return float64(nums1[pos1])
        } else {
            return float64(nums2[pos2])
        }        
    } else {
        var a int
        if nums1[pos1] > nums2[pos2] {
            a = nums1[pos1]
            if pos1-1 >= 0 && nums1[pos1-1] > nums2[pos2] {
                return (float64(a) + float64(nums1[pos1-1])) /2
            } else {
                return (float64(a) + float64(nums1[pos2])) /2
            }
        }
        a = nums2[pos2]
        if pos2-1 >= 0 && nums2[pos2-1] > nums1[pos1] {
            return (float64(a) + float64(nums2[pos2-1]))/2
        } else {
            return (float64(a) + float64(nums1[pos1]))/2
        }        
    }
    
}