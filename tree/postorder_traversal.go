package tree

func PostOrderTraversal[T any](head BinaryTreeNode[T]) []T {
	path := []T{}

	post_walk(&head, &path)

	return path
}

func post_walk[T any](node *BinaryTreeNode[T], path *[]T) {
	if node == nil {
		return
	}

	post_walk(node.Left, path)
	post_walk(node.Right, path)
	*path = append(*path, node.Value)
}
