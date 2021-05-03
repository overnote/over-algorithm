package array

import (
	"testing"
)

func TestDynamicArray(t *testing.T) {

	// 初始化一个动态数组
	arr := NewDynamicArray()
	arr.Display()

	// 测试插入
	arr.Insert(10, 0)
	arr.Insert(11, 1)
	arr.Insert(12, 2)
	arr.Insert(13, 3)
	arr.Insert(14, 4)
	arr.Display()

	// 测试扩容
	arr.Insert(15, 5)
	arr.Insert(16, 6)
	arr.Display()
	t.Log("容量：", arr.Capacity())

	// 测试删除
	arr.Delete(0)
	arr.Display()
	arr.Delete(2)
	arr.Display()

	// 测试缩容
	t.Log("容量：", arr.Capacity())
	arr.Delete(4)
	arr.Delete(0)
	arr.Delete(0)
	arr.Delete(0)
	arr.Display()
	t.Log("容量：", arr.Capacity())

	// 测试更新
	arr.Update(100,0)

	// 测试查找
	t.Log("位置：",  arr.Locate(100))

	// 测试取值
	t.Log("取值：", arr.Get(0))

}
