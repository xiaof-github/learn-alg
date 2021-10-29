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
    a := 0
    for i:=0;i<len(height);i++{        
        for j:=i+1;j<len(height);j++ {
            a = height[i]
            if height[i] > height[j] {
                a = height[j]
            }
            rst := a * (j-i)
            if max < rst {
                max = rst
            }
        }
    }   

    return max
}