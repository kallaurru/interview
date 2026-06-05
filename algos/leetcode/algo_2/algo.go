package algo_2

type SingleListNode struct {
	Val  int
	Next *SingleListNode
}

func AddTwoNumbers(root1, root2 *SingleListNode) *SingleListNode {
	var (
		rem    = 0
		cursor *SingleListNode
		out    = &SingleListNode{Val: 0, Next: nil}
	)
	sum := root1.Val + root2.Val + rem
	if sum > 10 {
		sum -= 10
		rem = 1
	}
	out.Val = sum
	cursor = out
	root1 = root1.Next
	root2 = root2.Next
	for root1 != nil || root2 != nil {
		l, r := 0, 0
		if root1 != nil {
			l = root1.Val
			root1 = root1.Next
		}
		if root2 != nil {
			r = root2.Val
			root2 = root2.Next
		}
		sum = l + r + rem
		if sum >= 10 {
			cursor.Next = &SingleListNode{Val: sum - 10, Next: nil}
			rem = 1
		} else {
			cursor.Next = &SingleListNode{Val: sum, Next: nil}
			rem = 0
		}
		cursor = cursor.Next
	}
	// если последняя сумма была больше 1
	if rem == 1 {
		cursor.Next = &SingleListNode{Val: 1, Next: nil}
	}
	return out
}

func BuildList(val []int) *SingleListNode {
	if len(val) == 0 {
		return nil
	}
	var head *SingleListNode

	for i := len(val) - 1; i >= 0; i-- {
		node := &SingleListNode{Val: val[i]}
		if head == nil {
			head = node
			continue
		}
		old := head
		head = node
		head.Next = old
	}

	return head
}
