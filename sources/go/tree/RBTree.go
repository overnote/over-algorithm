package tree

const RED 	= false				 // 默认值为false：红色，因为这样做可以尽快让树满足红黑树特质
const BLACK = true

// 节点对象
type node struct {
	data	int
	left	*node
	right 	*node
	parent	*node
	color 	bool
}

// 红树对象
type RBTree struct {ngth 	int
	root	*node
	le
}

// 初始化红黑树
func NewBSTree() *RBTree{
	return &RBTree{
		root:   nil,
		length: 0,
	}
}

// 添加元素
func (t *RBTree)Insert(data int) {

}