package main

import "fmt"

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func main(){
	root := &TreeNode{5, nil, nil}
	hasPathSum(root, 5)
}

func hasPathSum(root *TreeNode, target int) bool {
	var sum map[int]*TreeNode
	sum = make(map[int]*TreeNode)
    ret := visitNode(root, 0, sum, target)
	if ret == nil {
		fmt.Println("not found")
		return false
	} else {
		fmt.Printf("found target, node: %v", ret)
		return true
	}
}

func visitNode(node *TreeNode, localSum int, sum map[int]*TreeNode, target int) *TreeNode {
	var ret *TreeNode

    if node == nil {
        return nil
    }

	if node.Left == nil && node.Right == nil {
		// calc path value
		pathValue := node.Val + localSum
		sum[pathValue] = node
		fmt.Println("path value: ", pathValue)
		// 比较map
		val, ok := sum[target]
		if ok {
			fmt.Println("found target")
			return val
		}
	}	

	if node.Left != nil {
		ret = visitNode(node.Left, localSum + node.Val, sum, target)
		if ret != nil {
			return ret
		}

	}	

	if node.Right != nil {
		ret = visitNode(node.Right, localSum + node.Val, sum, target)
		if ret != nil {
			return ret
		}
	}

	return nil
}