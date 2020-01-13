package stack

import "testing"

func TestShareStack_Push(t *testing.T) {
	stack := NewShareStack()
	stack.Push("a1", 1)
	stack.Push("a2", 1)
	stack.Push("b1", 2)
	stack.Push("b2", 2)
	t.Log(stack)
}

func TestShareStack_Pop(t *testing.T) {
	stack := NewShareStack()
	stack.Push("a1", 1)
	stack.Push("a2", 1)
	stack.Push("b1", 2)
	stack.Push("b2", 2)
	stack.Pop(1)
	t.Log(stack)
	stack.Pop(1)
	t.Log(stack)
	stack.Pop(2)
	t.Log(stack)
	stack.Pop(1)
	t.Log(stack)
}

func TestShareStack_Top(t *testing.T) {
	stack := NewShareStack()
	stack.Push("a1", 1)
	stack.Push("a2", 1)
	stack.Push("b1", 2)
	stack.Push("b2", 2)
	t.Log(stack.Top(1))
}