package algo_19

type SingleListNode struct {
	Val  int
	Next *SingleListNode
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
func RemoveNthFromEndAlgo19(root *SingleListNode, n int) *SingleListNode {
	head := &SingleListNode{Next: root}

	fast := head
	slow := head

	for i := 0; i <= n; i++ {
		fast = fast.Next
	}

	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}

	slow.Next = slow.Next.Next

	return head.Next
}
