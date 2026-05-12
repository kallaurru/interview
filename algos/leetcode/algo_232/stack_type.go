package algo_232

const defStackSize = 8

type MyStack struct {
	stor   []int // добавляем в конец массива и достаем. Имитируем stack
	size   int
	master bool
}

func NewStack(master ...uint8) *MyStack {
	return &MyStack{
		stor:   make([]int, 0, defStackSize),
		size:   defStackSize,
		master: len(master) == 0,
	}
}

func (s *MyStack) Push(val int) bool {
	if len(s.stor) == s.size {
		return false
	}

	if len(s.stor) == 0 {
		s.stor = append(s.stor, val)
		return true
	}
	tmp := make([]int, 0, s.size)
	tmp = append(tmp, val)
	s.stor = tmp

	return true
}

func (s *MyStack) Peek() int {
	el := s.stor.Front()

	return el.Value.(int)
}

func (s *MyStack) Pop() int {
	el := s.stor.Front()
	s.stor.Remove(el)

	return el.Value.(int)
}

func (s *MyStack) Size() int {
	return s.stor.Len()
}

func (s *MyStack) IsEmpty() bool {
	return s.stor.Len() == 0
}
