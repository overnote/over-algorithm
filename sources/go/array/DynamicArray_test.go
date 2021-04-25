package array

import (
	"fmt"
	"testing"
)

func TestDynamicArray(t *testing.T) {

	// 初始化一个动态数组
	arr := NewDynamicArray()
	arr.Display()

	arr.PushElem(1)
	arr.PushElem(2)
	arr.PushElem(3)
	arr.PushElem(4)
	arr.PushElem(5)
	fmt.Println("arr[4]=", arr.IndexElem(4))
	arr.Display()

	arr.PushElem(6)
	arr.PushElem(7)
	arr.Display()

	arr.PopElem()
	arr.PopElem()
	arr.PopElem()
	arr.PopElem()
	arr.PopElem()
	arr.Display()
}
