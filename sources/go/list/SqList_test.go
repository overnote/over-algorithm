package list

import "testing"


func TestSqList(t *testing.T) {

	// 初始化一个动态数组
	l := NewSqList()
	l.Display()

	l.Insert(11, 1)
	l.Insert(12, 2)
	l.Insert(17, 7)
	l.Display()

	l.Insert(13, 3)
	l.Insert(14, 4)
	l.Insert(15, 5)
	l.Display()

	l.Insert(16, 6)
	l.Insert(5, 5)
	l.Display()

	l.Delete(3)
	l.Display()
}
