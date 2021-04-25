/*
 *  动态数组
 */

package array

import "fmt"

const MaxSize = 5 // 默认容量，取小值便于测试

type DynamicArray struct {
	data     []interface{} // 动态数组内的真实数组
	capacity int           // 动态数组容量
	length   int           // 动态数组元素个数
}

// 创建动态数组对象
func NewDynamicArray() *DynamicArray {
	return &DynamicArray{
		data:     make([]interface{}, MaxSize),
		capacity: MaxSize,
		length:   0,
	}
}

// 扩容方法：超过数组容量，按照当前容量的 2 倍扩容
func (arr *DynamicArray) expandCap() {
	if arr.length == arr.capacity {
		// 拷贝数据
		newCap := arr.capacity * 2
		newData := make([]interface{}, newCap)
		for i := 0; i < arr.length; i++ {
			newData[i] = arr.data[i]
		}
		arr.data = newData

		// 变更容量
		arr.capacity = newCap
	}
}

// 缩容方法：数组元素为当前容量的 1/4，缩容为当前容量的一半
func (arr *DynamicArray) reduceCap() {
	if arr.length <= arr.capacity/4 {

		newCap := arr.capacity / 2
		newData := make([]interface{}, newCap)
		for i := 0; i < arr.length; i++ {
			newData[i] = arr.data[i]
		}
		arr.data = newData

		// 变更新容量
		arr.capacity = newCap
	}
}

// 增：增加最后一位
func (arr *DynamicArray) PushElem(e interface{}) {
	if arr.data == nil {
		fmt.Println("数组不存在")
		return
	}

	arr.expandCap() // 扩容
	arr.length++
	arr.data[arr.length-1] = e
}

// 删:移除最后一位
func (arr *DynamicArray) PopElem() {
	if arr.length == 0 {
		fmt.Println("数组元素个数为0")
		return
	}

	arr.reduceCap() // 缩容
	arr.length--
}

// 查：根据索引获取数据
func (arr *DynamicArray) IndexElem(index int) interface{} {
	if index < 0 || index > arr.length-1 {
		fmt.Println("数组索引越界")
		return nil
	}
	return arr.data[index]
}

// 获取容量
func (arr *DynamicArray) Capacity() int {
	return arr.capacity
}

// 获取长度
func (arr *DynamicArray) Length() int {
	return arr.length
}

// 销毁：自动内存管理，无需该方法

// 显示动态数组数据
func (arr *DynamicArray) Display() {
	fmt.Println(arr.data[0:arr.length])
}
