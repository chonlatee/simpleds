package tree_test

import (
	"github.com/chonlatee/simpleds/tree"
	"testing"
)

func TestNewTree(t *testing.T) {
	tr := tree.NewTree()
	if tr.GetRoot() != nil {
		t.Error("expect tree equal nil")
	}
}

func TestInsertNewNode(t *testing.T) {
	tr := tree.NewTree()
	err := tr.Insert(10)

	if err != nil {
		t.Error("expect err equal nil")
	}

	n := tr.GetRoot()
	v := tr.GetValue(n)
	if v != 10 {
		t.Errorf("expect 10 but got %d", v)
	}
}

func TestTree_IsFullTree(t *testing.T) {
	t.Run("should be true when tree is empty", func(t *testing.T) {
		tr := tree.NewTree()
		r := tr.IsFullTree(tr.GetRoot())
		if !r {
			t.Errorf("expect %t but got %t", true, r)
		}
	})

	t.Run("should be true when tree is one node", func(t *testing.T) {
		tr := tree.NewTree()
		err := tr.Insert(10)
		shouldBeNil(err, t)
		r := tr.IsFullTree(tr.GetRoot())
		if !r {
			t.Errorf("expect %t but got %t", true, r)
		}
	})

	t.Run("Should be tree when valid full tree", func(t *testing.T) {
		tr := tree.NewTree()
		err := tr.Insert(100)
		shouldBeNil(err, t)
		err = tr.Insert(50)
		shouldBeNil(err, t)
		err = tr.Insert(150)
		shouldBeNil(err, t)
		err = tr.Insert(40)
		shouldBeNil(err, t)
		err = tr.Insert(60)
		shouldBeNil(err, t)
		err = tr.Insert(145)
		shouldBeNil(err, t)
		err = tr.Insert(155)
		shouldBeNil(err, t)

		r := tr.IsFullTree(tr.GetRoot())
		if !r {
			t.Errorf("expect %t but got %t", true, r)
		}
	})
}

func shouldBeNil(err error, t *testing.T) {
	if err != nil {
		t.Errorf("expect error is nil but got %+v", err)
	}
}