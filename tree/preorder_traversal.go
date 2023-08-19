package tree

type BinaryTreeNode[T any] struct {
	value T
	Left  *BinaryTreeNode[T]
	Right *BinaryTreeNode[T]
}

func NewBinaryTreeNode[T any](value T) BinaryTreeNode[T] {
	return BinaryTreeNode[T]{
		value: value,
	}
}

func PreOrderTraversal[T any](head BinaryTreeNode[T]) []T {
	path := []T{}

	pre_walk(&head, &path)

	return path
}

func pre_walk[T any](node *BinaryTreeNode[T], path *[]T) {
	if node == nil {
		return
	}

	*path = append(*path, node.value)
	pre_walk(node.Left, path)
	pre_walk(node.Right, path)

}
