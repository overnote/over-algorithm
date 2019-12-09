package testStruct

import (
	"fmt"
	"structure/list/sqlist"
)

func TestSqList() {

	fmt.Println("*********** 测试开始 ***********")
	l := sqlist.NewSqList(3, 2)

	fmt.Println("----------- 增加元素 -----------")
	l.Append(6)
	l.Insert(1, 7)
	l.Insert(3, 4)
	l.Insert(4, 9)
	l.Display()

	fmt.Println("----------- 删除元素 -----------")
	l.Delete(4)
	l.Display()

	fmt.Println("----------- 修改元素 -----------")
	l.Update(3, 22)
	l.Display()

	fmt.Println("----------- 查询元素 -----------")
	e, _ := l.GetElem(3)
	i, _ := l.Locate(22)
	fmt.Println("查询索引值：", e)
	fmt.Println("查询值索引：", i)
	pe, _ := l.PrevElem(6)
	ne, _ := l.NextElem(6)
	fmt.Printf("查找前驱：%v\n", pe)
	fmt.Printf("查找后继：%v\n", ne)

	fmt.Println("*********** 测试结束 ***********")

}
