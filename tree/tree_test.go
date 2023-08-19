package tree_test

import (
	"dsa/tree"
	"reflect"
	"testing"
)

func create_traversal_graph() tree.BinaryTreeNode[int] {
	traversal_head := tree.NewBinaryTreeNode(7)
	l1 := tree.NewBinaryTreeNode(23)
	l2 := tree.NewBinaryTreeNode(5)
	l3 := tree.NewBinaryTreeNode(4)
	r1 := tree.NewBinaryTreeNode(3)
	r2 := tree.NewBinaryTreeNode(18)
	r3 := tree.NewBinaryTreeNode(21)

	traversal_head.Left = &l1
	l1.Left = &l2
	l1.Right = &l3
	traversal_head.Right = &r1
	r1.Left = &r2
	r1.Right = &r3

	return traversal_head
}

func TestPreOrder(t *testing.T) {
	head := create_traversal_graph()
	expected := []int{7, 23, 5, 4, 3, 18, 21}
	actual := tree.PreOrderTraversal(head)

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected: %+v, Receieved: %+v", expected, actual)
	}
}

func TestInOrder(t *testing.T) {
	head := create_traversal_graph()
	expected := []int{5, 23, 4, 7, 18, 3, 21}
	actual := tree.InOrderTraversal(head)

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected: %+v, Receieved: %+v", expected, actual)
	}
}

func TestPostOrder(t *testing.T) {
	head := create_traversal_graph()
	expected := []int{5, 4, 23, 18, 21, 3, 7}
	actual := tree.PostOrderTraversal(head)

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected: %+v, Receieved: %+v", expected, actual)
	}
}
