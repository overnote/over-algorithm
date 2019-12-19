package bstree

type node struct {
	data	int
	left	*node
	right	*node
}

type BSTree struct {
	root 	*node
	length	int
}

func NewBSTree() *BSTree{
	return &BSTree{
		root:   nil,
	}
}

// 插入方法
func (bt *BSTree)Insert(data int) {

	// 构造插入节点
	insertNode := &node{
		data:  data,
		left:  nil,
		right: nil,
	}

	// 判断根节点是否有值
	if bt.root == nil {
		bt.root = insertNode
		return
	}

	// 执行插入
	bt.insertRC(bt.root, insertNode)
}

// 插入时的递归函数
func (bt *BSTree)insertRC(current *node, insert *node){

	if insert.data < current.data {					// 向左查找
		if current.left == nil {
			current.left = insert
		} else {
			bt.insertRC(current.left, insert)
		}
	} else {										// 向右查找
		if current.right == nil {
			current.right = insert
		} else {
			bt.insertRC(current.right, insert)
		}
	}
}

// 遍历：二叉搜索树仍然是二叉树，可以直接使用先序、后序、中序、层序遍历

// 因为二叉搜索树获取最大值、最小很方便，所以需要提供最大值、最小值获取操作
// 获取最大值：二叉搜索树最右边的值
func (bt *BSTree)MaxData() interface{}{
	currentNode := bt.root
	var data interface{}
	for currentNode != nil {
		data = currentNode.data
		currentNode = currentNode.right
	}
	return data
}

// 获取最小值：二叉搜索树最左边的值
func (bt *BSTree)MinData() interface{}{
	currentNode := bt.root
	var data interface{}
	for currentNode != nil {
		data = currentNode.data
		currentNode = currentNode.left
	}
	return data
}

// 查找数据是否在二叉树
func (bt *BSTree)SearchData(data int) bool {
	currentNode := bt.root
	for currentNode != nil {
		if data < currentNode.data {
			currentNode = currentNode.left
		} else if data > currentNode.data{
			currentNode = currentNode.right
		} else {
			return true
		}
	}
	return false
}

/**
	删除节点：
	- 寻找要删除的节点，此时需确定节点是否存在
	- 被删除节点是叶节点（无子节点），此时需要把其父节点的left/right设置为null
	  - 额外考虑如果树只有一个根节点的情况
	- 该节点有一个子节点，此时需要让其父节点的left/right指向其子节点
	- 该节点有两个子节点
 */
func (bt *BSTree)Remove(data int) bool {

	// 寻找要删除的节点
	currentNode := bt.root		// 要删除节点
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
		if currentNode == bt.root {
			bt.root = nil
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
		if currentNode == bt.root {
			bt.root = currentNode.left
		} else if isLeft {
			parentNode.left = currentNode.left
		} else {
			parentNode.right = currentNode.left
		}
		return true
	}
	// 删除的节点只有一个子节点-右子节点
	if currentNode.left == nil {
		if currentNode == bt.root {
			bt.root = currentNode.right
		} else if isLeft {
			parentNode.left = currentNode.right
		} else {
			parentNode.right = currentNode.right
		}
		return true
	}

	// 删除的节点有两个子节点：可以通过获取前驱、后继的办法，这里采用获取后继后提升后继的办法
	sur := bt.removeBySur(currentNode)
	if currentNode == bt.root {
		bt.root = sur
	} else if isLeft {
		parentNode.left = sur
	} else {
		parentNode.right = sur
	}
	sur.left = currentNode.left

	return true
}

// 使用查找后继方式删除节点
func (bt *BSTree)removeBySur(n *node) *node{
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