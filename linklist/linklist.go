package linklist

import (
	"fmt"
	"log"
)

type elm struct {
	val  int
	next *elm
}

type SingleLinklist struct {
	len  int
	head *elm
	tail *elm
}

func (s *SingleLinklist) InsertFront(val int) {
	l := &elm{
		val: val,
	}
	if s.IsEmpty() {
		s.head = l
		s.tail = l
	} else {
		l.next = s.head
		s.head = l
	}
	s.len++
}

func (s *SingleLinklist) InsertBack(val int) {
	l := &elm{
		val: val,
	}

	if s.IsEmpty() {
		s.head = l
		s.tail = l
	} else {
		s.tail.next = l
		s.tail = l
	}

	s.len++
}

// InsertBefore insert before found value InsertBefore(value, before interface{})
func (s *SingleLinklist) InsertBefore(val, before int) {

	l := &elm{
		val: val,
	}
	prev := s.head
	current := s.head

	for current != nil && current.val != before {

		prev = current
		current = current.next
	}

	if current == nil {
		log.Fatalln("Item not found")
	}

	prev.next = l
	l.next = current

	s.len++
}

// InsertAfter insert after found value InsertAfter(value, after interface{})
func (s *SingleLinklist) InsertAfter(val, after int) {

	l := &elm{
		val: val,
	}

	current := s.head
	prev := s.head

	for current != nil && current.val != after {

		prev = current
		current = current.next
	}

	if current == nil {
		log.Fatalln("Item not found")
	}

	if current.val == s.tail.val {
		s.tail.next = l
		s.tail = l
	} else {
		l.next = current.next
		prev.next = l

	}

	s.len++
}

func (s *SingleLinklist) RemoveFront() {
	if s.IsEmpty() {
		log.Fatalln("Link list is empty.")
	}
	s.head = s.head.next
	s.len--
}

func (s *SingleLinklist) RemoveBack() {
	prev := s.head
	current := s.head
	for current.next != nil {
		prev = current
		current = current.next
	}
	s.tail = prev
	s.tail.next = nil
	s.len--
}

// RemoveBefore use for remove node before given value RemoveBefore(before interface{})
func (s *SingleLinklist) RemoveBefore(before interface{}) {

	current := s.head
	target := s.head
	prevTarget := s.head

	for current != nil && target.next.val != before && current.next != nil {
		prevTarget = target
		target = current
		current = current.next
	}

	if current == nil {
		log.Fatalln("Item not found")
	}

	prevTarget.next = target.next

	s.len--
}

func (s *SingleLinklist) RemoveAfter(after interface{}) {

	current := s.head
	prev := s.head

	for current != nil && prev.val != after {
		prev = current
		current = current.next
	}

	if current == nil {
		log.Fatalln("Item not found")
	}

	prev.next = current.next

	if prev.next == nil {
		s.tail = prev
	}

	s.len--
}

func (s *SingleLinklist) Traverse() {
	current := s.head
	for current != nil {
		fmt.Printf("%d ", current.val)
		current = current.next
	}
	fmt.Println()
}

func (s *SingleLinklist) Size() int {
	return s.len
}

func (s *SingleLinklist) IsEmpty() bool {
	return s.head == nil && s.tail == nil && s.len == 0
}

func (s *SingleLinklist) ShowHeadValue() {
	fmt.Println(s.head.val)
}

func (s *SingleLinklist) ShowTailvalue() {
	fmt.Println(s.tail.val)
}

func NewSingleLinklist() *SingleLinklist {
	return &SingleLinklist{}
}
