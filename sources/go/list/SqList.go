/**
*  顺序表：SqList
*	说明：Go拥有切片数据结构Slice，可简单视为动态数组，Slice提供了许多方法，会隐藏数组元素移动的底层细节
 */

package list

import (
	"fmt"
)

const SqListSize = 5

// 顺序表结构体
type SqList struct {
	items    	[]interface{} 	// 线性表内数据
	cap   		int         	// 该线性表最大容量
	length 		int         	// 该线性表元素个数
}

// 构造实例
func NewSqList() *SqList {
	return &SqList{
		items:		make([]interface{}, SqListSize),
		cap:   		SqListSize,
		length: 	0,
	}
}

// 增：在末尾追加元素
func (l *SqList)PushBack(e interface{}) {

	// 容量已满，申请新空间
	if l.length == l.cap {
		newArr := make([]interface{}, l.cap * 2)
		copy(newArr, l.items)			// 拷贝原有元素到新的切片中
		l.items = newArr
		l.cap = l.cap * 2
	}

	l.items[l.length] = e
	l.length++
}

// 增：任意位置插入元素
func (l *SqList) Insert(index int, e interface{}) bool {

	// 如果索引越界
	if index < 1 || index > l.length + 1 {
		fmt.Println("插入位置不正确")
		return false
	}

	// 如果是空表
	if l.length == 0 && index == 1 {		// index此时必定是1
		l.items[0] = e
		l.length++
		return true
	}

	// 容量已满，申请新空间
	if l.Length() == l.Cap() {
		fmt.Println("申请新容量")
		newArr := make([]interface{}, l.Cap() * 2)
		copy(newArr, l.items)			// 拷贝原有元素到新的切片中
		l.items = newArr
		l.cap = l.Cap() * 2
	}

	// 执行插入
	for i := l.length + 1; i > index; i-- {
		l.items[i - 1] = l.items[i - 2]
	}
	l.items[index - 1] = e
	l.length++
	return true
}

// 获取顺序表长度
func (l *SqList)Length() int {
	return l.length
}

// 获取顺序表容量
func (l *SqList)Cap() int {
	return l.cap
}

// 清空
func (l *SqList)Clear() {
	l.length = 0
	l.items = make([]interface{}, l.Cap())
}
