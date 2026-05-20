package algo_232

type MyQueue struct {
	s1 *MyStack
	s2 *MyStack
}

func New() *MyQueue {
	return &MyQueue{
		s1: NewStack(),
		s2: NewStack(),
	}
}

func (mq *MyQueue) Push(x int) {
	if mq.s1.IsEmpty() {
		mq.s1.Push(x)
		return
	}
	mq.s2.Push(x)
}

func (mq *MyQueue) Pop() int {
	val := mq.s1.Pop()
	if mq.s2.IsEmpty() {
		return val
	}
	mq.rearrange()

	return val
}

func (mq MyQueue) Peek() int {
	return mq.s1.Peek()
}

func (mq *MyQueue) Empty() bool {
	return mq.s1.IsEmpty()
}

func (mq *MyQueue) rearrange() {
	tmp := make([]int, 0, mq.s2.Size())
	for !mq.s2.IsEmpty() {
		tmp = append(tmp, mq.s2.Pop())
	}
	mq.s1.Push(tmp[len(tmp)-1])
	tmp = tmp[0 : len(tmp)-1]
	for _, val := range tmp {
		mq.s2.Push(val)
	}
}
