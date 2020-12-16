package circularQueue

import (
	"errors"
	"sync"
)

var ErrQueueIsFull = errors.New("queue is full")
var ErrQueueIsEmpty = errors.New("queue is empty")

type MyQueue struct {
	mu sync.Mutex
	front int
	rear int
	capacity int
	queue []interface{}
}

func (m *MyQueue) Enqueue(v interface{}) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.IsFull() {
		return ErrQueueIsFull
	}
	if m.front == -1 {
		m.front = 0
	}

	m.rear = (m.rear + 1) % m.capacity
	m.queue[m.rear] = v
	return nil
}

func (m *MyQueue) Dequeue() (interface{}, error){
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.IsEmpty() {
		return nil, ErrQueueIsEmpty
	}

	e := m.queue[m.front]
	if m.front == m.rear {
		m.front = -1
		m.rear = -1
		return e, nil
	}

	m.front = (m.front + 1) % m.capacity

	return e, nil

}

func (m *MyQueue) IsEmpty() bool {
	return m.front == -1
}

func (m *MyQueue) IsFull() bool {
	if m.front == 0 && m.rear == m.capacity - 1 {
		return true
	}
	if m.front == m.rear + 1 {
		return true
	}

	return false
}

func NewQueue(cap int) *MyQueue {
	return &MyQueue{
		mu: sync.Mutex{},
		front:    -1,
		rear:     -1,
		capacity: cap,
		queue: make([]interface{}, cap),
	}
}