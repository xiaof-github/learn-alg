package main

import (
	"fmt"
)

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

var (
	r = TreeNode{0, nil, nil}
)

// 求二叉树层序遍历
func main() {
	// nums1 := []int{23,2,2,1,1,1,1,1,21,230,3,3,91}
	// nums2 := []int{1,2,3,4,5,91,21,230,3,23,3,3,5}

	a := 1
	insertNode(r, 1)

}

func insertNode(root *TreeNode, Val int) {
	NodeS := make([]*TreeNode, 10000)
	NodeS1 := make([]*TreeNode, 10000)
	if (TreeNode.Left == nil && TreeNode.Right == nil) {
		TreeNode.Val = Val
		return;
	}

	NodeS[0] = root
	done := false
	for (!done) {
		for k, v = range NodeS {
			if (v == nil) {
				NodeS[k] = new(TreeNode)
				NodeS[k].Val = Val
				done = true
				break			
			}
			NodeS1 = append(NodeS1, k.Left, k.Right)
		}
		copy(NodeS1, NodeS)
		NodeS1 = NodeS1[:0]
		fmt.Println("NodeS1, len: cap: ", len(NodeS1), cap(NodeS1))
	}
	
}



func levelOrder(root *TreeNode) [][]int {
	return [][]int{{0},}
}