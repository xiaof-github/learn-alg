package main

type ListNode struct {
    Val int
    Next *ListNode
}

func main() {
	var array = [8]int{1,2,3,4,5,6,7,8}
	// var array = [8]int{6,7,8,1,2,3,4,5}
	// 拆分为两个数组，使用二分查找方法

	insert


}

fun insert(val int, root *ListNode) int {
	var a int 
	
	if(root == nil) {
		return -1
	}

	p := root
	node := new(ListNode)

	for p;p != nil ; p = p.Next{
		if (p.Next == nil)
			break;
	}

	p.Next = ListNode
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
