package tree

import (
	"fmt"
	"structure/binarytree/bstree"
)

func TestBSTree() {

	fmt.Println("*********** 测试开始 ***********")

	t := bstree.NewBSTree()
	t.InsertByGen(11)
	t.InsertByGen(7)
	t.InsertByGen(15)
	t.InsertByGen(5)
	t.InsertByGen(3)
	t.InsertByGen(9)
	t.InsertByGen(8)
	t.InsertByGen(10)
	t.InsertByGen(13)
	t.InsertByGen(12)
	t.InsertByGen(14)
	t.InsertByGen(20)
	t.InsertByGen(18)
	t.InsertByGen(25)

	fmt.Println("最大值：", t.MaxData())
	fmt.Println("最小值：", t.MinData())
	fmt.Println("查找数值是否存在：", t.SearchData(25))

	fmt.Println("*********** 测试结束 ***********")

}
