package bstreeNoParent

// 节点设计中 无父节点的二叉搜索树
type node struct {
	data	int
	left	*node
	right 	*node
}

type BSTree struct {
	root	*node
	length 	int
}

func NewBSTree() *BSTree{
	return &BSTree{
		root:   nil,
		length: 0,
	}
}

// 删除节点：不利用parent属性
func (t *BSTree)RemoveBySelf(data int) bool {

	// 寻找要删除的节点
	currentNode := t.root		// 要删除节点
	var parentNode *node		// 要删除节点的父节点
	var isLeft bool				// 要删节点与父节点关系
	for currentNode.data != data {

		if currentNode == nil {
			return false					// 节点不存在
		} else {
			parentNode = currentNode
		}

		if data < currentNode.data {		// 往左走
			isLeft = true
			currentNode = currentNode.left
		} else {							// 往右走
			isLeft = false
			currentNode = currentNode.right
		}

	}

	// 删除的节点是叶节点
	if currentNode.left == nil && currentNode.right == nil {
		// 如果树只有根节点
		if currentNode == t.root {
			t.root = nil
			return true
		}
		// 如果该叶节点有父节点，此时需要判断当前节点是父节点的左子节点还是右子节点
		if isLeft {
			parentNode.left = nil
		} else {
			parentNode.right = nil
		}
		return true
	}

	// 删除的节点只有一个子节点-左子节点
	if currentNode.right == nil {
		if currentNode == t.root {
			t.root = currentNode.left
		} else if isLeft {
			parentNode.left = currentNode.left
		} else {
			parentNode.right = currentNode.left
		}
		return true
	}
	// 删除的节点只有一个子节点-右子节点
	if currentNode.left == nil {
		if currentNode == t.root {
			t.root = currentNode.right
		} else if isLeft {
			parentNode.left = currentNode.right
		} else {
			parentNode.right = currentNode.right
		}
		return true
	}

	// 删除的节点有两个子节点：可以通过获取前驱、后继的办法，这里采用获取后继后提升后继的办法
	sur := t.removeBySur(currentNode)
	if currentNode == t.root {
		t.root = sur
	} else if isLeft {
		parentNode.left = sur
	} else {
		parentNode.right = sur
	}
	sur.left = currentNode.left

	return true
}

// 使用查找后继方式删除节点
func (t *BSTree)removeBySur(n *node) *node{
	sur := n
	currentNode := n.right
	surParent := n
	for currentNode != nil{
		surParent = sur
		sur = currentNode
		currentNode = currentNode.left
	}
	// 判断寻找的后继节点是否直接就是要删除的节点
	if sur != n.right {
		surParent.left = sur.right
		sur.right = n.right
	}
	return sur
}
