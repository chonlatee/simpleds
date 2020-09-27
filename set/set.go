package set

type Set struct {
	listItem []int
}

func (s *Set) Add(v int) {
	if !s.Contain(v) {
		s.listItem = append(s.listItem, v)
	}
}

func (s *Set) Remove(value int) {
	nl := []int{}

	for _, v := range s.listItem {
		if v != value {
			nl = append(nl, v)
		}
	}
	s.listItem = nl
}

func (s *Set) Contain(value int) bool {
	for _, v := range s.listItem {
		if v == value {
			return true
		}
	}

	return false
}

func (s *Set) Union(ns Set) *Set {
	return nil
}

func (s *Set) Intersect(ns Set) *Set {
	return nil
}

func (s *Set) Difference(ns Set) *Set {
	return nil
}

func (s *Set) Length() int {
	return len(s.listItem)
}
