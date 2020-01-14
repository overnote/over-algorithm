package queue

import "testing"

func TestCircleQueue_Enqueue(t *testing.T) {
	queue := NewCircleQueue()
	queue.Enqueue("aa")
	queue.Enqueue("bb")
	queue.Enqueue("cc")
	queue.Enqueue("dd")
	queue.Enqueue("ee")
	queue.Enqueue("ff")
	t.Log("个数：", queue.length)
	t.Log("队首：", queue.front)
	t.Log("队尾:", queue.rear)
	t.Log("元素:", queue.items)
}

func TestCircleQueue_DeQueue(t *testing.T) {
	queue := NewCircleQueue()
	queue.Enqueue("aa")
	queue.Enqueue("bb")
	queue.DeQueue()
	t.Log("个数：", queue.length)
	t.Log("队首：", queue.front)
	t.Log("队尾:", queue.rear)
	t.Log("元素:", queue.items)
	queue.Enqueue("cc")
	queue.Enqueue("dd")
	queue.Enqueue("ee")
	queue.Enqueue("ff")
	queue.DeQueue()
	t.Log("个数：", queue.length)
	t.Log("队首：", queue.front)
	t.Log("队尾:", queue.rear)
	t.Log("元素:", queue.items)
}