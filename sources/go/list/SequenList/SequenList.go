/**
 * 顺序表
 */

package SequenList

import (
	"errors"
	"fmt"
) 

// 线性表结构体对象		由于Go没有泛型，这里很难实现：调用者传递一个数组参数的方式New一个SequenList，只能定死如下一个整型数组
type SequenList struct {
	size int									// 该线性表最大容量
	length int									// 该线性表最大长度
	data [10]int								// 线性表内数据，这里为了演示默认设置为10长度的int数组，所有元素默认为0(Go的0值机制)
}

// 创建线性表实例    按笔者认为这里应该传入一个泛型数组，通过泛型数组来更高抽象顺序表	
func New() *SequenList {
	
	var arr [10]int

	return &SequenList{
		size: 10,
		length: 0,
		data: arr,
	}
}

// 打印线性表
func (sl *SequenList)Show() {
	fmt.Println(sl)
}

// 插入元素：从末尾append一个数据
func (sl *SequenList)Append(data int)  error{

	// 判断空间是否已满
	if sl.IsFull() {
		return errors.New("SequenList overflow")
	}

	sl.data[sl.length] = data
	sl.length++

	return nil
}

// 插入元素：任意位置插入元素
func (sl *SequenList)Insert(index int, data int) error {

	if sl.IsFull() {
		return errors.New("SequenList overflow")
	}

	if index < 0 || index > sl.length {
		return errors.New("index overflow")
	}

	// 这里如果按照正序循环则书写极其麻烦,从最后一位开始往后移动很简便
	for i := sl.length; i >= index; i-- {	

		if i == sl.length {					// 如果是在末尾插入 时间复杂度为O(1)
			sl.data[i] = data
			break
		}

		sl.data[i] = sl.data[i - 1]
	}

	sl.length++
	return nil
}

// 删除元素：从末尾pop一个数据
func (sl *SequenList)Pop() (int, error) {

	if sl.IsEmpty() {
		return 0, errors.New("SequenList is empty")
	}

	e := sl.data[sl.length - 1]
	sl.data[sl.length - 1] = 0
	sl.length --
	return e, nil
}

// 获取顺序表长度
func (sl *SequenList)Length() int{
	return sl.length
}

// 判断顺序表是否已满
func (sl *SequenList)IsFull() bool {
	if sl.length == sl.size {
		return true
	}
	return false
}

// 判断顺序表是否为空方法
func (sl *SequenList)IsEmpty() bool {
	if sl.length == 0 {
		return true
	}
	return false
}

// 获取顺序表容量
func (sl *SequenList)Size() int{
	return sl.size
}
