package SeqQueue

/*
链式队列
*/

import (
	"fmt"
)

// 队列结点
type node struct {
	data interface{}
	next *node
}

// 链式队列对象
type LinkedQueue struct {
	front *node
	rear *node
	length int
}

// 创建队列
func New() *LinkedQueue {
	return &LinkedQueue{
		nil,
		nil,
		0,
	}
}

// 新增
func (lq *LinkedQueue) Push(data interface{}) {
	insertNode := &node{
		data,
		nil,
	}
	lq.rear.next = insertNode
	lq.rear = insertNode
	lq.length++
}

// 删除
func (lq *LinkedQueue) Delete() interface{} {

	if lq.IsEmpty() {
		fmt.Println("queue is empty")
		return nil
	}

	node := lq.front.next
	lq.front.next = node.next
	lq.length--

	if node == lq.rear {
		lq.front = lq.rear
	}

	return node
}

// 判断队列是否为空
func (lq *LinkedQueue) IsEmpty() bool{
	if lq.front == lq.rear {
		return true
	}
	return false
}

// 获取队列长度
func (lq *LinkedQueue) Length() int {
	return lq.length
}

// 清空队列长度
func (lq *LinkedQueue) Clear() {
	if lq.IsEmpty() {
		return
	}
	lq.front = lq.rear
	lq.length = 0
}