package list

import "testing"

func TestLinkedList_PushBack(t *testing.T) {
	l := NewLinkedList()
	l.PushBack(11)
	l.PushBack(12)
	t.Log(l.head)
	t.Log(l.head.next)
	t.Log(l.head.next.next)
}

func TestLinkedList_PushFront(t *testing.T) {
	l := NewLinkedList()
	l.PushFront(11)
	l.PushFront(12)
	t.Log(l.head)
	t.Log(l.head.next)
	t.Log(l.head.next.next)
}

func TestLinkedList_Insert(t *testing.T) {
	l := NewLinkedList()
	l.Insert(1,11)
	l.Insert(2,12)
	l.Insert(2,13)
	l.Insert(4,14)
	t.Log(l.head.next)
	t.Log(l.head.next.next)
	t.Log(l.head.next.next.next)
	t.Log(l.head.next.next.next.next)
}

func TestLinkedList_Delete(t *testing.T) {
	l := NewLinkedList()
	l.PushBack(11)
	l.PushBack(12)
	l.PushBack(13)
	l.PushBack(14)
	l.Delete(14)
	t.Log(l.head.next.next.next)
	l.Delete(12)
	t.Log(l.head.next.next)
	l.Delete(11)
	t.Log(l.head.next)
}