package queue

import (
	"errors"
	"sync"
)

var ErrQueueIsFull = errors.New("queue is full")
var ErrQueueIsEmpty = errors.New("queue is empty")

type MyQueue struct {
	mu sync.Mutex
	capacity int
	len int
	queue    []interface{}
}

func (m *MyQueue) EnQueue(v interface{}) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.IsFull() {
		return ErrQueueIsFull
	}
	m.queue = append(m.queue, v)
	m.len++
	return nil
}

func (m *MyQueue) DeQueue() (interface{}, error) {
	m.mu.Lock()
	m.mu.Unlock()
	if m.IsEmpty() {
		return nil, ErrQueueIsEmpty
	}
	q := m.queue[0]
	m.queue = m.queue[1:]
	m.len--
	return q, nil
}

func (m *MyQueue) Peek() interface{} {
	q := m.queue[0]
	return q
}

func (m *MyQueue) IsFull() bool {
	return m.capacity == m.len
}

func (m *MyQueue) IsEmpty() bool {
	return m.len == 0
}

func NewQueue(cap int) *MyQueue {
	q := &MyQueue{
		mu: sync.Mutex{},
		capacity: cap,
		len: 0,
		queue: make([]interface{}, 0),
	}
	return q
}
