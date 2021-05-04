package main

import (
	"fmt"
)

func main() {
	var arr = []int{10,931,3,2,1,443,2,6,34}
	fmt.Println("before arr:", arr)
	selectionSort(arr)
	fmt.Println("after arr:", arr)
}

func selectionSort(arr []int){
	for i:=len(arr)-1;i>=0;i-- {
		for j:=0;j<i;j++ {
			if (arr[j] > arr[j+1]) {
				swap(arr, j, j+1)
			}
		}
	}
}

func swap(arr []int, i int, j int){
	tmp := arr[i]
	arr[i] = arr[j]
	arr[j] = tmp
}