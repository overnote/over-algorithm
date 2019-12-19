package tree

import (
	"fmt"
	"structure/tree/binarytree"
)

func TestBinaryTree() {

	fmt.Println("*********** 测试开始 ***********")

	fmt.Println("请依次输入 A B D H # # I # # E J # # # C F # # G # # 构造一个完全二叉树")
	t := binarytree.NewBiTree()

	fmt.Println("----------- 树长度 -----------")
	fmt.Println(t.Length())

	fmt.Println("----------- 树遍历 -----------")
	t.PreOrderTraverse()

	fmt.Println("*********** 测试结束 ***********")

}
