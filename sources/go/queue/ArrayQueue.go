/*
	循环队列：数组实现
*/
package queue

import (
	"errors"
	"fmt"
)

type ArrayQueue struct {
	data     	[]interface{}	// 队列中的数据存放数组（切片）
	size 		int				// 队列的容量
	head		int				// 队首下标
	tail		int				// 队尾下标
}

func NewArrayQueue(size int) *ArrayQueue{
	return &ArrayQueue{
		data:   make([]interface{}, size),
		size:   size,
		head:	0,
		tail:	0,
	}
}

// 入队
func (q *ArrayQueue) Enqueue(el interface{}) bool{

	if q.tail == q.size - 1 {
		fmt.Println("队列容量已满")
		return false
	}

	q.data[q.tail] = el
	q.tail++

	return true
}

func (q *ArrayQueue)EnQueue(e interface{}) error{

	if q.IsFull() {
		fmt.Println("队列已满")
		return errors.New("队列已满")
	}

	// 由于go语言没有泛型，这里在第一次入队时，将所有元素都初始化一次
	if len(q.data) == 0 {
		q.data = make([]interface{}, q.size)
		for i := 0; i < q.size; i++ {
			q.data = append(q.data, e)
		}
		q.data = q.data[q.size:]
	}

	// 将元素e赋值给队尾
	q.data[q.head] = e
	// 队尾索引后移一位
	q.tail = (q.tail + 1) % q.size

	return nil
}

// 出队
func (q *ArrayQueue)DeQueue() (interface{}, error){

	if q.head == q.tail {
		fmt.Println("队列为空")
		return nil, errors.New("队列为空")
	}

	// 获取当前队首元素
	e := q.data[q.head]

	// front索引后移
	q.head = (q.head + 1) %  q.size

	return e, nil
}

// 显示队列数据
func (q *ArrayQueue)Display() {

	if q.Length() == 0 {
		fmt.Println("队列为空")
		return
	}
	fmt.Println("队首索引：", q.head)
	fmt.Println("队尾索引：", q.tail)
	fmt.Println("队列元素：", q.data)
}

// 获取元素个数
func (q *ArrayQueue)Length() int {
	return (q.tail - q.head + q.size) % q.size
}

// 获取队列容量
func (q *ArrayQueue)Size() int {
	return q.size
}