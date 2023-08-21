package tree

import "dsa/structures"

func BreadthFirstSearch[T comparable](head BinaryTreeNode[T], needle T) bool {
	queue := structures.NewQueue[*BinaryTreeNode[T]]()
	queue.Enqueue(&head)

	for queue.Length() > 0 {
		curr, err := queue.Dequeue()
		if err != nil {
			panic("Queue is empty")
		}
        if curr == nil {
            continue
        }

		if curr.Value == needle {
			return true
		}

		queue.Enqueue(curr.Left)
		queue.Enqueue(curr.Right)
	}

	return false
}
