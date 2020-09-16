package tree

import (
	"testing"
)

func TestNewTree(t *testing.T) {
	tree := NewTree()

	if tree == nil {
		t.Error("expect tree equal nil")
	}
}

func TestInsertNewNode(t *testing.T) {
	tree := NewTree()

	err := tree.Insert(10)
	if err != nil {
		t.Error("expect err equal nil")
	}

	if tree.rootNode == nil {
		t.Error("expect rootNode != nil")
	}

	if tree.rootNode.value != 10 {
		t.Error("expect value of root node equal 10")
	}

}

func TestInsertRightNode(t *testing.T) {
	tree := NewTree()

	err := tree.Insert(10)

	if err != nil {
		t.Error("expect err equal nil")
	}

	err = tree.Insert(20)

	if err != nil {
		t.Error("expect err equal nil")
	}

	if tree.rootNode.rightNode == nil {
		t.Error("expect right node of root node != nil")
	}

	if tree.rootNode.rightNode.value != 20 {
		t.Error("expect value of right node equal 20")
	}

}

func TestInsertLeftNode(t *testing.T) {
	tree := NewTree()

	err := tree.Insert(10)

	if err != nil {
		t.Error("expect err equal nil")
	}

	err = tree.Insert(8)

	if err != nil {
		t.Error("expect err equal nil")
	}

	if tree.rootNode.leftNode == nil {
		t.Error("expect right node of root node != nil")
	}

	if tree.rootNode.leftNode.value != 8 {
		t.Error("expect value of right node equal 8")
	}
}

func TestInsertNode(t *testing.T) {
}
