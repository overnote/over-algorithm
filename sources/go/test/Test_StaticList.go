package test

import (
	"fmt"
	"structure/list/staticlist"
)

func TestStaticList() {

	fmt.Println("*********** 测试开始 ***********")
	l := staticlist.NewStaticList(10)

	fmt.Println("----------- 增加元素 -----------")
	l.Insert(1, 11)
	l.Insert(2, 12)
	l.Insert(3, 13)
	l.Insert(4, 14)
	l.Insert(5, 15)
	l.Display()

	fmt.Println("----------- 删除元素 -----------")
	l.Delete(1)
	l.Display()

	fmt.Println("*********** 测试结束 ***********")

}
