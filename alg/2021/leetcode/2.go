package main

import (
    "fmt"
)

type ListNode struct {
    Val int
    Next *ListNode
}

func main() {
	var array1 = [8]int{1,2,3,4,5,6,7,8}
	var array2 = [8]int{3,3,3,4,1,1,3,2}	
    root1 := new(ListNode)
    root1.Val = array1[0]
    root2 := new(ListNode)
    root2.Val = array2[0]
    for i:=1;i<len(array1);i++ {
        insert(array1[i], root1)
        insert(array2[i], root2)
    }
    
    rst := addTwoNumbers(root1,root2)
    for cur := rst; cur!= nil ;cur=cur.Next {
        fmt.Println(cur.Val)
    }
}

func insert(val int, root *ListNode) {
	
	if(root == nil) {
		return
	}

	p := root
	node := new(ListNode)

	for ;p != nil ; p = p.Next{
		if p.Next == nil {
            p.Next = node
            node.Val = val
            break;
        }			
	}
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    var l *ListNode = new(ListNode)
    var sum int
    var p1,p2,cur *ListNode = l1,l2,l
    
    if (l1 == nil && l2 == nil){
        return nil
    }
    
    l.Next = nil    

    var add int
    for (add > 0 || p1 != nil || p2 != nil){
        if (p1 != nil) {                                 
            sum += p1.Val
            p1 = p1.Next                            
        }
        if (p2 != nil){
            sum += p2.Val            
            p2 = p2.Next            
        }
        if(add != 0){
            sum+=add
        }
        cur.Val = sum % 10
        if (sum >= 10) {            
            add = 1
        }else{            
            add = 0
        }
        if (add > 0 || p1 != nil || p2 != nil){
            cur.Next = new(ListNode)
            cur = cur.Next
        }        
        sum = 0                
    }
    return l
    
}
