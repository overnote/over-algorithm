package testStruct

import (
	"fmt"
	"structure/stack/sqstack"
)

func TestSqStack() {

	fmt.Println("*********** 测试开始 ***********")
	s := sqstack.NewSqStack(5)

	fmt.Println("----------- 元素进栈 -----------")
	s.Push(7)
	s.Push(4)
	s.Display()

	fmt.Println("----------- 元素出栈 -----------")
	s.Pop()
	s.Display()

	fmt.Println("*********** 测试结束 ***********")

}
