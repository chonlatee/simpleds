package heap

import (
	"errors"
	"strconv"
	"sync"
)

var ErrHeapIsEmpty = errors.New("heap is empty")

type MyHeap struct {
	h []int
	size int
	mu sync.Mutex
}

func (m *MyHeap) heapify(i int) {
	lg := i
	l := 2 * i + 1
	r := 2 * i + 2
	if l < m.size && m.h[l] > m.h[lg] {
		lg = l
	} else if r < m.size && m.h[r] > m.h[lg] {
		lg = r
	}

	if lg != i {
		m.h[i], m.h[lg] = m.h[lg], m.h[i]
		m.heapify(lg)
	}
}

func (m *MyHeap) Peek() (int, error) {
	if m.size < 1 {
		return -1, ErrHeapIsEmpty // should check err because real data might be -1
	}
	return m.h[0], nil
}

func (m *MyHeap) Remove(num int) {
	m.mu.Lock()
	defer m.mu.Unlock()
	i := 0
	length := m.size - 1
	for i < length {
		if num == m.h[i] {
			break
		}
		i++
	}

	m.h[i], m.h[m.size-1] = m.h[i], m.h[m.size - 1]
	m.h = m.h[:length]
	m.size--

	j := m.size / 2 - 1
	for j >= 0 {
		m.heapify(j)
		j--
	}
}

func (m *MyHeap) ToString() string {
	str := ""
	for _, v := range m.h {
		str += strconv.Itoa(v)
	}
	return str
}



func (m *MyHeap) Insert(val int) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.size == 0 {
		m.h = append(m.h, val)
		m.size++
	} else {
		m.h = append(m.h, val)
		m.size++
		i := m.size / 2 - 1
		for i >= 0 {
			m.heapify(i)
			i--
		}
	}
}

func (m *MyHeap) length() int {
	return len(m.h)
}

func NewHeap() *MyHeap {
	return &MyHeap{
		mu: sync.Mutex{},
		h: make([]int, 0),
	}
}
