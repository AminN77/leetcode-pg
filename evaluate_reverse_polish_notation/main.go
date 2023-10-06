package main

import (
	"log"
	"strconv"
)

func main() {
	input := []string{"2", "1", "+", "3", "*"}
	log.Println(evalRPN(input))
}

func evalRPN(tokens []string) int {
	stack := NewStack()
	for i := 0; i < len(tokens); i++ {
		switch tokens[i] {

		case "+":
			num2, _ := strconv.Atoi(stack.Pop())
			num1, _ := strconv.Atoi(stack.Pop())
			stack.Push(strconv.Itoa(num1 + num2))

		case "-":
			num2, _ := strconv.Atoi(stack.Pop())
			num1, _ := strconv.Atoi(stack.Pop())
			stack.Push(strconv.Itoa(num1 - num2))

		case "*":
			num2, _ := strconv.Atoi(stack.Pop())
			num1, _ := strconv.Atoi(stack.Pop())
			stack.Push(strconv.Itoa(num1 * num2))

		case "/":
			num2, _ := strconv.Atoi(stack.Pop())
			num1, _ := strconv.Atoi(stack.Pop())
			stack.Push(strconv.Itoa(num1 / num2))
			
		default:
			stack.Push(tokens[i])
		}
	}

	res, _ := strconv.Atoi(stack.Pop())
	return res
}

type Stack struct {
	data []string
}

func NewStack() *Stack {
	return &Stack{
		data: make([]string, 0),
	}
}

func (s *Stack) Push(v string) {
	s.data = append(s.data, v)
}

func (s *Stack) Pop() string {
	top := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return top
}