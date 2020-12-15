package queue_test

import (
	"github.com/chonlatee/simpleds/queue"
	"testing"
)

func TestQueue(t *testing.T) {
	t.Parallel()
	t.Run("success enqueue and dequeue", func(t *testing.T) {
		q := queue.NewQueue(2)
		_ = q.EnQueue(1)
		_ = q.EnQueue(2)

		firstVal, _ := q.DeQueue()
		secondVal, _ := q.DeQueue()
		if firstVal != 1 {
			t.Errorf("expect 1 got %d" , firstVal)
		}
		if secondVal != 2 {
			t.Errorf("expect 2 got %d", secondVal)
		}
	})

	t.Run("error queue is full if enqueue over capacity", func(t *testing.T) {
		q := queue.NewQueue(2)
		_ = q.EnQueue(1)
		_ = q.EnQueue(2)
		err := q.EnQueue(3)

		if err == nil {
			t.Errorf("expect error %s but got %+v", queue.ErrQueueIsFull, err)
		}
	})

	t.Run("error queue is empty when dequeue", func(t *testing.T) {
		q := queue.NewQueue(2)
		_, err := q.DeQueue()
		if err == nil {
			t.Errorf("expect error %s but got %+v", queue.ErrQueueIsEmpty, err)
		}
	})

}