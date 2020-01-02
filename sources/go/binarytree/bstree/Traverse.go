package bstree

import (
	"container/list"
	"fmt"
)

// 对于所有的树来说，遍历都是通用的

// 前序遍历
func PreOrderTraverse(node *Node) {
	if node == nil {
		return
	}
	fmt.Printf("%v ", node.data)			// 前序遍历就是从node开始遍历，所以要先打印
	PreOrderTraverse(node.left)
	PreOrderTraverse(node.right)
}

// 中序遍历
func InOrderTraverse(node *Node) {
	if node == nil {
		return
	}
	// 会产生式升序结果
	InOrderTraverse(node.left)
	fmt.Printf("%v ", node.data)
	InOrderTraverse(node.right)

	// 会产生降序结果
	//InOrderTraverse(node.right)
	//fmt.Printf("%v ", node.data)
	//InOrderTraverse(node.left)
}

// 后序遍历
func PostOrderTraverse(node *Node) {
	if node == nil {
		return
	}
	PostOrderTraverse(node.left)
	PostOrderTraverse(node.right)
	fmt.Printf("%v ", node.data)
}

// 层序遍历
/**
	实现思路：无法用递归实现
	1 将各节点入队
	2 循环执行以下操作，直到队列为空
		取出队头节点出队，进行访问
		将队头节点的左子节点入队
		将队头节点的右子节点入队
 */
func LevelOrderTraverse(node *Node) {

	if node == nil {
		return
	}

	queue := list.New()		// 制作一个队列
	queue.PushBack(node)

	for queue.Len() != 0 {
		queueHead := queue.Remove(queue.Front())	// 队首出队
		tempNode := queueHead.(*Node)				// 类型断言
		fmt.Printf("%v ", tempNode.data)
		if tempNode.left != nil {
			queue.PushBack(tempNode.left)
		}
		if tempNode.right != nil {
			queue.PushBack(tempNode.right)
		}
	}

}