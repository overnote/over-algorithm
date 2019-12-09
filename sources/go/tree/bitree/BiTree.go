/*
	二叉树
	对于一些节点如果不存在值，可以用 字符 # 代替，成为一个虚节点，那么每次创建的二叉树都是一个完整二叉树了
	笔者这里额外添加了length属性，引入的数据结构设计的复杂性，如果要简单实现插入节点，可以借鉴02节中的的排序算法
 */

package bitree

import (
	"fmt"
)

type node struct {
	data	interface{}
	lchild	*node
	rchild	*node
}

type BiTree struct {
	root	*node
	length  int
}

func NewBiTree() *BiTree{
	fmt.Println("注意：空节点使用 # 字符代替")
	root := new(node)
	t := &BiTree{
		root:   root,
		length: 0,
	}
	t.createNode(root)
	return t
}

// 构造数据
func (t *BiTree)createNode(n *node){
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
func (t *BiTree)Length() int{
	return t.length
}

// 获取根节点数据
func (t *BiTree)Root() interface{} {
	return t.root.data
}

// 前序遍历
func (t *BiTree)PreOrderTraverse() {
	t.preOrderRecursion(t.root)
	fmt.Println()
}

// 前序遍历的递归算法
func (t *BiTree)preOrderRecursion(n *node) {

	if n == nil {
		return
	}

	if n.data == "#" {
		return
	}

	fmt.Printf("%v ", n.data)

	t.preOrderRecursion(n.lchild)
	t.preOrderRecursion(n.rchild)
}