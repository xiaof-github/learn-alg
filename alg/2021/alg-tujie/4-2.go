package main

import (
	"fmt"
)

func main() {
	var a =[]int{1,2,3,3,4,4,5,5,6}	
	c := count(a)
	fmt.Println(c)	
}

func count(arr []int) int{
	if len(arr) == 1 {
		return 1
	}
	c := count(arr[1:]) + 1
	return c
}