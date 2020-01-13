package stack

import "testing"

func TestLinkedStack_Push(t *testing.T) {
	stack := NewLinkedStack()
	stack.Push("aa")
	stack.Push("bb")
	stack.Push("cc")
	t.Log("第1个节点：", stack.top)
	t.Log("第2个节点：", stack.top.next)
	t.Log("第3个节点：", stack.top.next)
	t.Log("第4个节点：", stack.top.next.next)
}

func TestLinkedStack_Pop(t *testing.T) {
	stack := NewLinkedStack()
	stack.Push("aa")
	stack.Push("bb")
	stack.Push("cc")
	stack.Pop()
	t.Log("第1个节点：", stack.top)
	t.Log("第2个节点：", stack.top.next)
	t.Log("第3个节点：", stack.top.next.next)
}

func TestLinkedStack_Top(t *testing.T) {
	stack := NewLinkedStack()
	stack.Push("aa")
	stack.Push("bb")
	stack.Push("cc")
	t.Log( stack.top)
}