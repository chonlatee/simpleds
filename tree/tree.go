package tree

import (
	"fmt"
	"log"
	"strconv"
)

type node struct {
	value     int
	rightNode *node
	leftNode  *node
}

// Tree struct for tree data structure
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

func (t *Tree) InsertLeft(parent, new int) {
	n := t.GetNodeByValue(parent)
	if n == nil {
		log.Fatalln("node is nil")
		return
	}
	if n.leftNode != nil {
		return
	}

	n.leftNode = &node{value: new}
}

func (t *Tree) InsertRight(parent, new int) {
	n := t.GetNodeByValue(parent)
	if n == nil {
		log.Fatalln("node is nil")
		return
	}
	if n.rightNode != nil {
		return
	}
	n.rightNode = &node{value: new}
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

func (t *Tree) IsBalance() bool {
	return maxDept(t.rootNode) - minDept(t.rootNode)  <= 1
}

func (t *Tree) GetRoot() *node {
	return t.rootNode
}

func (t *Tree) GetValue(root *node) int {
	return root.value
}

func (t *Tree) GetNodeByValue(val int) *node {
	st := make([]*node, 0)
	st = append(st, t.rootNode) // stack.push()
	for len(st) != 0 {
		n := st[len(st) - 1] // stack.peek()
		st = st[:len(st) - 1] // stack.pop()

		if n == nil {
			continue
		}

		if n.value == val {
			return n
		}

		st = append(st, n.rightNode) // stack.push()
		st = append(st, n.leftNode) // stack.push()
	}

	return nil
}

func minDept(root *node)  int {
	if root == nil {
		return 0
	}
	return 1 + min(minDept(root.leftNode), minDept(root.rightNode))
}

func maxDept(root *node) int {
	if root == nil {
		return 0
	}

	return 1 + max(maxDept(root.leftNode), maxDept(root.rightNode))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}



// NewTree for new tree data structure
func NewTree() *Tree {
	return &Tree{}
}
