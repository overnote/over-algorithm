/**
 * 栈:链式栈
 */

package LinkedStack

import "fmt"

type Node struct {
	data interface{}
	next *Node
}

type LinkedStack struct {
	Top    *Node // 栈顶元素
	Length int
}

func New() *LinkedStack {
	return &LinkedStack{
		nil,
		0,
	}
}

// 压栈
func (ls *LinkedStack) Push(data interface{}) {
	node := new(Node)
	node.data = data
	node.next = ls.Top
	ls.Top = node
	ls.Length++
}

// 出栈
func (ls *LinkedStack) Pop() interface{} {
	if ls.Length == 0 {
		fmt.Println("stack is empty")
		return nil
	}
	value := ls.Top.data
	node := ls.Top
	ls.Top = node.next
	ls.Length--
	return value
}
