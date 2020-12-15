package stack

import (
	"errors"
	"sync"
)

var ErrStackIsEmpty = errors.New("stack is empty")

type MyStack struct {
	mu sync.Mutex
	stack    []interface{}
}

func (m *MyStack) Push(v interface{}) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.stack = append(m.stack, v)
}

func (m *MyStack) Pop() (interface{}, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.Size() == 0 {
		return nil, ErrStackIsEmpty
	}
	s := m.stack[len(m.stack)-1]
	m.stack = m.stack[:len(m.stack)-1]
	return s, nil
}

func (m *MyStack) Peek() (interface{}, error) {
	if m.Size() == 0 {
		return nil, ErrStackIsEmpty
	}
	return m.stack[len(m.stack)-1], nil
}

func (m *MyStack) IsEmpty() bool {
	return len(m.stack) == 0
}

func (m *MyStack) Size() int {
	return len(m.stack)
}

// NewStack ...
func NewStack() *MyStack {
	n := &MyStack{
		mu: sync.Mutex{},
		stack: make([]interface{}, 0),
	}
	return n
}
