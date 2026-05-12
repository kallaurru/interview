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
	}
}

func (mq *MyQueue) Pop() int {
	return mq.s1.Pop()
}

func (mq MyQueue) Peek() int {
	return mq.s1.Peek()
}

func (mq *MyQueue) Empty() bool {
	return mq.s1.IsEmpty()
}
