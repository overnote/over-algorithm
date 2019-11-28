/**
 *  顺序表：Go拥有切片数据结构，可简单视为动态数组，可以隐藏数组元素移动的底层细节
 */

package sqlist

import (
	"errors"
	"fmt"
)

const ListMaxSize int = 100  // 最大容量
const ListIncrement int = 10 // 每次容量增长幅度

type ElemeType int // 给 int起个别名

// 顺序表结构体
type SqList struct {
	size   int         // 该线性表最大容量
	length int         // 该线性表最大长度
	arr    []ElemeType // 线性表内数据，在Go中是一个切片
}

// 构造实例
func NewSqList() *SqList {
	return &SqList{
		size:   ListMaxSize,
		length: 0,
		arr:    nil,
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
	if l.length == 0 {
		fmt.Printf("为空表\n")
	} else {
		for i := 0; i < l.length; i++ {
			fmt.Printf("元素[%d] = %d\n", i+1, l.arr[i])
		}
	}
}

// 追加元素：在末尾追加元素
func (l *SqList) Append(e ElemeType) {

	if l.length == l.size {

		fmt.Println("数据结构已满，申请新空间")
		oldElem := l.arr
		newElem := make([]ElemeType, ListMaxSize+ListIncrement)
		for i := 0; i < len(oldElem); i++ {
			newElem[i] = oldElem[i]
		}
		l.size += ListIncrement
		l.arr = newElem

	}
	l.arr[l.length] = e
	l.length++
}

// 插入元素：任意位置插入元素
func (l *SqList) Insert(index int, e ElemeType) error {

	// 如果是空表
	if l.length == 0 && index == 1 {
		l.arr[0] = e
		l.length++
		return nil
	}

	// 如果索引越界
	if index < 1 || index > l.length+1 {
		return errors.New("位序越界")
	}

	if l.length == l.size {

		fmt.Println("数据结构已满，申请新空间")
		oldElem := l.arr
		newElem := make([]ElemeType, ListMaxSize+ListIncrement)
		for i := 0; i < len(oldElem); i++ {
			newElem[i] = oldElem[i]
		}
		l.size += ListIncrement
		l.arr = newElem

	}

	for i := l.length + 1; i > index; i-- {
		l.arr[i-1] = l.arr[i-2]
	}
	l.arr[index-1] = e
	l.length++
	return nil
}

// 删除元素：按照位序删除
func (l *SqList) Delete(index int) error {

	if index < 1 || index > l.length {
		return errors.New("索引越界")
	}

	for i := index - 1; i < l.length-1; i++ {
		l.arr[i] = l.arr[i+1]
	}
	l.length--

	return nil
}

// 修改：按照位序修改
func (l *SqList) Update(index int, e ElemeType) error {
	if index < 1 || index > l.length {
		return errors.New("位序越界")
	}
	l.arr[index-1] = e
	return nil
}

// 查询：按照位序查询值
func (l *SqList) GetElem(index int) (e ElemeType, err error) {
	if index < 1 || index > l.length {
		err = errors.New("位序不合法")
	}
	e = l.arr[index-1]
	return
}

// 查询：按值查询位序
func (l *SqList) LocateElem(e ElemeType) (index int, err error) {

	err = errors.New("未找到")

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
func (l *SqList) PriorElem(e ElemeType) (pe ElemeType, err error) {
	if l.length <= 1 {
		err = errors.New("顺序表元素不足，无法查询")
		return
	}

	if l.arr[0] == e {
		err = errors.New("第一个元素无前驱")
		return
	}

	// 查找 e 所在位置
	index, err := l.LocateElem(e)
	if err != nil {
		return
	}

	pe = l.arr[index-2]
	return
}

// 查询后继
func (l *SqList) NextElem(e ElemeType) (ne ElemeType, err error) {

	if l.length <= 1 {
		err = errors.New("数据结构为空，无法查询")
		return
	}
	if l.arr[l.length-1] == e {
		err = errors.New("最后一个元素无后继")
		return
	}

	// 查找 e 所在位置
	index, err := l.LocateElem(e)
	if err != nil {
		return
	}
	ne = l.arr[index]
	return
}

// 清空
func (l *SqList) Clear() {
	l.length = 0
	l.arr = make([]ElemeType, l.Size())
}
