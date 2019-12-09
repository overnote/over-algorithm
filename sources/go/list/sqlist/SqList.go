/**
 *  顺序表：SqList
 *	说明：Go拥有切片数据结构Slice，可简单视为动态数组
 *	说明：Slice提供了许多方法，会隐藏数组元素移动的底层细节
 */

package sqlist

import (
	"errors"
	"fmt"
)					

// 顺序表结构体
type SqList struct {
	size   		int         	// 该线性表最大容量
	length 		int         	// 该线性表最大长度
	arr    		[]interface{} 	// 线性表内数据
	increment	int				// 线性表每次扩容的幅度
}

// 构造实例
func NewSqList(size int, increment int) *SqList {
	return &SqList{
		size:   	size,
		length: 	0,
		arr:    	make([]interface{}, size),
		increment:	increment,
	}
}

// 获取顺序表长度
func (l *SqList) Length() int {
	return l.length
}

// 获取顺序表容量
func (l *SqList) Size() int {
	return l.size
}

// 打印线性表
func (l *SqList) Display() {

	fmt.Printf("数据结构长度：%v，容量：%v\n", l.length, l.size)

	if l.length == 0 {
		return
	} 

	fmt.Printf("数据元素显示：")
	for i := 0; i <= l.length - 1; i++ {
		fmt.Printf("%v ", l.arr[i])
	}
	fmt.Println("")
}

// 增：在末尾追加元素
func (l *SqList) Append(e interface{}) {

	// 容量已满，申请新空间
	if l.length == l.size {
		fmt.Println("申请新容量")
		newArr := make([]interface{}, l.size + l.increment)
		copy(newArr, l.arr)			// 拷贝原有元素到新的切片中
		l.arr = newArr
		l.size += l.increment	
	}

	l.arr[l.length] = e
	l.length++
}

// 增：任意位置插入元素
func (l *SqList) Insert(index int, e interface{}) error {

	// 如果是空表
	if l.length == 0 && index == 1 {
		l.arr[0] = e
		l.length++
		return nil
	}

	// 如果索引越界
	if index < 1 || index > l.length + 1 {
		fmt.Println("位序越界")
		return errors.New("位序越界")
	}

	// 容量已满，申请新空间
	if l.length == l.size {
		fmt.Println("申请新容量")
		newArr := make([]interface{}, l.size + l.increment)
		copy(newArr, l.arr)			// 拷贝原有元素到新的切片中
		l.arr = newArr
		l.size += l.increment	
	}

	fmt.Println(cap(l.arr))

	// 执行插入
	for i := l.length + 1; i > index; i-- {
		l.arr[i - 1] = l.arr[i - 2]
	}
	l.arr[index - 1] = e
	l.length++
	return nil
}

// 删：按照位序删除
func (l *SqList) Delete(index int) error {

	if index < 1 || index > l.length {
		fmt.Println("位序越界")
		return errors.New("位序越界")
	}

	for i := index; i <= l.length; i++ {
		if i == l.length {
			l.arr[i] = nil
		} else {
			l.arr[i] = l.arr[i + 1]
		}	
	}
	l.length--

	return nil
}

// 修：按照位序修改
func (l *SqList) Update(index int, e interface{}) error {

	if index < 1 || index > l.length {
		fmt.Println("位序越界")
		return errors.New("位序越界")
	}

	l.arr[index - 1] = e
	return nil
}

// 查询：按照位序查询值
func (l *SqList) GetElem(index int) (e interface{}, err error) {
	if index < 1 || index > l.length {
		fmt.Println("位序越界")
		err = errors.New("位序越界")
		return
	}
	e = l.arr[index - 1]
	return
}

// 查询：按值查询位序
func (l *SqList) Locate(e interface{}) (index int, err error) {

	err = errors.New("未找到数据")

	for i := 0; i < l.length; i++ {
		if l.arr[i] == e {
			index = i + 1
			err = nil
			break
		}
	}
	return
}

// 查询前驱
func (l *SqList) PrevElem(e interface{}) (interface{}, error) {

	if l.length <= 1 {
		return nil, errors.New("数据结构为空")
	}

	// 查找 e 所在位置
	index, err := l.Locate(e)
	if err != nil {
		return nil, err
	}

	if index == 1 {
		fmt.Println("第一个元素无前驱")
		return nil, errors.New("第一个元素无前驱")
	}

	return l.arr[(index - 1) - 1], nil
}

// 查询后继
func (l *SqList) NextElem(e interface{}) (interface{}, error) {

	if l.length <= 1 {
		fmt.Println("数据结构为空")
		return nil, errors.New("数据结构为空")
	}

	// 查找 e 所在位置
	index, err := l.Locate(e)
	if err != nil {
		return nil, err
	}

	if index == l.length {
		fmt.Println("最后一个元素无后继")
		return nil, errors.New("最后一个元素无后继")
	}
 
	return l.arr[(index -1) + 1], nil
}

// 清空
func (l *SqList) Clear() {
	l.length = 0
	l.arr = make([]interface{}, l.size)
}
