package algo_234

type SingleListNode struct {
	Val  int
	Next *SingleListNode
}

func IsPalindromeAlgo234(head *SingleListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	// Step 1: Find middle of linked list (slow/fast pointers)
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// Step 2: Reverse second half
	var prev *SingleListNode
	curr := slow
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	// prev now points to head of reversed second half

	// Step 3: Compare first half with reversed second half
	left, right := head, prev
	for right != nil {
		if left.Val != right.Val {
			return false
		}
		left = left.Next
		right = right.Next
	}

	return true
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
