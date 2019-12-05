/*
	循环队列：也称为双端队列
 */

package cirqueue

import (
	"errors"
	"fmt"
)

type CirQueue struct {
	data     	[]interface{}	// 队列中的数据存放数组（切片）
	front		int				// 队首索引
	rear		int				// 队尾索引
	size 		int				// 队列的容量
}

func NewCirQueue(size int) *CirQueue{
	return &CirQueue{
		data:   nil,
		front:	0,
		rear:	0,
		size:   size,
	}
}

// 获取元素个数
func (q *CirQueue)Length() int {
	return (q.rear - q.front + q.size) % q.size
}

// 判断队列是否已满：标准判定规则是
func (q *CirQueue)IsFull() bool{
	if  (q.rear + 1) % q.size == q.front {
		return true
	} else {
		return false
	}
}

// 显示
func (q *CirQueue)Display() {

	if q.Length() == 0 {
		fmt.Println("数据结构为空")
		return
	}
	fmt.Println("队首索引：", q.front)
	fmt.Println("队尾索引：", q.rear)
	fmt.Println("队列元素：", q.data)
}

// 入队
func (q *CirQueue)EnQueue(e interface{}) error{

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
	q.data[q.rear] = e
	// 队尾索引后移一位
	q.rear = (q.rear + 1) % q.size

	return nil
}

// 出队
func (q *CirQueue)DeQueue() (interface{}, error){

	if q.front == q.rear {
		fmt.Println("队列为空")
		return nil, errors.New("队列为空")
	}

	// 获取当前队首元素
	e := q.data[q.front]

	// front索引后移
	q.front = (q.front + 1) %  q.size

	return e, nil
}