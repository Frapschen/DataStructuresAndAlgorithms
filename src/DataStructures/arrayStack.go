package datastructures

const MAX int = 5

type stack struct {
	data []int
	top  int
}

//初始化
func InitStack() stack {
	return stack{top: -1,
		data: make([]int, MAX)}
}

// 入栈
func (s *stack) Push(data int) bool {
	if s.top == MAX-1 {
		return false
	}
	s.top++
	s.data[s.top] = data
	return true
}

//出栈
func (s *stack) Pop() int {
	if s.top == -1 {
		return -1
	}
	s.top--
	return s.data[s.top+1]
}

//获得栈顶元素
func (s stack) GetTop() int {
	if s.top == -1 {
		return -1
	}
	return s.data[s.top]
}
