/*
 *  双向循环链表
 */
package DoubleList

import (
	"errors"
	"fmt"
)

type ElemeType int // 给 int起个别名

// 单链表 结点结构体
type Node struct {
	data ElemeType
	next *Node
	prev *Node
}

// 单链结构体
type DoubleList struct {
	length int // 该线性表最大长度
	first  *Node
}

// 构造表
func NewDoubleList() *DoubleList {
	return &DoubleList{
		length: 0,
		first:  nil,
	}
}

// 获取长度
func (l *DoubleList) Length() int {
	return l.length
}

// 显示表
func (l *DoubleList) Display() {

	if l.length == 0 {
		fmt.Println("数据结构内无元素")
		return
	}

	currentNode := l.first
	fmt.Printf("链表数据元素为：%d ", currentNode.data)
	for currentNode.next != l.first {
		currentNode = currentNode.next
		fmt.Printf("  %d ", currentNode.data)
	}
	fmt.Printf(" \n ")

}

func (l *DoubleList) Append(e ElemeType) {

	// 构造要插入的节点
	insertNode := &Node{
		data: e,
		next: nil,
		prev: nil,
	}

	// 第一次追加
	if l.length == 0 {
		insertNode.next = insertNode
		insertNode.prev = insertNode
		l.first = insertNode
		l.length++
		return
	}

	currentNode := l.first
	for currentNode.next != l.first {
		currentNode = currentNode.next
	}
	currentNode.next = insertNode
	insertNode.prev = currentNode
	insertNode.next = l.first
	l.first.prev = insertNode
	l.length++
}

// 插入元素
func (l *DoubleList) Insert(index int, e ElemeType) error {

	if index < 1 || index > l.length+1 {
		return errors.New("位序不正确")
	}

	// 构造要插入的节点
	insertNode := &Node{
		data: e,
		next: nil,
		prev: nil,
	}

	// 查找插入位置
	i := 1
	currentNode := l.first
	for currentNode.next != l.first {
		if i == index {
			break
		}
		i++
		currentNode = currentNode.next
	}

	// 执行插入
	prevNode := currentNode.prev
	insertNode.prev = prevNode
	insertNode.next = currentNode

	prevNode.next = insertNode
	currentNode.prev = insertNode
	l.length++

	// 考虑头指针是否改变
	if index == 1 {
		l.first = insertNode
	}

	return nil

}

// 删除 按照位序删除
func (l *DoubleList) Delete(index int) error {

	if index < 1 || index > l.length {
		return errors.New("位序不正确")
	}

	// 如果删除的是第一个元素
	if index == 1 && l.length == 1 {
		l.first = nil
		l.length = 0
		return nil
	}

	// 找到要删除元素的元素
	i := 1
	currentNode := l.first
	for currentNode.next != l.first {
		if i == index {
			break
		}
		i++
		currentNode = currentNode.next
	}

	// 执行删除
	prevNode := currentNode.prev
	nextNode := currentNode.next
	prevNode.next = nextNode
	nextNode.prev = prevNode
	l.length--

	// 考虑头指针是否会改变
	if index == 1 {
		l.first = nextNode
	}

	return nil

}

// 修改 按照位序修改
func (l *DoubleList) Update(index int, e ElemeType) error {

	if index < 1 || index > l.length {
		return errors.New("位序不正确")
	}

	i := 1
	currentNode := l.first
	for currentNode.next != l.first {
		if i == index {
			break
		}
		i++
		currentNode = currentNode.next
	}

	currentNode.data = e
	return nil
}

// 查询 按照位序查询值
func (l *DoubleList) GetElem(index int) (e ElemeType, err error) {

	if index < 1 || index > l.length {
		err = errors.New("位序不合法")
		return
	}

	i := 1
	currentNode := l.first
	for currentNode.next != l.first {
		if i == index {
			break
		}
		i++
		currentNode = currentNode.next
	}
	e = currentNode.data
	return
}

// 查询 按照值查询位序
func (l *DoubleList) LocateElem(e ElemeType) (index int, err error) {
	i := 1
	currentNode := l.first
	for currentNode.next != l.first {
		if currentNode.data == e {
			break
		}
		i++
		currentNode = currentNode.next
	}

	if i == l.length && currentNode.data != e {
		err = errors.New("未找到元素")
		return
	}

	index = i
	return
}

// 查询前驱
func (l *DoubleList) PriorElem(e ElemeType) (pe ElemeType, err error) {

	if l.length <= 1 {
		err = errors.New("数据结构元素不足，无法查询")
		return
	}

	if l.first.data == e {
		err = errors.New("首元素无前驱")
		return
	}

	i := 1
	currentNode := l.first
	for currentNode.next != l.first {
		if currentNode.next.data == e {
			pe = currentNode.data
			return
		}
		i++
		currentNode = currentNode.next
	}

	err = errors.New("未找到传入元素")
	return
}

// 查询后继
func (l *DoubleList) NextElem(e ElemeType) (ne ElemeType, err error) {

	if l.length <= 1 {
		err = errors.New("数据元素不足，无法查询")
		return
	}

	i := 1
	currentNode := l.first
	for currentNode.next != l.first {
		if currentNode.data == e {
			break
		}
		i++
		currentNode = currentNode.next
	}

	if i == l.length && currentNode.data != e {
		err = errors.New("未找到传入元素")
		return
	}

	if i == l.length && currentNode.data == e {
		err = errors.New("最后元素无后继")
		return
	}

	ne = currentNode.next.data
	return
}

// 清空
func (l *DoubleList) Clear() {
	l.first = nil
	l.length = 0
}
