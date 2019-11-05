package test

import (
	"fmt"
	doublelist "structure/list/doublelist"
)

func TestDoubleList() {

	fmt.Printf("********** 测试开始 **********\n")

	l := doublelist.NewDoubleList()

	fmt.Println("追加元素-------------")
	l.Append(17)
	l.Append(20)
	l.Display()

	fmt.Println("插入元素-------------")
	l.Insert(2, 18)
	l.Insert(3, 19)
	l.Display()

	fmt.Println("删除元素-------------")
	l.Delete(1)
	l.Display()

	fmt.Println("修改元素-------------")
	l.Update(3, 22)
	l.Display()

	fmt.Println("按照位序查找元素-------------")
	ge, _ := l.GetElem(2)
	fmt.Printf("查找元素值：%d\n", ge)

	fmt.Println("按照值查找位序-------------")
	le, _ := l.LocateElem(18)
	fmt.Printf("查找元素位序：%d\n", le)

	fmt.Println("查找元素前驱-------------")
	pe, _ := l.PriorElem(24)
	fmt.Printf("查找元素前驱：%d\n", pe)

	fmt.Println("查找元素后继-------------")
	ne, _ := l.NextElem(18)
	fmt.Printf("查找元素前驱：%d\n", ne)

	fmt.Println("删除元素-------------")
	l.Delete(2)
	l.Display()

	fmt.Printf("********** 测试结束 **********\n")

}
