package main

import "log"
func main(){
    var a = []int{45,145,245,32,5,2,69,239,12,40}
    log.Print("before: ",a)
    bubblesort(a)
	log.Print("after: ",a)
}
/**
 * 思路：从第一个数开始，依次和后面的数比较，通过交换的方法，将最大的数交换到最后的位置。再继续从头以上述方式调整最大数的位置，直到调整到第二个数。
 */
func bubblesort(arr []int){
	for i:=len(arr);i>0;i-- {
		for j:=0;j<i-1;j++ {
			if (arr[j] > arr[j+1]){
				swap(arr, j, j+1)
			}
		}
	}
}

func swap(arr []int, j int, k int){
	tmp := arr[j]
	arr[j] = arr[k]
	arr[k] = tmp	
}