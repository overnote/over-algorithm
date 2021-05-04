package list

import (
	"testing"
)

func TestLinkList(t *testing.T) {
	l := NewLinkList()
	l.Insert(11, 1)
	l.Insert(12, 2)
	l.Insert(13, 3)
	l.Insert(14, 4)
	l.Insert(15, 5)
	l.Display()

	t.Log(l.Delete(5))
	l.Display()
	t.Log(l.Delete(3))
	l.Display()

	l.Update(3, 33)
	l.Display()

	l.Clear()
	l.Display()
}