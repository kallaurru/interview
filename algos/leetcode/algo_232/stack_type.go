package algo_232

const defStackSize = 8

type MyStack struct {
	stor []int // добавляем в конец массива и достаем. Имитируем stack
}

func NewStack() *MyStack {
	return &MyStack{
		stor: make([]int, 0, defStackSize),
	}
}

func (s *MyStack) Push(val int) {
	s.stor = append(s.stor, val)
}

func (s *MyStack) Peek() int {
	last := len(s.stor) - 1

	return s.stor[last]
}

func (s *MyStack) Pop() int {
	val := s.Peek()
	last := len(s.stor) - 1
	newStor := s.stor[0:last]
	s.stor = newStor

	return val
}

func (s *MyStack) Size() int {
	return len(s.stor)
}

func (s *MyStack) IsEmpty() bool {
	return s.stor == nil || len(s.stor) == 0
}
