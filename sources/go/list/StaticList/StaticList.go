/**
 * 静态链表
 */

package list

import (
	"errors"
	"fmt"
	"os"
)

// 节点结构体
type node struct {
	data interface{}
	cur  int 				// 游标：为0时表示无指向
}

// 静态链表
type StaticList struct {
	data     []node
	size int
	length  int
}

func New(size int) (*StaticList, error){

	if size < 3 {
		return nil, errors.New("size overflow")
	}

	s := make([]node, size)

	for i := 0; i <= size-1; i++ {
		s[i].cur = i + 1
		if i == size - 1 {
			s[size-1].cur = 0 
		}
	}

	return &StaticList{s, size, 0},nil 
}

// 判断是否为空
func (sl *StaticList) IsEmpty() bool {
	if sl.length == 0 {
		return true
	}
	return false
}

// 分配节点
func (sl *StaticList) malloc() int {
	i := sl.data[0].cur
	if i == 0 {
		os.Exit(0)
	}
	sl.data[0].cur = sl.data[i].cur
	return i
}

// 回收节点
func (sl *StaticList) free(index int) {
	sl.data[index].cur = sl.data[0].cur
	sl.data[0].cur = index
}

// 回收链表到备用链表
func (sl *StaticList) DestroyList() {

	j := sl.data[sl.size-1].cur

	if j == 0 {
		return
	}

	sl.data[sl.size-1].cur = 0

	i := sl.data[0].cur
	sl.data[0].cur = j
	if j > 0 {
		j = sl.data[j].cur
	}
	sl.data[j].cur = i
}

// 插入节点
func (sl *StaticList) Insert(data interface{}, index int) error{

	if index < 1 || index > sl.length {
		return errors.New("index overflow")
	}

	i := sl.data[sl.size-1].cur
	j := 1
	for i > 0 && j < index-1 {
		j++
		i = sl.data[i].cur
	}
	tmp := sl.data[i].cur
	cur := sl.malloc()
	sl.data[cur].data = data
	sl.data[cur].cur = tmp
	sl.data[i].cur = cur
	return nil
}

// 显示链表结构
func (sl *StaticList) Show() {
	for i, v := range sl.data {
		fmt.Printf("%5d:%5d,%5s", i, v.cur, v.data)
	}
}

// 获取数据元素位置
func (sl *StaticList) Location(data interface{}) int {
	location := 0
	i := sl.data[sl.size-1].cur
	for i > 0 {
		location++
		if sl.data[i].data == data {
			return location
		}
		i = sl.data[i].cur
	}
	return location
}

// 获取链表长度
func (sl *StaticList) Length() int {
	return sl.length
}

// 获取链表容量
func (sl *StaticList) Size() int {
	return sl.size
}