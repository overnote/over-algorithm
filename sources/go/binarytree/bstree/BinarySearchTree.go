package bstree

/**
对于所有的不同类型的二叉树来说，其遍历的方式都是一致的
*/

// 将父节点也设计进去
type Node struct {
	data	int
	left	*Node
	right 	*Node
	parent	*Node
}

type BSTree struct {
	root	*Node
	length 	int
}

func NewBSTree() *BSTree{
	return &BSTree{
		root:   nil,
		length: 0,
	}
}

// 插入元素：迭代方式插入
func (t *BSTree)InsertByGen(data int) {

	// 如果树刚刚初始化
	if t.root == nil {
		root := &Node{
			data:   data,
			left:   nil,
			right:  nil,
			parent: nil,
		}
		t.root = root
		t.length++
		return
	}

	// 添加的不是根节点
	currentNode := t.root			// 循环到的当前节点
	parentNode := t.root			// 当前节点的父节点
	isLeft := true					// 当前是在左侧插入还是右侧插入

	// 创建插入节点
	insertNode := &Node{
		data:   data,
		left:   nil,
		right:  nil,
		parent: nil,
	}

	for currentNode != nil {		// 插入其实是直接往叶节点处插入，所以循环退出条件是到达叶节点后

		if insertNode.data == currentNode.data {
			currentNode = insertNode	// 万一以后要将data属性改变为自定义对象呢？所以这里要做一下相等处理
			return
		}

		parentNode = currentNode

		if insertNode.data > currentNode.data {		//  往右插入
			currentNode = currentNode.right
			isLeft = false
		} else {							// 往左插入
			currentNode = currentNode.left
			isLeft = true
		}
		

		// 执行插入
		insertNode.parent = parentNode
		if isLeft {
			parentNode.left = insertNode
		} else {
			parentNode.right = insertNode
		}
		t.length++

	}

}

// 插入元素：递归方式
func (t *BSTree)InsertByRc(data int) {

	// 构造插入节点
	insertNode := &Node{
		data:  data,
		left:  nil,
		right: nil,
		parent: nil,
	}

	// 执行递归插入
	t.insertRC(t.root, insertNode)
}

// 插入时的递归函数
func (t *BSTree)insertRC(currentNode *Node, insertNode *Node){

	// 判断根节点是否存在
	if t.root == nil {
		t.root = insertNode
		t.length++
		return
	}

	// 如果插入数据和当前节点数据相同
	if currentNode.data == insertNode.data {
		currentNode = insertNode
		return
	}

	parentNode := currentNode

	if currentNode.data > insertNode.data {		// 向左查找
		if currentNode.left == nil {
			insertNode.parent = parentNode
			currentNode.left = insertNode
			t.length++
		} else {
			parentNode = currentNode.left
			t.insertRC(currentNode.left, insertNode)
		}
	}

	if currentNode.data < insertNode.data { // 向右查找
		if currentNode.right == nil {
			insertNode.parent = parentNode
			currentNode.right = insertNode
			t.length++
		} else {
			parentNode = currentNode.right
			t.insertRC(currentNode.right, insertNode)
		}
	}
}



// 因为二叉搜索树获取最大值、最小很方便，所以需要提供最大值、最小值获取操作
// 获取最大值：二叉搜索树最右边的值
func (t *BSTree)MaxData() interface{}{
	currentNode := t.root
	var data interface{}
	for currentNode != nil {
		data = currentNode.data
		currentNode = currentNode.right
	}
	return data
}

// 获取最小值：二叉搜索树最左边的值
func (t *BSTree)MinData() interface{}{
	currentNode := t.root
	var data interface{}
	for currentNode != nil {
		data = currentNode.data
		currentNode = currentNode.left
	}
	return data
}

// 查找数据是否在二叉树
func (t *BSTree)SearchData(data int) bool {
	currentNode := t.root
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
func (t *BSTree)Remove(data int) bool {

	// 寻找要删除的节点
	currentNode := t.root		// 要删除节点
	var parentNode *Node		// 要删除节点的父节点
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
func (t *BSTree)removeBySur(n *Node) *Node{
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