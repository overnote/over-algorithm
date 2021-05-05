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
func (l *LinkList)Insert(e interface{}, index int) error{
	
	// 找到其前一个元素位置
	p := l.Locate(index - 1)
	if p == nil {
		return nil
	}

    // 创建要插入的结点
	temp := newNode(e)
	temp.next = p.next
	p.next = temp
	
	l.data = l.data.(int) + 1
	return nil
}

// 删
func (l *LinkList)Delete(index int) interface{}{
	// 找到前一个元素位置
	p := l.Locate(index - 1)
	if p == nil || p.next == nil{
        return nil;
    }
	// 执行删除
	tempData := p.next.data
	p.next = p.next.next
	l.data = l.data.(int) - 1
	return tempData
}

// 改
func (l *LinkList)Update(index int, e interface{}){
	p := l.Locate(index)
	if p == nil {
		return
	}
	p.data = e
}

// 查
func (l *LinkList)Search(e interface{}) *Node{
	p := l.next
	for p != nil && p.data != e {
		p = p.next
	}
	return p
}

// 定位：根据位置查询结点地址
func (l *LinkList)Locate(index int) *Node{
	if index < 0 || index > l.data.(int) + 1{
		fmt.Println("获取位置不合法")
		return nil
	}

	k := 0
	p := l
	for p != nil && k < index {
		p = p.next
		k++
	}
	return p
}

// 获取长度
func (l *LinkList)Length() int {
	return l.data.(int)	// 如果没有头结点一般使用循环获取长度
}

// 清空表：仅保留头结点
func (l *LinkList)Clear(){
	l.data = 0
	l.next = nil
}

// 显示单链表
func (l *LinkList)Display(){
	if l.data.(int) == 0 {
		fmt.Println("空链表")
		return
	}

	p := l
	pos := 0
	for p != nil {
		if pos == l.data.(int) { // 最后一位
			fmt.Println(p.data)
			break
		} else {
			fmt.Print(p.data, "->")
			p = p.next
			pos++
		}
	}
}