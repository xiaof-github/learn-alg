package main

import "fmt"

// 实现LRU缓存，最近最少使用缓存

type LinkNode struct {
    Time int
    Key int
    Next *LinkNode    
}

type LRUCache struct {    
    cap int
    cur int
    cache map[int]int
    root *LinkNode
}

// 初始化链表
func initLink() *LinkNode {
    return new(LinkNode)
}

// 访问缓存时，更新其他key的时间
func (this *LRUCache)lruIncOtherKeyTime(key int) {
    i := 0    
    // for i;i < this.cap; i++ {
    //     if (this.hot.key != key && this.hot.exist){
    //         this.hot.time++
    //     }else if (this.hot.key == key){
    //         this.hot.time = 0
    //     }
    // }
}
// 删除节点
func deleteLinkNode(root *LinkNode, node *LinkNode) {
    var n *LinkNode
    last := root
    for n = root->Next; n!=null; last = n, n = n.Next {
        if (n == node){
            last.Next = n.Next
            n.Next = nil
            n = nil
            break;
        }
    }
}

// 加入新节点，并更新老节点时间值
func insertLinkNodeUpdate(root *LinkNode, key int) {
    n := root
    node := new(LinkNode)
    node.Key = key
    for n = root.Next;n!=null;n = n.Next {
        if (n != root){
            n.Time++
        }        
        if (n.Next == null){
            n.Next = node
            break;
        }
    }
    root.Key++
}

// 查询最大的节点
func findBigLinkNode(root *LinkNode) *LinkNode {
    n := root
    t := root.Next
    for n = root.Next;n!=null;n = n.Next {
        if (t.Val < n.Val){
            t = n
        }
    }
    return t
}


func Constructor(capacity int) LRUCache {
    lru := new(LRUCache)
    lru.cap =  capacity
    lru.cache = make(map[int]int, capacity)
    lru.root = initLink()
    return lru
}

func (this *LRUCache) Get(key int) int {

    val, ok := map[key]
    if (ok) {
        fmt.Printf("key %d exist", key)
        this.lruIncOtherKeyTime(key)
        return val
    }

    fmt.Printf("key %d not exist", key)
    return -1
}

func (this *LRUCache) Put(key int, value int) {
    if(cur < cap){
        map[key] = value        
        this.cur++
        insertLinkNode(this.root, key)
    }else{        
        fmt.Println("map before deletion", this.map)
        big := findBigLinkNode(this.root)
        delete(this.cache, big.key)
        deleteLinkNode(this.root, big)
        insertLinkNodeUpdate(root, key)        
        fmt.Println("map after deletion", this.map)
    }
    lruIncOtherKeyTime(key)
}

func main() {
    nums1 := []int{23,2,2,1,1,1,1,1,21,230,3,3,91}
    nums2 := []int{1,2,3,4,5,91,21,230,3,23,3,3,5}

    a := intersect(nums1, nums2)
    fmt.Printf("a:%+v", a)

}