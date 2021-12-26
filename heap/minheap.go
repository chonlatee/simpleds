package heap

import (
	"errors"
	"sync"
)

var ErrHeapIsEmpty = errors.New("heap is empty")

type minHeap struct {
	arr  []int
	size int
	mu   sync.Mutex
}

func parent(index int) int {
	return (index - 1) / 2
}

func left(index int) int {
	return index*2 + 1
}

func right(index int) int {
	return index*2 + 2
}

func (h *minHeap) swap(first, second int) {
	h.arr[first], h.arr[second] = h.arr[second], h.arr[first]
}

func (h *minHeap) Insert(item int) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.arr = append(h.arr, item)

	h.heapifyUp(len(h.arr) - 1)
}
func (h *minHeap) heapifyUp(index int) {
	in := index
	for h.arr[in] < h.arr[parent(in)] {
		h.swap(in, parent(in))
		in = parent(in)
	}
}

func (h *minHeap) Remove() (int, error) {
	h.mu.Lock()
	defer h.mu.Unlock()

	if len(h.arr) == 0 {
		return 0, ErrHeapIsEmpty
	}

	removed := h.arr[0]
	arrLength := len(h.arr) - 1

	h.arr[0] = h.arr[arrLength]
	h.arr = h.arr[:arrLength]
	h.heapifyDown(0)

	return removed, nil
}

func (h *minHeap) heapifyDown(index int) {
	lastIndexToCheck := len(h.arr) - 1
	leftChild, rightChild := left(index), right(index)
	childIndex := leftChild

	for leftChild <= lastIndexToCheck {
		if h.arr[index] <= h.arr[childIndex] {
			return
		}

		if leftChild == lastIndexToCheck || h.arr[leftChild] < h.arr[rightChild] {
			childIndex = leftChild
		} else {
			childIndex = rightChild
		}

		if h.arr[index] > h.arr[childIndex] {
			h.swap(index, childIndex)
			index = childIndex
			leftChild = left(index)
			rightChild = right(index)
		}

	}
}

func (h *minHeap) Value() []int {
	return h.arr
}

func NewMinHeap() *minHeap {
	return &minHeap{
		mu:  sync.Mutex{},
		arr: make([]int, 0),
	}
}
