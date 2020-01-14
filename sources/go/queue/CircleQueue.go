/*
	循环队列：一般使用数组实现，因为可以合理利用数组的空间，可以顺便解决顺序存储的不足。当然链表也可以实现循环队列，而且更加方便
	扩容：这里忽略扩容情况。如果要扩容，其步骤是：
		- head应该设置为0
		- 重新计算所有元素的下标：(i + head) % cap   i为原数组从索引0开始循环的遍历i值
*/
package queue

import (
	"fmt"
)

const CircleQueueCap = 5

type CircleQueue struct {
	items     	[]interface{}	// 队列中的数据存放数组（切片）
	cap 		int				// 队列的容量
	length		int 			// 队列元素个数
	front		int				// 队首下标
	rear		int				// 队尾下标
}

func NewCircleQueue() *CircleQueue{
	return &CircleQueue{
		items:  make([]interface{}, CircleQueueCap),
		cap:    CircleQueueCap,
		front:	0,
		rear:	0,
	}
}

// 入队
func (q *CircleQueue)Enqueue(e interface{}) bool{

	if q.Length() == q.Cap() {
		fmt.Println("队列容量已满")
		return false
	}

	q.rear = q.realIndex(q.Length())	// 计算入队后tail的真实位置
	q.items[q.rear] = e
	q.length++
	return true
}

// 出队
func (q *CircleQueue)DeQueue() interface{}{

	if q.Length() == 0 {
		fmt.Println("队列为空")
		return nil
	}

	// 获取当前队首元素
	e := q.items[q.front]				// 先获取旧元素，因为要return出去
	q.front = q.realIndex(1)		// 计算出队后head的真实位置
	q.length--
	return e
}

// 获取在循环队列数组中的真实索引
func (q *CircleQueue)realIndex(index int) int{
	return (q.front + index) % q.Cap()
}

// 获取元素个数
func (q *CircleQueue)Length() int {
	return q.length
}

// 获取队列容量
func (q *CircleQueue)Cap() int {
	return q.cap
}