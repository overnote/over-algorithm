package list

import (
	"fmt"
	"testing"
)

func TestSqList(t *testing.T) {

	l := NewSqList()
	l.Display()

    l.Insert(11, 1)
    l.Insert(12, 2)  
    l.Insert(13, 3)
    l.Insert(14, 4)
    l.Insert(15, 5)
    l.Insert(17, 7)
    l.Insert(15, 5)
	l.Display()

	fmt.Println("删除元素:", l.Delete(3))
	l.Display()

    l.Update(1, 100);
	l.Display()
	
	err1, e1 := l.Search(2)
	fmt.Println("查询元素:", err1, e1)

	err2, e2 := l.Locate(100)
	fmt.Println("查询元素:",err2, e2)
}
