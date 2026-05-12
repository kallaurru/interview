package algo_101

import (
	"container/list"
)

type Queue struct {
	list *list.List
}

func NewQ() *Queue {
	return &Queue{
		list: list.New(),
	}
}

func (q *Queue) Pop() *Node {
	el := q.list.Front()
	q.list.Remove(el)

	return el.Value.(*Node)
}

func (q *Queue) Push(n *Node) {
	q.list.PushBack(n)
}

func (q *Queue) IsEmpty() bool {
	return q.list.Len() == 0
}
