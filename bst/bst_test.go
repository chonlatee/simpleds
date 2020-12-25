package bst_test

import (
	"github.com/chonlatee/simpleds/bst"
	"testing"
)

func TestBST(t *testing.T) {
	createTree := func () *bst.Tree {
		tree := bst.NewBST()
		tree.Insert(100)
		tree.Insert(50)
		tree.Insert(40)
		tree.Insert(60)
		tree.Insert(150)
		tree.Insert(140)
		tree.Insert(160)
		tree.Insert(55)
		tree.Insert(65)
		tree.Insert(155)
		tree.Insert(135)
		return tree
	}
	t.Run("should return valid bst inorder", func(tt *testing.T) {
		tree := createTree()
		actual := tree.Inorder()
		expected := "100-50-40-60-55-65-150-140-135-160-155"
		if actual != expected {
			t.Errorf("expect %s but got %s", expected, actual)
		}
	})
	
	t.Run("valid delete leaf node", func(tt *testing.T) {
		//tree := createTree()
		//tree.Delete(135)
		//actual := tree.Inorder()
		//expected := "100-50-40-60-55-65-150-140-160-155"
		//if actual != expected {
		//	t.Errorf("expect %s but got %s", expected, actual)
		//}
	})

}