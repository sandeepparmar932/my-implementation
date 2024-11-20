package main

import "fmt"

type MyStack struct {
	values []interface{}
	top    int
}

func (s *MyStack) push(val interface{}) interface{} {
	s.values = append(s.values, val)
	s.top += 1
	return val
}

func (s *MyStack) pop() interface{} {
	val := s.values[s.top]
	s.values = append(s.values[:s.top])
	s.top -= 1
	return val
}

func (s *MyStack) peek() interface{} {
	return s.values[s.top]
}

func (s *MyStack) isEmpty() bool {
	if s.top == -1 {
		return true
	}
	return false
}
func main() {
	stack := MyStack{nil, -1}
	stack.push(1)
	stack.push("xyz")
	stack.push(2)
	stack.push(3)
	stack.push("abc")
	fmt.Println("stack is ", stack.values)
	stack.pop()
	fmt.Println("stack is ", stack.values)

}
