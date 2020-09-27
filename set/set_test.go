package set

import "testing"

func TestEmptySet(t *testing.T) {
	s := &Set{}

	l := s.Length()

	if l != 0 {
		t.Errorf("Expect length equal 0 but got %d", l)
	}
}

func TestAddItem(t *testing.T) {
	s := &Set{}

	s.Add(3)

	if !s.Contain(3) {
		t.Errorf("expect Set contain 3")
	}
}

func TestRemoveItem(t *testing.T) {
	s := &Set{
		listItem: []int{20, 90, 10},
	}

	s.Remove(20)

	if s.Contain(20) {
		t.Errorf("Expect 20 no contain in set")
	}
}

func TestDuplicateItem(t *testing.T) {
	s := &Set{}

	s.Add(10)
	s.Add(20)
	s.Add(10)

	f := []int{}

	for _, v := range s.listItem {
		if v == 10 {
			f = append(f, v)
		}
	}

	length := len(f)

	if len(f) > 1 {
		t.Errorf("Expect length of found item equal 0 but got %d", length)
	}

}

func TestLength(t *testing.T) {
	s := &Set{}

	s.Add(10)
	s.Add(20)
	s.Add(100)

	length := s.Length()

	if length != 3 {
		t.Errorf("Expect length equal 3 but got %d", length)
	}
}
