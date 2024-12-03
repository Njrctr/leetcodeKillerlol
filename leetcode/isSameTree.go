package leetcode

import (
	"alg/tracking"
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func wtf() {
	// p = [1,2,3], q = [1,2,3]
	p11 := &TreeNode{Val: 1}
	p12 := &TreeNode{Val: 2, Right: p11, Left: nil}
	p13 := &TreeNode{Val: 3, Right: nil, Left: p11}
	p11.Right, p11.Left = p13, p12

	q11 := &TreeNode{Val: 1}
	q12 := &TreeNode{Val: 2, Right: q11, Left: nil}
	q13 := &TreeNode{Val: 3, Right: nil, Left: q11}
	q11.Right, q11.Left = q13, q12

	defer tracking.Duration(tracking.Track("searchInsert: "))
	fmt.Println(isSameTree(p11, q11))
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil || q == nil {
		return p == q
	}
	if p.Val != q.Val {
		return false
	}
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
