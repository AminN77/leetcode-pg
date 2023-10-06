package main

import "log"

func main() {
	input := []int{73, 74, 75, 71, 69, 72, 76, 73}
	log.Println(dailyTemperatures(input))
}

func dailyTemperatures(temperatures []int) []int {
	res := make([]int, len(temperatures))
	stack := NewStack()
	for i := 0; i < len(temperatures); i++ {
		if stack.Top() != nil && stack.Top().data < temperatures[i] {
			top := stack.Top()
			for top != nil && top.data < temperatures[i] {
				res[top.arrive] = i - top.arrive
				stack.Pop()
				top = stack.Top()
			}
		}

		stack.Push(&StackLine{
			data:   temperatures[i],
			arrive: i,
		})

	}

	return res
}

type StackLine struct {
	data   int
	arrive int
}
type Stack struct {
	data []*StackLine
}

func NewStack() *Stack {
	return &Stack{
		data: make([]*StackLine, 0),
	}
}

func (s *Stack) Push(v *StackLine) {
	s.data = append(s.data, v)
}

func (s *Stack) Pop() *StackLine {
	top := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return top
}

func (s *Stack) Top() *StackLine {
	if len(s.data) == 0 {
		return nil
	}

	return s.data[len(s.data)-1]
}
