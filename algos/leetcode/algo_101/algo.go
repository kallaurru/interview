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
	tree := BuildTree(in, missVal)
	return false
}

// BuildTree собираем дерево из полученных данных
// @param mv - передаем число которое обозначет пустые данные. Узел формироваться не будет
func BuildTree(in []int, mv int) *Node {
	nodesCounter := func(line int) int {
		if line == 0 {
			return 1
		}
		return line * 2
	}

	nodeBuilder := func(root *Node, l, r int, missVal int) *Node {
		if l != missVal {
			root.L = &Node{
				Val: l,
				L:   nil,
				R:   nil,
			}
		}
		if r != missVal {
			root.R = &Node{
				Val: r,
				L:   nil,
				R:   nil,
			}
		}
		return root
	}

	var i int = 1
	var line int = 1
	var root *Node = &Node{
		Val: in[0],
		L:   nil,
		R:   nil,
	}

	for i < len(in) {
		nodes := nodesCounter(line)
		for n := 0; n <= nodes; n++ {
			root = nodeBuilder(root, in[i+1], in[i+2], mv)
		}
		line++
	}
	return root
}
