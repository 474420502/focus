package listtree

func (tree *ListTree) getHeight() int {
	root := tree.getRoot()
	if root == nil {
		return 0
	}

	var height = 1

	var traverse func(cur *Node, h int)
	traverse = func(cur *Node, h int) {

		if cur == nil {
			return
		}

		if h > height {
			height = h
		}

		traverse(cur.children[0], h+1)
		traverse(cur.children[1], h+1)
	}

	traverse(root, 1)

	return height
}
