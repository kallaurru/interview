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
	masterSize := 8
	// проблема при добавлении после 8 элементов
	if mq.s1.IsEmpty() {
		mq.s1.Push(x)
		return
	}
	if mq.s1.Size() == masterSize {
		mq.s2.Push(x)
		return
	}
	// нужно переложить
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

func (mq *MyQueue) rearrange(val int) {
	tmp := make([]int, 0, mq.s1.Size())
	for !mq.s1.IsEmpty() {
		tmp = append(tmp, mq.s1.Pop())
	}
	mq.s1.Push(val)
	for i := len(tmp) - 1; i >= 0; i-- {
		mq.s1.Push(tmp[i])
	}
}
