package main

import (
	"log"
	"time"
)
func main(){
    var a = []int{45,145,245,32,5,2,69,239,12,40}
    log.Print("before: ", a)
	start := time.Now().UnixNano()
    heapsort(a)
	end := time.Now().UnixNano()
    log.Print("after: ", a)
	Milliseconds:= float64((end - start) / 1e6)
	log.Print("start_time, end_time, cost: ", start, end, Milliseconds)
}

func heapsort(arr []int) {
	len := len(arr)
	buildMaxHeap(arr)
	swap(arr, 0, len-1)
	for i := len-1;i>0;i-- {
		heapify(arr[0:i], 1)
		swap(arr[0:i], 0, i-1)
	}	
}

func buildMaxHeap(arr []int) {
	lenH := len(arr)
	for i:=lenH/2;i>0;i-- {
		heapify(arr,i)
	}
}

func swap(arr []int, i int, j int) {
	temp := arr[i]
	arr[i] = arr[j]
	arr[j] = temp
}
/*
 * 指定位置的节点按最大堆调整
 *
 */
func heapify(arr []int, heap int) {

	if (heap>len(arr) || heap*2>len(arr)){
		return
	}
	// 比较左子节点 
	if (arr[heap-1] < arr[2*heap-1]) {
		swap(arr, heap-1, 2*heap-1)
	}
	// 比较右子节点
	if (heap*2+1>len(arr)) {
		return
	}
	if (arr[heap-1] < arr[2*heap]) {
		swap(arr, heap-1, 2*heap)
	}

	// 子节点
	heapify(arr, 2*heap)
	heapify(arr, 2*heap+1)
}