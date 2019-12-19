package tree

import (
	"fmt"
	"structure/tree/bstree"
)

func TestBSTree() {

	fmt.Println("*********** 测试开始 ***********")

	t := bstree.NewBSTree()
	t.Insert(11)
	t.Insert(7)
	t.Insert(15)
	t.Insert(5)
	t.Insert(3)
	t.Insert(9)
	t.Insert(8)
	t.Insert(10)
	t.Insert(13)
	t.Insert(12)
	t.Insert(14)
	t.Insert(20)
	t.Insert(18)
	t.Insert(25)

	fmt.Println("最大值：", t.MaxData())
	fmt.Println("最小值：", t.MinData())
	fmt.Println("查找数值是否存在：", t.SearchData(25))

	fmt.Println("*********** 测试结束 ***********")

}
