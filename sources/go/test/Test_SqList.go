package test

import (
	"fmt"
	"structure/list/sqlist"
)

func TestSqList() {

	fmt.Printf("********** 顺序表测试开始 **********\n")

	l := sqlist.NewSqList()
	l.Display()

	fmt.Printf("增加元素-------------\n")
	l.Append(17)
	l.Insert(2, 18)
	l.Insert(3, 19)
	fmt.Printf("长度 %d\n", l.Length())
	l.Display()

	fmt.Printf("修改元素-------------\n")
	l.Update(2, 22)
	l.Display()

	pe, _ := l.PriorElem(22)
	ne, _ := l.NextElem(22)
	fmt.Printf("查找第二个元素前驱：%d\n", pe)
	fmt.Printf("查找第二个元素后继：%d\n", ne)

	fmt.Printf("删除元素-------------\n")
	l.Delete(2)
	l.Display()

	fmt.Printf("********** 顺序表测试结束 **********\n")

}
