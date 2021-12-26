package heap_test

import (
	"testing"

	"github.com/chonlatee/simpleds/heap"
	"github.com/stretchr/testify/require"
)

func TestMaxHeap_Insert(t *testing.T) {
	t.Run("valid max heap", func(t *testing.T) {
		h := heap.NewMaxHeap()
		h.Insert(1)
		h.Insert(2)
		h.Insert(3)
		h.Insert(4)
		h.Insert(5)

		require.Equal(t, []int{5, 4, 2, 1, 3}, h.Value())
	})

	t.Run("valid max heap when remove", func(t *testing.T) {
		h := heap.NewMaxHeap()
		h.Insert(1)
		h.Insert(2)
		h.Insert(3)
		h.Insert(4)
		h.Insert(5)

		h.Remove()

		require.Equal(t, []int{4, 3, 2, 1}, h.Value())
	})

	t.Run("error when remove empty heap", func(t *testing.T) {
		h := heap.NewMaxHeap()
		h.Insert(1)
		h.Insert(2)

		i, err := h.Remove()
		require.NoError(t, err)
		require.Equal(t, 2, i)

		i, err = h.Remove()
		require.NoError(t, err)
		require.Equal(t, 1, i)

		_, err = h.Remove()
		require.Equal(t, heap.ErrHeapIsEmpty, err)
	})
}
