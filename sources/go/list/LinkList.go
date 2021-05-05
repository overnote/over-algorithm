/**
 * 单链表
 */
package list

import (
	"fmt"
)

// 单链表结点类型
type Node struct {
	data	interface{}
	next 	*Node
}

// 单链表
type LinkList struct {
	head	*Node
	length  int
}

// 构造Node
func newNode(e interface{}) *Node{
	return &Node {
		data: e,
		next: nil,
	}
}

// 构造单链表
func NewLinkList() *LinkList {
	p := newNode(0)
	return &LinkList{
		head: p,
		length: 0,
	}
}

// 增：插入结点
// 约定：带头结点的链表，插入时，只能在头结点之后插入，也不允许插入超过最大元素个数的位置
func (l *LinkList)Insert(e interface{}, index int){
	
	if index < 1 || index > l.length + 1 {
		fmt.Println("插入位置非法")
		return
	}

    // 找到插入位置前一个位置：也可以使用 Locate函数
	p := l.head
	k := 0
	for p.next != nil && k < index - 1 {
		p = p.next
		k++
	}

    // 创建要插入的结点
	q := newNode(e)
	q.next = p.next
	p.next = q
	
	l.length++
}

// 删
func (l *LinkList)Delete(index int) interface{}{

	if index < 1 || index > l.length {
		fmt.Println("删除位置非法")
		return nil
	}

    // 找到删除位置前一个位置：也可以使用 Locate函数
	p := l.head
	k := 0
	for p.next != nil && k < index - 1 {
		p = p.next
		k++
	}

	// 执行删除
	tempData := p.next.data
	p.next = p.next.next

	l.length--
	return tempData
}

// 改
func (l *LinkList)Update(index int, e interface{}){
	// 找到修改位置：这里使用 locate 函数
	p := l.Locate(index)
	if p == nil {
		return
	}
	p.data = e
}

// 查
func (l *LinkList)Search(e interface{}) *Node{
	p := l.head
	for p.next != nil{
		if p.data == e {
			return p
		}
		p = p.next
	}
	return nil
}

// 定位：根据位置查询结点地址
func (l *LinkList)Locate(index int) *Node{
	if index < 0 || index > l.length + 1{
		fmt.Println("获取位置不合法")
		return nil
	}

	p := l.head
	k := 0
	for p.next != nil && k < index {
		p = p.next
		k++
	}
	return p
}

// 获取长度
func (l *LinkList)Length() int {
	return l.length	// 如果没有头结点一般使用循环获取长度
}

// 清空表：仅保留头结点
func (l *LinkList)Clear(){
	l.length = 0
	l.head.next = nil
}

// 显示单链表
func (l *LinkList)Display(){
	if l.length == 0 {
		fmt.Println("空链表")
		return
	}

	p := l.head
	pos := 0
	for p != nil {
		if pos == l.length{ // 最后一位
			fmt.Println(p.data)
			break
		} else {
			fmt.Print(p.data, "->")
			p = p.next
			pos++
		}
	}
}