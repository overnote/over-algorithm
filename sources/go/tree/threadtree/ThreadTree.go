/*
	二叉线索树：
		仍然使用值 # 表示值为空虚节点
*/

package threadtree

import (
	"fmt"
)

type node struct {
	data	interface{}
	lchild	*node
	rchild	*node
	Ltag	int
	Rtag 	int
}

type ThreadTree struct {
	root	*node
	length  int
}

// 临时节点：用于记录在某个节点进行递归循环时，其前一个循环的节点
var prevNode *node

func NewThreadTree() *ThreadTree{
	fmt.Println("注意：空节点使用 # 字符代替")
	root := new(node)
	t := &ThreadTree{
		root:   root,
		length: 0,
	}
	t.createNode(root)
	return t
}

// 构造数据
func (t *ThreadTree)createNode(n *node){
	var e string
	fmt.Scanf("%v", &e)
	n.data = e							// 给当前节点赋值
	if e == "#" {
		return
	}
	t.length++
	n.lchild = new(node)
	n.rchild = new(node)
	t.createNode(n.lchild)				// 构造左子树
	t.createNode(n.rchild)				// 构造右子树
}

// 获取长度
func (t *ThreadTree)Length() int{
	return t.length
}

// 获取根节点数据
func (t *ThreadTree)Root() interface{} {
	return t.root.data
}

// 中序遍历 线索化二叉树
func (t *ThreadTree)InOrderThread() {
	t.inOrderThreadRecursion(t.root)
	fmt.Println()
}

// 中序遍历具体递归过程
func (t *ThreadTree)inOrderThreadRecursion(n *node) {

	if n == nil {
		return
	}

	// 处理左子树
	t.inOrderThreadRecursion(n.lchild)

	// 如果左指针为空，则将左指针指向前驱节点
	if n.lchild == nil {
		n.lchild = prevNode
		n.Ltag = 1
	}

	// 前一个节点的后继节点指向当前节点
	if prevNode != nil && prevNode.rchild == nil {
		prevNode.rchild = n
		prevNode.Rtag = 1
	}

	prevNode = n

	// 处理右子树
	t.inOrderThreadRecursion(n.rchild)

}

// 线索树生成后的遍历方式---找到最左节点后遍历方式
func (t *ThreadTree)InOrderTraverse() {
	root := t.root
	t.inOrderTraverseRecursion(root)
}

// 具体执行步骤
func (t *ThreadTree)inOrderTraverseRecursion(n *node) {

	// 找到最左子节点
	for n != nil && n.Ltag == 0 {
		n = n.lchild
	}

	// 开始遍历
	for n != nil {

		if n.data != "#" {
			fmt.Printf("%v ", n.data)
		}

		if n.Ltag == 1 {
			n = n.rchild
		} else {
			n = n.rchild
			for n != nil && n.Ltag == 0 {
				n = n.lchild
			}
		}
	}

	fmt.Println()
}