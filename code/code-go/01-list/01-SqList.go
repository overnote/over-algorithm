/**
 * 顺序表：SqList
 * 为了简便，这里不考虑数组扩容，缩容问题，简单申请一个数组即可
 */

package list

import (
	"errors"
	"fmt"
)

// 顺序表结构体
type SqList struct {
	data    		[]any				// 使用切片结构，可以摆脱数组容量限制
	length 			int         	// 该顺序表元素个数
}

// 构造实例
func NewSqList() *SqList {
	return &SqList{
		data:		make([]any, 5),// 给一个初始默认长度（非容量）
		length: 	0,
	}
}


// 获取顺序表长度
func (l *SqList)Length() int{
	return l.length
}

// 清空
func (l *SqList)Clear(){
	l.length = 0
}

// 显示顺序表
func (l *SqList) Display() {
	fmt.Println(l.data[0:l.length])
}

// 增：按照一个位置插入数据
func (l *SqList)InsertElem(e any, idx int) error{

	if idx < 1 || idx > l.length + 1 {
		return errors.New("插入位置不合法")
	}

	// 将切片在插入位置一份为2，再合并，即可实现插入效果
	l.data= append(l.data[: idx - 1], append([]any{e}, l.data[idx - 1 :]...)...)
	l.length++
	return nil
}

// 删：删除位置上的元素，并获取到删除元素的值
func (l *SqList)DeleteElem(idx int) (t any, err error) {
	if idx < 1 || idx > l.length {
		err =  errors.New("删除位置不合法")
		return
	}

	t = l.data[idx - 1]

	// 也可以使用append方式
	for i := idx; i < l.length; i++ {
		l.data[i - 1] = l.data[i]
	}
	l.length--
	return
}

// 改
func (l *SqList)UpdateElem(idx int, e any) error{
	if idx < 1 || idx > l.length {
		return errors.New("修改位置不合法")
	}
	l.data[idx - 1] = e
	return nil
}

// 按位查找
func (l *SqList)GetElem(index int)(t any, err error){
	if index < 1 || index > l.length {
		err = errors.New("按位查找:越界")
	}
	t = l.data[index - 1]
	return
}

// 查：根据值查询位置
func (l *SqList)LocateElem(e any) (int, error){
	for i:= 0; i < l.length; i++ {
		if l.data[i] == e {
			return i, nil
		}
	}
	return 0, errors.New("按值查找:未找到")
}



