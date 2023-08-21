package tree

func InOrderTraversal[T any](head BinaryTreeNode[T]) []T {
	path := []T{}

	in_walk(&head, &path)

	return path
}

func in_walk[T any](node *BinaryTreeNode[T], path *[]T) {
	if node == nil {
		return
	}

	in_walk(node.Left, path)
	*path = append(*path, node.Value)
	in_walk(node.Right, path)
}
