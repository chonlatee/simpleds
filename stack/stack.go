package stack

import (
	"log"
)

type MyStack struct {
	capacity int
	stack    []interface{}
}

func (m *MyStack) Push(v interface{}) {
	m.stack = append(m.stack, v)
	m.capacity++
}

func (m *MyStack) Pop() interface{} {
	if m.Size() == 0 {
		log.Fatalln("Stack is empty")
	}
	s := m.stack[len(m.stack)-1]
	m.stack = m.stack[:len(m.stack)-1]
	m.capacity--
	return s
}

func (m *MyStack) Peek() interface{} {
	return m.stack[len(m.stack)-1]
}

func (m *MyStack) IsEmpty() bool {
	return len(m.stack) == 0
}

func (m *MyStack) Size() int {
	return len(m.stack)
}

// NewStack ...
func NewStack(size int) *MyStack {
	n := &MyStack{
		capacity: size,
	}
	return n
}
