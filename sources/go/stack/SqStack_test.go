package stack

import "testing"

func TestSqStack_Push(t *testing.T) {
	stack := NewSqStack()
	stack.Push("aaa")
	stack.Push("bbb")
	t.Log(stack)
}

func TestSqStack_Pop(t *testing.T) {
	stack := NewSqStack()
	stack.Push("aaa")
	stack.Push("bbb")
	stack.Pop()
	t.Log(stack)
	stack.Pop()
	t.Log(stack)
	stack.Pop()
	t.Log(stack)
}

func TestSqStack_Top(t *testing.T) {
	stack := NewSqStack()
	stack.Push("aaa")
	stack.Push("bbb")
	t.Log(stack.Top())
}