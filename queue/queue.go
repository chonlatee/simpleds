package queue

type MyQueue struct {
	capacity int
	queue    []interface{}
}

func (m *MyQueue) EnQueue(v interface{}) {
	m.queue = append(m.queue, v)
}

func (m *MyQueue) DeQueue() interface{} {
	q := m.queue[0]
	m.queue = m.queue[1:]
	return q
}

func (m *MyQueue) Peek() interface{} {
	q := m.queue[0]
	return q
}

func (m *MyQueue) IsFull() bool {
	return len(m.queue) == m.capacity
}

func (m *MyQueue) IsEmpty() bool {
	return len(m.queue) == 0
}

func NewQueue(cap int) *MyQueue {

	q := &MyQueue{
		capacity: cap,
	}
	return q
}
