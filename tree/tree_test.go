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
		tr.InsertLeft(100, 50)
		tr.InsertRight(100, 150)
		tr.InsertLeft(50, 40)
		tr.InsertRight(50, 60)
		tr.InsertLeft(150, 145)
		tr.InsertRight(150, 160)

		r := tr.IsFullTree(tr.GetRoot())
		if !r {
			t.Errorf("expect %t but got %t", true, r)
		}
	})
}

func TestTree_IsBalance(t *testing.T) {
	t.Run("tree is balance", func(tt *testing.T) {
		tr := tree.NewTree()
		err := tr.Insert(100)
		shouldBeNil(err, t)
		tr.InsertLeft(100, 50)
		tr.InsertRight(100, 150)
		tr.InsertLeft(50, 40)
		tr.InsertRight(50, 60)
		tr.InsertLeft(150, 145)
		tr.InsertRight(150, 160)

		r := tr.IsBalance()
		if !r {
			t.Errorf("expect tree is balance is true but got %t",r )
		}
	})

	t.Run("tree is not balance", func(tt *testing.T) {

	})
}

func shouldBeNil(err error, t *testing.T) {
	if err != nil {
		t.Errorf("expect error is nil but got %+v", err)
	}
}