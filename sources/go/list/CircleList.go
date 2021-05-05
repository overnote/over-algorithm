/**
 * 循环链表
 */

package list

import (
	"fmt"
)

// 循环链表结点类型
type CircleNode struct {
	data	interface{}
	next 	*CircleNode
}

// 循环链表
type CircleList struct {
	head	*CircleNode
	size  int
}

// 构造Node
func newCircleNode(e interface{}) *CircleNode{
	return &CircleNode {
		data: e,
		next: nil,
	}
}

// 构造循环链表
func NewCircleList() *CircleList{
	p := newCircleNode(0)
	p.next = p 	// 循环链表！！！
	return &CircleList{
		head: p,
		size: 0,
	}
}

// 增：插入结点
// 约定：带头结点的链表，插入时，只能在头结点之后插入，也不允许插入超过最大元素个数的位置
func (l *CircleList)Insert(e interface{}, index int){
	if index < 1 || index > l.size + 1 {
		fmt.Println("插入位置不合法")
		return
	}

	// 找到插入位置前一个位置：也可以使用 Locate函数
	p := l.head
	k := 0
	for p.next != l.head && k < index - 1 {
		p = p.next
		k++
	}

	// 创建要插入的节点
	q := newCircleNode(e)
	q.next = p.next
	p.next = q

	l.size++
}

// 删：根据位置删除，返回被删除的元素
func (l *CircleList)Delete(index int) interface{}{ 
	if index < 1 || index > l.size {
		fmt.Println("删除位置非法")
		return nil
	}

	// 找到删除位置前一个位置：也可以使用 Locate函数
	p := l.head
	k := 0
	for p.next != l.head && k < index - 1 {
		p = p.next
		k++
	}

	// 执行删除
	q := p.next
	tempData := q.data
	p.next = q.next

	l.size--
	return tempData
}

// 改
func (l *CircleList)Update(index int, e interface{}){
	p := l.Locate(index)
	if p == nil {
		return
	}
	p.data = e
}

// 查
func (l *CircleList)Search(e interface{}) *CircleNode{
	p := l.head
	for p.next != l.head {
		if p.data == e {
			return p
		}
		p = p.next
	}
	return nil
}

// 定位
func (l *CircleList)Locate(index int) *CircleNode{
	if index < 0 || index > l.size + 1 {
		fmt.Println("获取位置不合法")
		return nil
	}

	p := l.head
	k := 0
	for p.next != l.head && k < index {
		p = p.next
		k++
	}
	return p
}

// 获取表长度
func (l *CircleList)Length() int {
	return l.size
}

// 清空表：仅保留头节点
func (l *CircleList)Clear() {
	l.head.next = l.head
	l.size = 0
}

// 显示循环链表
func (l *CircleList)Display(){
	if l.size == 0 {
		fmt.Println("空链表")
		return
	}

	p := l.head
	pos := 0
	for p.next != nil {
		if pos == l.size {
			fmt.Print(p.data, "->", l.head.data, "->...\n", )
			break
		}
		fmt.Print(p.data, "->")
		p = p.next
		pos++
	}
}