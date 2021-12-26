package heap_test

import (
	"testing"

	"github.com/chonlatee/simpleds/heap"
	"github.com/stretchr/testify/require"
)

func TestMaxHeap(t *testing.T) {
	t.Run("valid insert min heap", func(t *testing.T) {
		h := heap.NewMinHeap()
		h.Insert(9)
		h.Insert(8)
		h.Insert(10)
		h.Insert(2)

		require.Equal(t, []int{2, 8, 10, 9}, h.Value())
	})

	t.Run("valid remove min heap", func(t *testing.T) {
		h := heap.NewMinHeap()
		h.Insert(9)
		h.Insert(8)
		h.Insert(7)
		h.Insert(6)
		h.Insert(5)
		h.Insert(4)

		_, err := h.Remove()
		require.NoError(t, err)

		require.Equal(t, []int{5, 6, 8, 9, 7}, h.Value())
	})

	t.Run("error heap is empty", func(t *testing.T) {
		h := heap.NewMinHeap()
		h.Insert(10)
		h.Insert(9)
		i, err := h.Remove()
		require.NoError(t, err)
		require.Equal(t, 9, i)
		i, err = h.Remove()
		require.NoError(t, err)
		require.Equal(t, 10, i)
		_, err = h.Remove()
		require.Equal(t, heap.ErrHeapIsEmpty, err)
	})
}
