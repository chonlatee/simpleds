package stack_test

import (
	"github.com/chonlatee/simpleds/stack"
	"testing"
)

func TestStack(t *testing.T) {
	t.Parallel()
	t.Run("valid push and pop", func(t *testing.T) {
		s := stack.NewStack()
		s.Push(10)
		val, err := s.Pop()
		if err != nil {
			t.Error("expect error is nil")
		}

		if val != 10 {
			t.Errorf("expect `10` but got %d", val)
		}
	})

	t.Run("should error when pop empty stack", func(t *testing.T) {
		s := stack.NewStack()
		val, err := s.Pop()
		if err != stack.ErrStackIsEmpty {
			t.Errorf("expect error %s but got nil",err)
		}

		if val != nil {
			t.Errorf("expect value is nil but got %+v", val)
		}
	})

	t.Run("should get valid value when call peek", func(t *testing.T) {
		s := stack.NewStack()
		s.Push(10)
		val, err := s.Peek()
		if err != nil {
			t.Errorf("expect error is nil but got %+v", err)
		}
		if val != 10 {
			t.Errorf("expect value is `10` but got %+v", val)
		}
	})

	t.Run("should error when call peek in empty stack", func(t *testing.T) {
		s := stack.NewStack()
		val, err := s.Peek()
		if val != nil {
			t.Errorf("expect value is nil but got %+v", val)
		}
		if err != stack.ErrStackIsEmpty {
			t.Errorf("expect error %s but got %+v", stack.ErrStackIsEmpty, err)
		}
	})
}