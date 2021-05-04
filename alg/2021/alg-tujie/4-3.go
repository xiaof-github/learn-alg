package main

import "fmt"

func main(){
	var a =[]int{1,2,3,3,44,4,5,5,6}
	
	c := big(a)
	fmt.Println(c)	
}

func big(arr []int) int {
	if (len(arr) == 1) {
		return arr[0]
	}			
	if (len(arr) == 2) {
		if (arr[0]>arr[1]){
			return arr[0]
		} else {
			return arr[1]
		}		
	}			
	c := big(arr[1:])
	if (arr[0]>c){
		return arr[0]
	} else {
		return c
	}		
}