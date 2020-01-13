/*
 *  双向循环链表
 */
package list

import (
	"errors"
	"fmt"
)

// 结点结构体
type DCircleNode struct {
	data 	interface{}
	next 	*DCircleNode
	prev 	*DCircleNode
}

// 单链结构体
type DcList struct {
	length 	int 			// 该线性表最大长度
	head  	*DCircleNode
}

// 构造表
func NewDcList() *DcList {
	return &DcList{
		length: 0,
		head:  nil,
	}
}

// 获取长度
func (l *DcList) Length() int {
	return l.length
}

// 显示表
func (l *DcList) Display() {

	if l.length == 0 {
		fmt.Println("数据结构内无元素")
		return
	}

	fmt.Printf("显示表：")
	currentNode := l.head
	for currentNode.next != l.head {
		fmt.Printf("%v ", currentNode.data)
		currentNode = currentNode.next
	}
	fmt.Printf("%v ", currentNode.data)
	fmt.Println()
}

func (l *DcList) Append(e interface{}) {

	// 构造要插入的节点
	insertNode := &DCircleNode{
		data: e,
		next: nil,
		prev: nil,
	}

	// 第一次追加
	if l.length == 0 {
		insertNode.next = insertNode
		insertNode.prev = insertNode
		l.head = insertNode
		l.length++
		return
	}

	// 常规追加：找到最后一个元素
	currentNode := l.head
	for currentNode.next != l.head {
		currentNode = currentNode.next
	}

	// 设置insertNode
	insertNode.next = l.head
	insertNode.prev = currentNode

	// 执行追加
	currentNode.next = insertNode
	l.head.prev = insertNode
	l.length++
}

// 插入元素
func (l *DcList) Insert(index int, e interface{}) error {

	if index < 1 || index > l.length + 1 {
		fmt.Println("位序不正确")
		return errors.New("位序不正确")
	}

	// 如果是末尾插入
	if l.length == 0 || l.length + 1 == index {
		l.Append(e)
		return nil
	}

	// 常规插入：构造要插入的节点
	insertNode := &DCircleNode{
		data: e,
		next: nil,
		prev: nil,
	}

	// 查找插入位置
	i := 1
	currentNode := l.head
	for currentNode.next != l.head {
		if i == index {
			break
		}
		i++
		currentNode = currentNode.next
	}
	
	// 设置insert
	insertNode.prev = currentNode.prev
	insertNode.next = currentNode
	
	// 执行插入
	currentNode.prev.next = insertNode
	currentNode.prev = insertNode
	l.length++
	
	// 考虑头指针是否改变
	if index == 1 {
		l.head = insertNode
	}

	return nil

}

// 删除 按照位序删除
func (l *DcList) Delete(index int) error {

	if index < 1 || index > l.length {
		fmt.Println("位序不正确")
		return errors.New("位序不正确")
	}

	// 如果删除的是第一个元素
	if index == 1 && l.length == 1 {
		l.head = nil
		l.length = 0
		return nil
	}

	// 找到要删除元素的元素
	i := 1
	currentNode := l.head
	for currentNode.next != l.head {
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
		l.head = nextNode
	}

	return nil

}

// 修改 按照位序修改
func (l *DcList) Update(index int, e interface{}) error {

	if index < 1 || index > l.length {
		fmt.Println("位序不正确")
		return errors.New("位序不正确")
	}

	i := 1
	currentNode := l.head
	for currentNode.next != l.head {
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
func (l *DcList) GetElem(index int) (e interface{}, err error) {

	if index < 1 || index > l.length {
		fmt.Println("位序不合法")
		err = errors.New("位序不合法")
		return
	}

	i := 1
	currentNode := l.head
	for currentNode.next != l.head {
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
func (l *DcList) Locate(e interface{}) (index int, err error) {
	i := 1
	currentNode := l.head
	for currentNode.next != l.head {
		if currentNode.data == e {
			break
		}
		i++
		currentNode = currentNode.next
	}

	if i == l.length && currentNode.data != e {
		fmt.Println("未找到元素")
		err = errors.New("未找到元素")
		return
	}

	index = i
	return
}

// 查询前驱
func (l *DcList) PrevElem(e interface{}) (pe interface{}, err error) {

	if l.length <= 1 {
		fmt.Println("数据结构元素不足，无法查询")
		err = errors.New("数据结构元素不足，无法查询")
		return
	}

	if l.head.data == e {
		fmt.Println("首元素无前驱")
		err = errors.New("首元素无前驱")
		return
	}

	i := 1
	currentNode := l.head
	for currentNode.next != l.head {
		if currentNode.next.data == e {
			pe = currentNode.data
			return
		}
		i++
		currentNode = currentNode.next
	}

	fmt.Println("数据结构元素不足，无法查询")
	err = errors.New("未找到传入元素")
	return
}

// 查询后继
func (l *DcList) NextElem(e interface{}) (ne interface{}, err error) {

	if l.length <= 1 {
		fmt.Println("数据结构元素不足，无法查询")
		err = errors.New("数据结构元素不足，无法查询")
		return
	}

	i := 1
	currentNode := l.head
	for currentNode.next != l.head {
		if currentNode.data == e {
			break
		}
		i++
		currentNode = currentNode.next
	}

	if i == l.length && currentNode.data != e {
		fmt.Println("未找到传入元素")
		err = errors.New("未找到传入元素")
		return
	}

	if i == l.length && currentNode.data == e {
		fmt.Println("末尾元素无后继")
		err = errors.New("末尾元素无后继")
		return
	}

	ne = currentNode.next.data
	return
}

// 清空
func (l *DcList) Clear() {
	l.head = nil
	l.length = 0
}
