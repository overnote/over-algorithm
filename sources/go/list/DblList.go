/**
 * 双向循环链表
 */
package list

import "fmt"

// 双向循环连表结点类型
type DblNode struct {
	data		interface{}
	prev		*DblNode
	next		*DblNode
}

// 双向循环链表
type DblList struct {
	head		*DblNode
	size		int
}

// 构造Node
func newDblNode(e interface{}) *DblNode{
	return &DblNode {
		data: e,
		prev: nil,
		next: nil,
	}
}

// 构造循环链表
func NewDblList() *DblList{
	p := newDblNode(0)
	p.next = p 	// 循环链表！！！
	p.prev = p
	return &DblList{
		head: p,
		size: 0,
	}
}

// 增：插入节点。约定只能在头节点之后插入，且不超过最大元素个数位置
func (l *DblList)Insert(e interface{}, index int){
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
	q := newDblNode(e)
	q.prev = p
	q.next = p.next

	// 执行插入
	p.next = q
	l.size++
}

// 删：根据位置删除，返回被删除的元素
func (l *DblList)Delete(index int) interface{}{ 

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

	// 获取删除元素数据
	q := p.next
	tempData := q.data

	// 执行删除
	q.next.prev = p
	p.next = q.next

	l.size--
	return tempData
}

// 改
func (l *DblList)Update(index int, e interface{}){
	p := l.Locate(index)
	if p == nil {
		return
	}
	p.data = e
}

// 查
func (l *DblList)Search(e interface{}) *DblNode{
	p := l.head
	for p.next != l.head {
		if p.data == e {
			break
		}
		p = p.next
	}
	if p.data == e {
		return p
	} else {
		return nil
	}
}

// 定位
func (l *DblList)Locate(index int) *DblNode{
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

// 根据值快速找到前驱
func (l *DblList)PrevElem(e interface{}) *DblNode{
	p := l.Search(e)
	if p == nil {
		fmt.Println("未找到该元素")
		return nil
	}
	return p.prev
}

// 根据值快速找到后继
func (l *DblList)NextElem(e interface{}) *DblNode{
	p := l.Search(e)
	if p == nil {
		fmt.Println("未找到该元素")
		return nil
	}
	return p.next
}

// 获取表长度
func (l *DblList)Length() int {
	return l.size
}

// 清空表：仅保留头节点
func (l *DblList)Clear() {
	l.head.next = l.head
	l.head.prev = l.head
	l.size = 0
}

// 显示循环链表
func (l *DblList)Display(){
	if l.size == 0 {
		fmt.Println("空链表")
		return
	}

	p := l.head
	pos := 0
	for p.next != nil {
		if pos == l.size {
			fmt.Print(p.data, "<->", l.head.data, "<->...\n", )
			break
		}
		fmt.Print(p.data, "<->")
		p = p.next
		pos++
	}
}