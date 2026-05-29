package algo_938

type Node struct {
	Val int
	L   *Node
	R   *Node
}

func RangeSumBST938(tree *Node, low, high int) int {
	sum := 0
	if tree == nil {
		return sum
	}
	if tree.Val <= high && tree.Val >= low {
		sum += tree.Val
		sum += RangeSumBST938(tree.L, low, high)
		sum += RangeSumBST938(tree.R, low, high)

		return sum
	}
	if tree.Val < low {
		sum += RangeSumBST938(tree.R, low, high)
		return sum
	}
	if tree.Val > high {
		sum += RangeSumBST938(tree.L, low, high)
		return sum
	}

	return sum
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
