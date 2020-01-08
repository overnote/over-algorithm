package rbtree

const RED 	= true
const BLACK = false

// 节点对象
type node struct {
	data	int
	left	*node
	right 	*node
	parent	*node
	color 	bool				// 默认值为false：黑色
}

// 红树对象
type RBTree struct {
	root	*node
	length 	int
}

// 初始化红黑树
func NewBSTree() *RBTree{
	return &RBTree{
		root:   nil,
		length: 0,
	}
}