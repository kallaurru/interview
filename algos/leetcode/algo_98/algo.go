package algo_98

import a "github.com/kallaurru/interview/algos/leetcode/algo_938"

func IsValidBSTAlgo98(root *a.Node) bool {
	if root.L == nil && root.R == nil {
		return true
	}
	if root.L != nil && root.L.Val > root.Val {
		return false
	}
	if root.R != nil && root.R.Val < root.Val {
		return false
	}
	l := IsValidBSTAlgo98(root.L)
	r := IsValidBSTAlgo98(root.R)

	return l && r
}

func BuildTree(in []int, mv int) *a.Node {
	return a.BuildTree(in, mv)
}
