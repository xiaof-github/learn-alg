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

func main() {
	NodeS := make([]*TreeNode, 1, 100)
	NodeS1 := make([]*TreeNode, 0, 100)

	NodeS[0] = new(TreeNode)
	NodeS1 = append(NodeS1, new(TreeNode), new(TreeNode))

	NodeS = NodeS[:0]
	fmt.Println(NodeS)
	NodeS = append(NodeS, NodeS1...)
	fmt.Println("append ", NodeS)
	fmt.Printf("%+v",NodeS)


}