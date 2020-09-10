package linklist

import "fmt"

type elm struct {
	val  interface{}
	next *elm
}

type SingleLinklist struct {
	len  int
	head *elm
}

func (s *SingleLinklist) InsertFront(val interface{}) {

}

func (s *SingleLinklist) InsertBack(val interface{}) {

}

func (s *SingleLinklist) InsertBefore(val, before interface{}) {

}

func (s *SingleLinklist) InsertAfter(val, after interface{}) {

}

func (s *SingleLinklist) RemoveFront() {

}

func (s *SingleLinklist) RemoveBack() {

}

func (s *SingleLinklist) RemoveBefore(val, after interface{}) {

}

func (s *SingleLinklist) RemoveAfter(val, after interface{}) {

}

func (s *SingleLinklist) Traverse() {
	current := s.head
	for current != nil {
		fmt.Println(current.val)
	}
}

func (s *SingleLinklist) Size() int {
	return s.len
}

func NewSingleLinklist() *SingleLinklist {
	return &SingleLinklist{}
}
