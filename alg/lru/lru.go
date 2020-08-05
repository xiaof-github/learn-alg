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
    var n *LinkNode    
    for n = this.root.Next; n!=nil; n = n.Next {
        if (n.Key == key){
            n.Time = 0            
            continue
        }
        n.Time++
    }
}
// 删除节点
func deleteLinkNode(root *LinkNode, node *LinkNode) {
    var n *LinkNode
    last := root
    for n = root.Next; n!=nil; last, n =n, n.Next {
        if (n == node){
            last.Next = n.Next
            n.Next = nil
            n = nil
            root.Key--
            break
        }
    }
}

// 加入新节点
func insertLinkNodeUpdate(root *LinkNode, key int) {
    n := root
    node := new(LinkNode)
    node.Key = key
    for n = root;n!=nil;n = n.Next {                 
        if (n.Next == nil){
            n.Next = node
            break
        }
    }
    root.Key++
}

// 查询最大的节点
func findBigLinkNode(root *LinkNode) *LinkNode {
    n := root
    t := root.Next
    for n = root.Next;n!=nil;n = n.Next {
        if (t.Time < n.Time){
            t = n
        }
    }
    return t
}


func Constructor(capacity int) *LRUCache {
    lru := new(LRUCache)
    lru.cap =  capacity
    lru.cache = make(map[int]int, capacity)
    lru.root = initLink()
    return lru
}

func (this *LRUCache) Get(key int) int {

    val, ok := this.cache[key]
    if (ok) {
        fmt.Printf("key %d exist\n", key)
        this.lruIncOtherKeyTime(key)
        return val
    }

    fmt.Printf("key %d not exist\n", key)
    return -1
}

func (this *LRUCache) Put(key int, value int) {
    if(this.cur < this.cap){
        this.cache[key] = value        
        this.cur++
        insertLinkNodeUpdate(this.root, key)
    }else{        
        fmt.Println("map before deletion", this.cache)
        big := findBigLinkNode(this.root)
        fmt.Printf("big: %+v", big)
        delete(this.cache, big.Key)
        deleteLinkNode(this.root, big)
        insertLinkNodeUpdate(this.root, key)        
        this.cache[key] = value
        fmt.Println("map after deletion", this.cache)
    }
    this.lruIncOtherKeyTime(key)
}

func main() {    
    cache := Constructor(2)
    cache.Put(3, 10)
    cache.Put(1, 5)
    cache.Get(1)
    cache.Put(7, 7)    
    cache.Put(8, 8)
    cache.Get(3)
    cache.Put(9, 9)
    cache.Get(8)

    

}