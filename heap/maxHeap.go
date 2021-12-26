package heap

import "sync"

type maxHeap struct {
	arr []int
	mu  sync.Mutex
}

func (m *maxHeap) Insert(item int) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.arr = append(m.arr, item)
	m.heapifyUp(len(m.arr) - 1)
}

func (m *maxHeap) Remove() (int, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	if len(m.arr) == 0 {
		return 0, ErrHeapIsEmpty
	}

	removed := m.arr[0]
	lastIndex := len(m.arr) - 1
	
	m.arr[0] = m.arr[lastIndex]
	m.arr = m.arr[:lastIndex]
	m.heapifyDown(0)

	return removed, nil
}

func (m *maxHeap) heapifyDown(index int) {
	lastIndex := len(m.arr) - 1
	leftChild, rightChild := left(index), right(index)
	childIndex := leftChild

	for leftChild <= lastIndex {
		if m.arr[index] >= m.arr[childIndex] {
			return
		}

		if leftChild == lastIndex || m.arr[leftChild] > m.arr[rightChild] {
			childIndex = leftChild
		} else {
			childIndex = rightChild
		}

		if m.arr[index] < m.arr[childIndex] {
			m.swap(index, childIndex)
			index = childIndex
			leftChild = left(index)
			rightChild = right(index)
		}
	}
}

func (m *maxHeap) heapifyUp(index int) {
	for m.arr[index] > m.arr[parent(index)] {
		m.swap(index, parent(index))
		index = parent(index)
	}
}

func (m *maxHeap) swap(first, second int) {
	m.arr[first], m.arr[second] = m.arr[second], m.arr[first]
}

func (m *maxHeap) Value() []int {
	return m.arr
}

func NewMaxHeap() *maxHeap {
	return &maxHeap{
		arr: make([]int, 0),
		mu:  sync.Mutex{},
	}
}
