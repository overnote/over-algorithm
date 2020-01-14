package queue

import "testing"

func TestLinkedQueue_EnQueue(t *testing.T) {
	lq := NewLinkQueue()
	lq.EnQueue("aa")
	lq.EnQueue("bb")
	lq.EnQueue("cc")
	t.Log("个数：", lq.length)
	t.Log("队首：", lq.front)
	t.Log("下一个元素：", lq.front.next)
	t.Log("队尾：", lq.rear)
}

func TestLinkedQueue_DeQueue(t *testing.T) {

	lq := NewLinkQueue()
	lq.EnQueue("aa")
	lq.EnQueue("bb")
	lq.EnQueue("cc")
	t.Log("个数：", lq.length)
	t.Log("队首：", lq.front)
	t.Log("下一个元素：", lq.front.next)
	t.Log("队尾：", lq.rear)
	lq.DeQueue()
	t.Log("个数：", lq.length)
	t.Log("队首：", lq.front)
	t.Log("队尾：", lq.rear)
}
