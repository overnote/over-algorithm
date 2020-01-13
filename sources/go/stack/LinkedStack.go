/*
 *  链式栈
 */

package stack

import (
	"fmt"
)

// 链栈节点
type LinkedNode struct {
	data 		interface{}
	next		*LinkedNode
}

// 链栈结构体
type LinkedStack struct {
	top 		*LinkedNode				// 栈顶指针
	length		int						// 栈的长度
}

func NewLinkedStack() *LinkedStack{
	return &LinkedStack{
		top:	nil,
		length:	0,
	}
}

// 压栈操作
func (s *LinkedStack)Push(e interface{}) {
	node := &LinkedNode{
		data: e,
		next: s.top,
	}
	s.top = node
	s.length++
}

// 弹栈操作：返回出栈的值
func (s *LinkedStack)Pop() interface{}{

	if s.top == nil  {
		fmt.Println("空栈无法pop")
		return nil
	}

	e := s.top.data
	node := s.top
	s.top = node.next
	s.length--
	return e

}

// 栈顶元素
func (s *LinkedStack)Top() interface{}{
	if s.top == nil  {
		fmt.Println("空栈无栈顶元素")
		return nil
	}
	return s.top
}