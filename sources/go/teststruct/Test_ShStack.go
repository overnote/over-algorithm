package teststruct

import (
	"fmt"
	"structure/stack/shstack"
)

func TestShStack() {

	fmt.Println("*********** 测试开始 ***********")
	s := shstack.NewShStack(5)

	fmt.Println("----------- 元素进栈 -----------")
	s.Push(0, 1)
	s.Push(1, 1)
	s.Push(2, 1)
	s.Push(3, 2)
	s.Push(4, 2)
	s.Push(5, 2)
	s.Display()

	fmt.Println("----------- 元素出栈 -----------")
	s.Pop(2)
	s.Pop(1)
	s.Display()

	fmt.Println("*********** 测试结束 ***********")

}
