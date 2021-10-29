package main

import (
	"fmt"
)

func main() {
    arr := []int{2,3,3,4,8,9,9,1}
    fmt.Println(maxArea(arr))
}

func maxArea(height []int) int {
    max := 0
    rst := 0
    a := 0
    j := len(height) - 1
    for i:=0;;{        
        a = height[i]
        if a > height[j] {
            a = height[j]
        }

        rst = a*(j-i)
        if max < rst {
            max = rst
        }
        if height[i] < height[j] {
            i++
        } else {
            j--
        }
        if i==j {
            break
        }
    }   

    return max
}