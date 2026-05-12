package algo_101

const ConstMissVal = -200

type Node struct {
	Val int
	L   *Node
	R   *Node
}

func SymmetricTreeR(in []int, missVal int) bool {
	node := BuildTree(in, missVal)
	if node == nil {
		return false
	}
	return isSymmetric(node.L, node.R)
}

func SymmetricTreeI(in []int, missVal int) bool {
	node := BuildTree(in, missVal)
	if node == nil {
		return false
	}
	q := NewQ()
	q.Push(node.L)
	q.Push(node.R)
	for !q.IsEmpty() {
		l := q.Pop()
		r := q.Pop()
		if l == nil && r == nil {
			continue
		}
		if l == nil || r == nil || l.Val != r.Val {
			return false
		}
		q.Push(l.L)
		q.Push(r.R)
		q.Push(l.R)
		q.Push(r.L)
	}

	return true
}

func isSymmetric(l *Node, r *Node) bool {
	if l == nil && r == nil {
		return true
	}
	if l == nil || r == nil {
		return false
	}

	return l.Val == r.Val && isSymmetric(l.L, r.R) && isSymmetric(l.R, r.L)
}

// BuildTree собираем дерево из полученных данных
// @param mv - передаем число которое обозначет пустые данные. Узел формироваться не будет
func BuildTree(in []int, mv int) *Node {
	countNodes := func(line int) int {
		if line == 0 {
			return 1
		}
		return line * 2
	}
	isNilNode := func(val int) bool {
		return val == mv
	}
	// проверь на пустоту предварительно
	buildNode := func(root *Node, val int, isLeft ...bool) *Node {
		if len(isLeft) > 0 {
			root.L = &Node{
				Val: val,
				L:   nil,
				R:   nil,
			}
			return root.L
		}
		root.R = &Node{
			Val: val,
			L:   nil,
			R:   nil,
		}

		return root.R
	}

	if len(in) == 0 || in == nil {
		return nil
	}
	if isNilNode(in[0]) {
		return nil
	}

	var (
		root = &Node{
			Val: in[0],
			L:   nil,
			R:   nil,
		}
		line = 1
		i    = 1
	)

	tmpItems := []*Node{root}

	for i < len(in) {
		countOfNodes := countNodes(line)
		roots := make([]*Node, len(tmpItems))
		copy(roots, tmpItems)
		tmpItems = make([]*Node, 0, countOfNodes)
		for _, rootItem := range roots {
			if !isNilNode(in[i]) && i < len(in) {
				tmpItems = append(tmpItems, buildNode(rootItem, in[i], true))
			}
			i++
			if !isNilNode(in[i]) && i < len(in) {
				tmpItems = append(tmpItems, buildNode(rootItem, in[i]))
			}
			i++
		}
		line++
	}

	return root
}
