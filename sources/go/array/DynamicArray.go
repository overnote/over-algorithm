/*
 *  可变数组
 */

package array

import (
	"fmt"
)

const MaxSize = 5 // 默认容量，取小值便于测试

type DynamicArray struct {
	data     []interface{} // 可变数组内的真实数组
	capacity int           // 可变数组容量
	length   int           // 可变数组元素个数
}

// 创建可变数组结构体
func NewDynamicArray() *DynamicArray {
	return &DynamicArray{
		data:     make([]interface{}, MaxSize),
		capacity: MaxSize,
		length:   0,
	}
}

// 增：根据索引插入数据
func (arr *DynamicArray) Insert(e interface{}, index int){
	if index < 0 || index > arr.length{
		fmt.Println("插入位置不合法")
		return
	}
	
	// 每次插入前执行扩容机制
	arr.expandCap()
	// 执行插入
	for i := arr.length; i >= index; i-- {
		if i == index {
			arr.data[index] = e
		} else {
			arr.data[arr.length] = arr.data[arr.length - 1]
		}
	}
	// 更新长度
	arr.length++
}

// 删：按索引位置删除元素
func (arr *DynamicArray) Delete(index int){
	if index < 0 || index > arr.length - 1 {
		fmt.Println("删除位置不合法")
		return
	}

	// 每次删除前执行缩容
	arr.reduceCap()

	// 执行删除
	for i := index ; i < arr.length; i++ {
		arr.data[i] = arr.data[ i + 1]
	}
	arr.length--
}

// 改：根据索引修改元素
func (arr *DynamicArray) Update(e interface{}, index int){
	if index < 0 || index > arr.length - 1 {
		fmt.Println("数组索引越界")
		return
	}
	arr.data[index] = e
}

// 查： 根据值查询索引
func (arr *DynamicArray) Locate(e interface{}) int{
	for i := 0; i < arr.length; i++{
		if arr.data[i] == e {
			return i
		}
	}
	return -1
}

// 取：根据索引进行数据存取
func (arr *DynamicArray) Get(index int) interface{} {
	if index < 0 || index > arr.length - 1 {
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

// 清空
func (arr *DynamicArray) Clear(){
	arr.length = 0
}

// 销毁：自动内存管理，无需该方法

// 扩容方法：超过数组容量，按照当前容量的 2 倍扩容
func (arr *DynamicArray) expandCap() {
	if arr.length < arr.capacity {
		return
	}
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

// 缩容方法：数组元素为当前容量的 1/4，缩容为当前容量的一半
func (arr *DynamicArray) reduceCap() {
	if arr.length > arr.capacity / 4 {
		return
	}

	newCap := arr.capacity / 2
	newData := make([]interface{}, newCap)
	for i := 0; i < arr.length; i++ {
		newData[i] = arr.data[i]
	}
	arr.data = newData

	// 变更新容量
	arr.capacity = newCap
}

// 显示可变数组数据
func (arr *DynamicArray) Display() {
	fmt.Println(arr.data[0:arr.length])
}
