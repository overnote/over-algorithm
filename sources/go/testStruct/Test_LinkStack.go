package testStruct

import (
	"fmt"
	"structure/stack/linkstack"
)

func TestLinkStack() {

	fmt.Println("*********** 测试开始 ***********")
	s := linkstack.NewLinkStack()

	fmt.Println("----------- 元素进栈 -----------")
	s.Push(7)
	s.Push(3)
	s.Push(4)
	s.Display()

	fmt.Println("----------- 元素出栈 -----------")
	s.Pop()
	s.Display()

	fmt.Println("*********** 测试结束 ***********")

}
