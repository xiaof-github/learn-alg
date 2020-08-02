package main

import "fmt"

// 实现LRU缓存，最近最少使用缓存

type LRUCache struct {    
    cap int
    cache map[int]int
}

func Constructor(capacity int) LRUCache {

}

func (this *LRUCache) Get(key int) int {

}

func (this *LRUCache) Put(key int, value int) {

}

func main() {
    nums1 := []int{23,2,2,1,1,1,1,1,21,230,3,3,91}
    nums2 := []int{1,2,3,4,5,91,21,230,3,23,3,3,5}

    a := intersect(nums1, nums2)
    fmt.Printf("a:%+v", a)

}