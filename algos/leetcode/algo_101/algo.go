package algo_101

const ConstMissVal = -200

type Node struct {
	Val int
	L   *Node
	R   *Node
}

func SymmetricTreeR(in []int, missVal int) bool {
	return false
}

func SymmetricTreeI(in []int, missVal int) bool {
	if len(in) == 0 {
		return false
	}
	if len(in)%2 != 1 {
		return false
	}
	tree := BuildTree(in)
	return false
}

func BuildTree(in []int) *Node {
	root := &Node{
		Val: in[0],
		L:   nil,
		R:   nil,
	}
}
