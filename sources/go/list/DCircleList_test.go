package list

import "testing"

func TestDCircleList_PushBack(t *testing.T) {
	l := NewDCircleList()
	l.PushBack("11")
	l.PushBack("12")
	l.PushBack("13")
	t.Log(l.head.next)
	t.Log(l.head.next.next)
	t.Log(l.head.next.next.next)
	t.Log(l.head.next.next.next.next)
}

func TestDCircleList_PushFront(t *testing.T) {
	l := NewDCircleList()
	l.PushFront("11")
	l.PushFront("12")
	l.PushFront("13")
	t.Log(l.head.next)
	t.Log(l.head.next.next)
	t.Log(l.head.next.next.next)
	t.Log(l.head.next.next.next.next)
}

func TestDCircleList_Insert(t *testing.T) {
	l := NewDCircleList()
	l.Insert(1,"11")
	l.Insert(1,"12")
	l.Insert(2,"13")
	l.Insert(4,"14")
	t.Log(l.head.next)
	t.Log(l.head.next.next)
	t.Log(l.head.next.next.next)
	t.Log(l.head.next.next.next.next)
}