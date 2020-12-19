package tree

import (
	"fmt"
	"strconv"
)

type node struct {
	value     int
	rightNode *node
	leftNode  *node
}

// Tree struct for tree datastructure
type Tree struct {
	rootNode *node
}

func (t *Tree) Insert(value int) error {
	n := &node{
		value: value,
	}

	if t.rootNode == nil {
		t.rootNode = n
		return nil
	}

	currentNode := t.rootNode

	for currentNode != nil {
		if currentNode.rightNode == nil && value >= currentNode.value {
			currentNode.rightNode = n
			break
		}

		if currentNode.leftNode == nil && value < currentNode.value {
			currentNode.leftNode = n
			break
		}

		if value >= currentNode.value {
			currentNode = currentNode.rightNode
		} else {
			currentNode = currentNode.leftNode
		}
	}

	return nil
}

func (t *Tree) InOrder(root *node) {
	if root  != nil {
		t.InOrder(root.leftNode)
		fmt.Printf(strconv.Itoa(root.value) + " -> ")
		t.InOrder(root.rightNode)
	}
}

func (t *Tree) PostOrder(root *node) {
	if root != nil {
		t.PostOrder(root.leftNode)
		t.PostOrder(root.rightNode)
		fmt.Printf(strconv.Itoa(root.value) + " -> ")
	}
}

func (t *Tree) PreOrder(root *node) {
	if root != nil {
		fmt.Printf(strconv.Itoa(root.value) + " -> ")
		t.PreOrder(root.leftNode)
		t.PreOrder(root.rightNode)
	}
}

func (t *Tree) IsFullTree(root *node) bool {
	if root == nil {
		return true
	}

	if root.leftNode == nil && root.rightNode == nil {
		return true
	}

	if root.leftNode != nil && root.rightNode != nil {
		return t.IsFullTree(root.leftNode) && t.IsFullTree(root.rightNode)
	}

	return false
}

func (t *Tree) GetRoot() *node {
	return t.rootNode
}

func (t *Tree) GetValue(root *node) int {
	return root.value
}


// NewTree for new tree data structure
func NewTree() *Tree {
	return &Tree{}
}
