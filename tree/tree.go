package tree

type node struct {
	value     int32
	rightNode *node
	leftNode  *node
}

// Tree struct for tree datastructure
type Tree struct {
	rootNode *node
}

func (t *Tree) Insert(value int32) error {
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

		if value >= currentNode.rightNode.value {
			currentNode = currentNode.rightNode
		} else {
			currentNode = currentNode.leftNode
		}
	}

	return nil
}

func (t *Tree) PreOrderTraversal() {

}

func (t *Tree) PostOrderTraversal() {

}

func (t *Tree) InOrderTraversal() {

}

// NewTree for new tree data structure
func NewTree() *Tree {
	return &Tree{}
}
