package list

import (
	"fmt"
	"testing"
)

func TestDblList(t *testing.T) {
	l := NewDblList()
	l.Insert(11, 1)
	l.Insert(12, 2)
	l.Insert(13, 3)
	l.Insert(14, 4)
	l.Insert(15, 5)
	l.Insert(33, 3)
	l.Display()

	fmt.Println("删除元素：", l.Delete(5))
	l.Display()
	fmt.Println("删除元素：", l.Delete(3))
	l.Display()

	l.Update(1, 21)
	l.Display()
	fmt.Println("查询元素 21, 结果：", l.Search(21).data)
	fmt.Println("查询前驱 21, 结果：", l.PrevElem(21).data)
	fmt.Println("查询前驱 21, 结果：", l.NextElem(21).data)

	l.Update(4, 44)
	l.Display()
	fmt.Println("查询元素 44, 结果：", l.Search(44).data)
	fmt.Println("查询前驱 4, 结果：", l.PrevElem(44).data)
	fmt.Println("查询前驱 44, 结果：", l.NextElem(44).data)

	l.Clear()
	l.Display()
}