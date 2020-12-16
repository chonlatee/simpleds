package circularQueue_test

import (
	circularQueue "github.com/chonlatee/simpleds/circular-queue"
	"testing"
)

func shouldBeNil(err error, t *testing.T) {
	if err != nil {
		t.Errorf("expect error is nil but got %+v", err)
	}
}

func TestCircularQueue(t *testing.T) {
	t.Parallel()
	t.Run("should error when enqueue", func(t *testing.T) {
		q := circularQueue.NewQueue(5)
		err := q.Enqueue(1)
		shouldBeNil(err, t)
		err = q.Enqueue(2)
		shouldBeNil(err, t)
		err = q.Enqueue(3)
		shouldBeNil(err, t)
		err = q.Enqueue(4)
		shouldBeNil(err, t)
		err = q.Enqueue(5)
		shouldBeNil(err, t)
	})

	t.Run("should error when enqueue with full queue", func(t *testing.T) {
		q := circularQueue.NewQueue(5)
		err := q.Enqueue(1)
		shouldBeNil(err, t)
		err = q.Enqueue(2)
		shouldBeNil(err, t)
		err = q.Enqueue(3)
		shouldBeNil(err, t)
		err = q.Enqueue(4)
		shouldBeNil(err, t)
		err = q.Enqueue(5)
		shouldBeNil(err, t)
		err = q.Enqueue(6)
		if err != circularQueue.ErrQueueIsFull {
			t.Errorf("expect error %s but got %+v", circularQueue.ErrQueueIsFull, err)
		}
	})

	t.Run("should valid dequeue", func(t *testing.T) {
		q := circularQueue.NewQueue(2)
		err := q.Enqueue(1)
		shouldBeNil(err, t)
		err = q.Enqueue(2)
		shouldBeNil(err, t)

		v1, err := q.Dequeue()
		shouldBeNil(err, t)

		v2, err := q.Dequeue()
		shouldBeNil(err, t)
		if v1 != 1 {
			t.Errorf("expect 1 but got %+v", v1)
		}

		if v2 != 2 {
			t.Errorf("expect 2 but got %+v", v2)
		}

		_, err = q.Dequeue()
		if err != circularQueue.ErrQueueIsEmpty {
			t.Errorf("expect error is %s but got %+v", circularQueue.ErrQueueIsEmpty, err)
		}
	})
	
	t.Run("should valid enqueue after dequeue", func(t *testing.T) {
		q := circularQueue.NewQueue(4)
		err := q.Enqueue(1)
		shouldBeNil(err, t)
		err = q.Enqueue(2)
		shouldBeNil(err, t)
		err = q.Enqueue(3)
		shouldBeNil(err, t)
		err = q.Enqueue(4)
		_, err = q.Dequeue()
		shouldBeNil(err, t)
		err = q.Enqueue(5)
		shouldBeNil(err, t)
	})

}
