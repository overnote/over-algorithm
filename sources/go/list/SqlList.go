/**
 * 顺序表：SqList
 */

package list

import (
	"errors"
	"fmt"
)

// 顺序表结构体
type SqList struct {
	data    		[]interface{} 	// 使用切片结构，可以摆脱数组容量限制
	size 			int         	// 该顺序表元素个数
}
 
// 构造实例
func NewSqList() *SqList {
	return &SqList{
		data:		make([]interface{}, 5),// 给一个初始默认长度（非容量）
		size: 	0,
	}
}
 
// 增：按照一个位置插入数据
func (l *SqList)Insert(e interface{}, index int) {
 
	if index < 1 || index > l.size + 1 {
		fmt.Println("插入位置不合法")
		return
	}
 
	// 将切片在插入位置一份为2，再合并，即可实现插入效果
	l.data= append(l.data[: index-1], append([]interface{}{e}, l.data[index-1:]...)...)
	l.size++
}
 
// 删：删除位置上的元素，并获取到删除元素的值
func (l *SqList)Delete(index int) interface{}{
	if index < 1 || index > l.size {
		fmt.Println("删除位置不合法")
		return nil
	}
 
	if l.size == 0 {
		fmt.Println("空表无元素可移除")
		return nil
	}

	e := l.data[index - 1]
 
	// 也可以使用append方式
	for i := index; i < l.size; i++ {
		l.data[i - 1] = l.data[i]
	}
	l.size--
	return e
}

// 改
func (l *SqList)Update(index int, e interface{}){
	if index < 1 || index > l.size {
		fmt.Println("修改位置不合法");
		return
	}
	l.data[index - 1] = e
}
 
// 查：根据位置查询值
func (l *SqList)Search(index int)(error, interface{}){
	if index < 1 || index > l.size {
		fmt.Println("参数位置不合法")
		return errors.New("参数位置不合法"), 0
	}
	return nil, l.data[index - 1]
}
 
// 查：根据值查询位置
func (l *SqList)Locate(e interface{}) (error, int){
	for i:= 0; i < l.size; i++ {
		if l.data[i] == e {
			return nil, i
		}
	}
	fmt.Println("未找到该元素")
	return errors.New("未找到该元素"), -1
}
 
// 获取顺序表长度
func (l *SqList)Length() int{
	return l.size
}
 
// 清空
func (l *SqList)Clear(){
	l.size = 0
}
 
// 销毁
// Go自动内存管理无需销毁
 
// 显示顺序表
func (l *SqList) Display() {
	fmt.Println(l.data[0:l.size])
}
 