package heap_test

import (
	"github.com/chonlatee/simpleds/heap"
	"testing"
)

func TestMaxHeap(t *testing.T) {
	t.Run("valid max heap", func(t *testing.T) {
		h := heap.NewHeap()
		h.Insert(1)
		h.Insert(4)
		h.Insert(2)
		h.Insert(5)
		h.Insert(7)
		h.Insert(9)

		actual := h.ToString()

		if actual != "957142" {
			t.Errorf("expect `957142` but got `%s`", actual)
		}
	})

	t.Run("valid peek", func(t *testing.T) {
		h := heap.NewHeap()
		h.Insert(1)
		h.Insert(7)
		h.Insert(2)
		h.Insert(3)

		val, err := h.Peek()
		if err != nil {
			t.Error("expect error is nil")
		}
		if val != 7 {
			t.Errorf("expect 7 but got %d", val)
		}
	})

	t.Run("should return error when peek empty heap", func(t *testing.T) {
		h := heap.NewHeap()
		_, err := h.Peek()
		if err != heap.ErrHeapIsEmpty {
			t.Errorf("expect err is %s but got %+v",heap.ErrHeapIsEmpty, err)
		}
	})

	t.Run("should valid remove", func(t *testing.T) {
		h := heap.NewHeap()
		h.Insert(3)
		h.Insert(9)
		h.Insert(2)
		h.Insert(1)
		h.Insert(4)
		h.Insert(5)

		h.Remove(4)

		str := h.ToString()

		if str != "93512"{
			t.Errorf("expect 93512 but got %s", str)
		}
	})

}