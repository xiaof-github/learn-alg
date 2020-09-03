package main

import (
	"fmt"
)

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

var r *TreeNode

// 求二叉树层序遍历
func main() {
	// nums1 := []int{23,2,2,1,1,1,1,1,21,230,3,3,91}
	// nums2 := []int{1,2,3,4,5,91,21,230,3,23,3,3,5}

	NodeS1 := make([]*TreeNode, 0, 100)
	NodeS1 = append(NodeS1, new(TreeNode))
	fmt.Printf("NodeS1: %+v\n", NodeS1[0])


	a := make([][]int, 0, 100)  
	var c []int

	for i:=0; i < 100; i++ {
		c = []int{}
		a = append(a, c)
	}

	b := 1
	r = new(TreeNode)
	r.Val = -1
	insertNode(r, b)
	insertNode(r, 2)
	insertNode(r, 3)
	insertNode(r, 4)
	insertNode(r, 4)
	insertNode(r, 4)
	insertNode(r, 4)
	insertNode(r, 4)

	rst := levelOrder1(r)
	fmt.Printf("rst: %+v\n", rst)
}
// 以满二叉树方式插入节点
func insertNode(root *TreeNode, Val int) {
	NodeS := make([]*TreeNode, 1, 100)
	NodeS1 := make([]*TreeNode, 0, 200)
	if (root.Val == -1) {
		root.Val = Val
		return;
	}

	NodeS[0] = root
	done := false	

	for (!done) {			
		for _, v := range NodeS{
			if ( v.Left == nil) {
				v.Left = new(TreeNode)
				v.Left.Val = Val
				fmt.Printf("add node val: %d, NodeS1, len:%d left \n", Val, len(NodeS1))
				done = true
				break
			} else if (v.Right == nil) {
				v.Right = new(TreeNode)
				v.Right.Val = Val
				fmt.Printf("add node val: %d, NodeS1, len:%d right \n", Val, len(NodeS1))
				done = true
				break
			}
			NodeS1 = append(NodeS1, v.Left, v.Right)
		}
		if (!done) {
			// 数的上一层节点满了，遍历下一层数节点
			// 扩容NodeS，防止拷贝丢失数据
			NodeS = make([]*TreeNode, len(NodeS1), 100)
			copy(NodeS, NodeS1)
			NodeS1 = NodeS1[:0]
			fmt.Printf("NodeS, len: %d cap: %d\n", len(NodeS), cap(NodeS))
			fmt.Printf("NodeS1, len: %d cap: %d\n", len(NodeS1), cap(NodeS1))
		}
	}
	
}


// 层序遍历
func levelOrder(root *TreeNode) [][]int {
	NodeS := make([]*TreeNode, 1, 100)
	NodeS1 := make([]*TreeNode, 0, 100)
	rst := make([][]int, 0, 100)	
	var list []int

	if (root.Val == -1) {
		return rst
	}


	NodeS[0] = root
	done := false
	for (!done){
		list = make([]int, 0, 100)
		for _, v := range NodeS {
			if (v != nil){
				fmt.Printf("value: %d\n", v.Val)
				list = append(list, v.Val)
				if(v.Left != nil) {
					NodeS1 = append(NodeS1,v.Left)
				}
				if (v.Right != nil) {
					NodeS1 = append(NodeS1,v.Right)
				}				
			}			
		}
		rst = append(rst, list)
		if (len(NodeS1) == 0) {
			break
		}
		NodeS = NodeS[:0]
		fmt.Println(NodeS)
		NodeS = append(NodeS, NodeS1...)
		fmt.Println("append ", NodeS)
			
		NodeS1 = NodeS1[:0]		
	}
	
	return rst
}


// 层序遍历优化
func levelOrder1(root *TreeNode) [][]int {
	var count int
	var list []int
	var node *TreeNode

	if (root == nil) {
		return nil
	}

	queue := make([]*TreeNode, 1, 100)
	queue[0] = root

	rst := make([][]int, 0, 100)

	for len(queue) > 0 {
		count = len(queue)
		list = make([]int, 0, 100)
		for count > 0 {
			//队头的值加入列表
			node = queue[0]
			list = append(list, node.Val)
			//从队头取子节点入队
			if (node.Left != nil) {
				queue = append(queue, node.Left)
			}
			if (node.Right != nil){
				queue = append(queue, node.Right)
			}			
			//队头出队
			queue = queue[1:]
			count--
		}
		rst = append(rst, list)		
	}

	return rst
}