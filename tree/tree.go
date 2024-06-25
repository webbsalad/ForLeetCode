package tree

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type NodeInfo struct {
	Val int
	Row int
	Col int
}

func ListToTree(arr []interface{}) *TreeNode {
	if len(arr) == 0 || arr[0] == nil {
		return nil
	}

	root := &TreeNode{Val: arr[0].(int)}
	queue := []*TreeNode{root}
	i := 1

	for i < len(arr) {
		current := queue[0]
		queue = queue[1:]

		if i < len(arr) && arr[i] != nil {
			current.Left = &TreeNode{Val: arr[i].(int)}
			queue = append(queue, current.Left)
		}
		i++

		if i < len(arr) && arr[i] != nil {
			current.Right = &TreeNode{Val: arr[i].(int)}
			queue = append(queue, current.Right)
		}
		i++
	}

	return root
}

func TreeToList(root *TreeNode) []interface{} {
	if root == nil {
		return []interface{}{}
	}

	var arr []interface{}
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current != nil {
			arr = append(arr, current.Val)
			queue = append(queue, current.Left, current.Right)
		} else {
			arr = append(arr, nil)
		}
	}

	for len(arr) > 0 && arr[len(arr)-1] == nil {
		arr = arr[:len(arr)-1]
	}

	return arr
}

func VisualizeTree(root *TreeNode) {
	if root == nil {
		return
	}

	var printNode func(node *TreeNode, level int)
	printNode = func(node *TreeNode, level int) {
		if node == nil {
			return
		}
		printNode(node.Right, level+1)

		fmt.Printf("%*s%d\n", level*4, "", node.Val)

		printNode(node.Left, level+1)
	}

	printNode(root, 0)
}
