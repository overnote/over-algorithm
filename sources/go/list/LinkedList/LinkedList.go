/**
 * 单链表
 */

package list

import (
	"errors"
	"fmt"
)

// 节点对象，存储单链表上某个节点数据
type node struct {
	data interface{}			// 数据域
	next *node					// 指针域
}

// 单链表对象：存储头节点即可，当然也有做法是直接将头结点视为单链表对象
type LinkedList struct {
	head *node					
}

// 创建单链表
func New() *LinkedList {
	head := &node{0, nil}		// 头节点存储链表中元素的个数
	return &LinkedList{
		head,
	}
}

// 判断空
func (ll *LinkedList)IsEmpty() bool {
	return ll.head.next == nil
}

// 增加：末尾添加
func (ll *LinkedList) Append(data interface{}){

	insertNode := &node{data, nil}					// 要插入的节点
	var len int = 0
	len = ll.head.data.(int)

	// 查询最后一个节点
	lastNode := ll.head.next
	if lastNode == nil {							// 第一次添加
		ll.head.next = insertNode
		len ++
		ll.head.data = len
		return
	}

	for lastNode.next != nil {						// 不是第一次添加
		lastNode = lastNode.next
	}
	lastNode.next = insertNode
	len ++
	ll.head.data = len

	return
}

// 增加：任意位置插入结点
func (ll *LinkedList) Insert(index int, data interface{}) error{

	var len int = 0
	len = ll.head.data.(int)

	if index < 1 || index > len {
		return errors.New("index overflow")
	}

	beforeNode := ll.head
	appendNode := &node{data, nil}

	for i := 0; i < index - 1; i++ {
		beforeNode = beforeNode.next		// 找到要插入的位置的前置元素
	}

	appendNode.next = beforeNode.next
	beforeNode.next = appendNode

	len ++
	ll.head.data = len

	return nil

}

// 删除：删除指定位置结点
func (ll *LinkedList) Delete(index int) (interface{}, error) {

	var len int = 0
	len = ll.head.data.(int)

	if index < 0 || index >= len {
		return nil,errors.New("index overflow")
	}

	currentNode := ll.head
	var beforeNode *node
	var delData interface{}					// 被删除的数据

	for i := 0; i < index; i++ {
		beforeNode = currentNode
		currentNode = currentNode.next
	}

	beforeNode.next = currentNode.next
	currentNode = nil

	len--
	ll.head.data = len

	return delData, nil
}

// 查询： 获取指定位置结点
func (ll *LinkedList) Node(index int) (interface{}, error) {

	var len int = 0
	len = ll.head.data.(int)

	if index < 0 || index >= len {
		return nil, errors.New("index overflow")
	}
	currentNode := ll.head
	for i := 0; i < index; i++ {
		currentNode = currentNode.next
	}
	return currentNode.data, nil
}

// 清空链表
func (ll *LinkedList) Clear() {
	currentNode := ll.head.next
	for currentNode != nil {
		temp := currentNode.next
		currentNode = nil
		currentNode = temp
	}
	ll.head.data = 0
	ll.head.next = nil
}

// 打印链表
func (ll *LinkedList) Show() {
	var len int = 0
	len = ll.head.data.(int)
	currentNode := ll.head
	for i := 0; i <= len; i++ {
		fmt.Print(currentNode.data)
		if i == len {
			fmt.Print(" \n")
		} else {
			fmt.Print(" ")
		}

		currentNode = currentNode.next
	}
}