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

	a := 1
	r = new(TreeNode)
	insertNode(r, 1)

}

func insertNode(root *TreeNode, Val int) {
	NodeS := make([]*TreeNode, 10000)
	NodeS1 := make([]*TreeNode, 10000)
	if (root.Left == nil && root.Right == nil) {
		root.Val = Val
		return;
	}

	NodeS[0] = root
	done := false
	i:=0	

	for (!done) {
		j := 1
		for k, v := range NodeS{						
			NodeS1 = append(NodeS1, k.Left, k.Right)
		}
		k := 0 
		j = j*2
		for ;k<j;k++ {
			if(NodeS1[k] == nil) {
				if (k % 2 == 1) {
					NodeS[k/2].Right = new(TreeNode)
					NodeS[k/2].Right.Val = Val
				} else {
					NodeS[k/2].Left = new(TreeNode)
					NodeS[k/2].Left.Val = Val
				}
				done = true
				break
			}
		}		
		copy(NodeS1, NodeS)
		NodeS1 = NodeS1[:0]
		fmt.Println("NodeS1, len: cap: ", len(NodeS1), cap(NodeS1))
	}
	
}



func levelOrder(root *TreeNode) [][]int {
	return [][]int{{0},}
}