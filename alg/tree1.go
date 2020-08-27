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

	a := 1
	r = new(TreeNode)
	r.Val = -1
	insertNode(r, a)
	insertNode(r, 2)
	insertNode(r, 3)
	insertNode(r, 4)
	insertNode(r, 4)
	insertNode(r, 4)
	insertNode(r, 4)
	insertNode(r, 4)

}
// 以满二叉树方式插入节点
func insertNode(root *TreeNode, Val int) {
	NodeS := make([]*TreeNode, 1, 100)
	NodeS1 := make([]*TreeNode, 0, 100)
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



func levelOrder(root *TreeNode) [][]int {
	return [][]int{{0},}
}