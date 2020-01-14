/*
 *  动态数组
 */

package array

import "fmt"

const DynamicArrayCap = 5		// 默认容量

type DynamicArray struct {
	items 	[]interface{}
	cap 	int            	// 动态数组容量
	length	int				// 动态数组元素个数
}

// 创建动态数组对象
func NewDynamicArray() *DynamicArray{
	return &DynamicArray{
		items:  make([]interface{}, DynamicArrayCap),
		cap:   DynamicArrayCap,
		length: 0,
	}
}

// 获取数组容量
func (arr *DynamicArray)Size() int {
	return arr.cap
}

// 获取数组元素个数
func (arr *DynamicArray)Length() int {
	return arr.length
}

// Push元素：添加元素到末尾
func (arr *DynamicArray)Push(e interface{}) {
	arr.length++
	arr.expandCap()			// 扩容
	arr.items[arr.Length() - 1] = e
}

// Pop元素：删除最后一个元素
func (arr *DynamicArray)Pop() interface{}{
	if arr.Length() <= 0 {
		fmt.Println("当前数组为空")
		return nil
	}

	e := arr.items[arr.Length() - 1]
	arr.items[arr.Length() - 1] = nil
	arr.length--

	arr.reduceCap()		// 缩容
	return e
}

// 扩容方法：超过数组容量，按照当前容量的 2 倍扩容
func (arr *DynamicArray)expandCap() {
	if arr.length <= arr.cap {
		return						// 无需扩容
	}

	newSize := arr.Size() * 2
	newItems := make([]interface{}, newSize)
	for i := 0; i < arr.Length() - 1; i++ {
		newItems[i] = arr.items[i]
	}
	arr.items = newItems
	arr.cap = newSize
}

// 缩容方法：数组元素个数为当前容量 1/4 时，缩容为当前容量的一半
func (arr *DynamicArray)reduceCap() {
	if arr.length > arr.cap / 4 {
		return						// 无需缩容
	}

	newSize := arr.Size() / 2
	newItems := make([]interface{}, newSize)
	for i := 0; i < arr.Length() - 1; i++ {
		newItems[i] = arr.items[i]
	}
	arr.items = newItems
	arr.cap = newSize
}

// 根据索引查找元素
func (arr *DynamicArray)Index(index int) interface{}{
	if index < 0 || index > arr.Length() - 1 {
		fmt.Println("数组索引越界")
		return nil
	}
	return arr.items[index]
}