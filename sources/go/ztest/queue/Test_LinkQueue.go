package queue

import (
	"fmt"
	"structure/queue/linkqueue"
)

func TestLinkQueue() {

	fmt.Println("*********** 测试开始 ***********")
	q := linkqueue.NewLinkQueue()

	fmt.Println("----------- 元素入队 -----------")
	q.EnQueue(0)
	q.EnQueue(1)
	q.EnQueue(2)
	q.EnQueue(3)
	q.EnQueue(4)
	q.Display()

	fmt.Println("----------- 元素出队 -----------")
	q.DeQueue()
	q.DeQueue()
	q.Display()

	fmt.Println("*********** 测试结束 ***********")

}
