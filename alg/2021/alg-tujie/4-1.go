package main

import "fmt"

func main(){
	// var a =[]int{1,2,3,3,4,4,5,5,6}
	var a =[]int{1}
	c := sum(a)
	fmt.Println(c)
	fmt.Println(a[1:])
}

func sum(arr []int) int {
	if (len(arr) == 0) {
		return 0
	}			
	
	return arr[0] + sum(arr[1:])
}