package SeqQueue

/*
循环队列
*/

import (
	"fmt"
	"errors"
)

// 循环队列对象
type SeqQueue struct {
	data []interface{}
	size int				// 容量
	front int				// 头指针
	rear int 				// 尾指针
}

// 创建循环队列
func New(size int) *SeqQueue {
	return &SeqQueue{
		make([]interface{}, size),
		size,
		0,
		0,
	}
}

// 判断队列是否为空
func (sq *SeqQueue) IsEmpty() bool{
	if sq.front == sq.rear || sq.data[sq.rear - 1] == nil {
		return true
	}
	return false
}

// 获取循环队列长度
func (sq *SeqQueue) Length() int {
	if sq.IsEmpty() {
		return 0
	}
	return (sq.rear - sq.front + sq.size) % sq.size
}

// 清空循环队列
func (sq *SeqQueue) Clear() {
	sq.front = 0
	sq.rear = 0
	sq.data = nil
}

// 获取头结点
func (sq *SeqQueue) Head() interface{} {
	if sq.IsEmpty() {
		return nil
	}
	return sq.data[sq.front]
}

// 新增
func (sq *SeqQueue) Push(data interface{}) error{
	if (sq.rear + 1) % sq.size == sq.front {
		return errors.New("queue is full")
	}
	sq.data[sq.rear] = data
	sq.rear = (sq.rear + 1) % sq.size
	return nil
}

// 删除
func (sq *SeqQueue) Delete() interface{} {
	if sq.IsEmpty() {
		fmt.Println("queue is empty")
		return nil
	}
	e := sq.data[sq.front]
	sq.data[sq.front] = nil
	sq.front = (sq.front + 1) % sq.size
	return e
}