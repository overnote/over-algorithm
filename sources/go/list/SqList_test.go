package list

import "testing"

func TestSqList_PushBack(t *testing.T) {
	l := NewSqList()
	l.PushBack(11)
	l.PushBack(12)
	l.PushBack(13)
	l.PushBack(14)
	l.PushBack(15)
	l.PushBack(16)
	t.Log(l.items)
}

func TestSqList_Insert(t *testing.T) {
	l := NewSqList()
	l.Insert(1,11)
	l.Insert(5,12)
	l.Insert(2,13)
	l.Insert(3,14)
	l.Insert(2,15)
	l.Insert(2,16)
	l.Insert(2,17)
	t.Log(l.items)
}