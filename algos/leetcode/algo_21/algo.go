package algo_21

type SingleListNode struct {
	Val  int
	Next *SingleListNode
}

func MergeTwoArrayAlgo21(l1, l2 []int) []int {
	return nil
}

func MergeTwoListAlgo21(l1, l2 *SingleListNode) *SingleListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	head := &SingleListNode{
		Val:  0,
		Next: nil,
	}

	tmp := head
	if l1.Val <= l2.Val {
		tmp.Next = l1
		l1 = l1.Next
	} else {
		tmp.Next = l2
		l2 = l2.Next
	}
	tmp = tmp.Next

	for {
		if l1 == nil {
			tmp.Next = l2
			break
		}
		if l2 == nil {
			break
		}
		if l1.Val == l2.Val {
			tmp.Next = l1
			tmp = tmp.Next
			tmp.Next = l2
			l1 = l1.Next
			l2 = l2.Next
			tmp = tmp.Next
			continue
		}
		if l1.Val < l2.Val {
			tmp.Next = l1
			l1 = l1.Next
			tmp = tmp.Next
			continue
		}

		if l1.Val > l2.Val {
			tmp.Next = l2
			l2 = l2.Next
			tmp = tmp.Next
			continue
		}

	}

	return head.Next
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
