package bst

import (
	"fmt"
	"strings"
)

type Tree struct {
	root *node
}

type node struct {
	value     int
	leftNode  *node
	rightNode *node
}

func (t *Tree) Insert(data int) {
	if t.root == nil {
		t.root = &node{
			value: data,
		}
		return
	}

	currentNode := t.root
	for currentNode != nil {
		if data < currentNode.value {
			if currentNode.leftNode == nil {
				currentNode.leftNode = &node{value: data}
				break
			} else {
				currentNode = currentNode.leftNode
				continue
			}
		}

		if data >= currentNode.value {
			if currentNode.rightNode == nil {
				currentNode.rightNode = &node{value: data}
				break
			} else {
				currentNode = currentNode.rightNode
				continue
			}
		}
	}
}

func (t *Tree) Delete(data int) {

}

func (t *Tree) Inorder() string {
	st := make([]*node, 0)
	st = append(st, t.root)

	var str string
	for len(st) != 0 {
		n := st[len(st)-1]
		st = st[:len(st)-1]
		if n == nil {
			continue
		}
		str += fmt.Sprintf("%d-", n.value)
		st = append(st, n.rightNode)
		st = append(st, n.leftNode)
	}

	str = strings.TrimSuffix(str, "-")
	return str
}

func NewBST() *Tree {
	return &Tree{}
}
