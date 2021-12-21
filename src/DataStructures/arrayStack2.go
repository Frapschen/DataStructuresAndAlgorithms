package datastructures

const MAX2 int = 20

type stack2 struct {
	data []interface{}
	top  int
}

//初始化
func InitStack2() stack2 {
	return stack2{top: -1,
		data: make([]interface{}, MAX2)}
}

// 入栈
func (s *stack2) Push2(data interface{}) bool {
	if s.top == MAX-1 {
		return false
	}
	s.top++
	s.data[s.top] = data
	return true
}

//出栈
func (s *stack2) Pop2() interface{} {
	if s.top == -1 {
		return -1
	}
	s.top--
	return s.data[s.top+1]
}

//获得栈顶元素
func (s stack2) GetTop2() interface{} {
	if s.top == -1 {
		return -1
	}
	return s.data[s.top]
}

//判断空
func (s stack2) IsEmpty() bool {
	if s.top == -1 {
		return true
	}
	return false
}
