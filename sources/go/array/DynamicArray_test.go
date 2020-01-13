package array

import "testing"

func TestDynamicArray_Push(t *testing.T) {
	arr := NewDynamicArray()
	arr.Push("0")
	arr.Push("1")
	arr.Push("2")
	arr.Push("3")
	arr.Push("4")
	t.Log("容量：", arr.Size())
	t.Log("元素：", arr.Length())
	arr.Push("5")
	t.Log("容量：", arr.Size())
	t.Log("元素：", arr.Length())
	t.Log(arr)
}

func TestDynamicArray_Pop(t *testing.T) {
	arr := NewDynamicArray()
	arr.Push("0")
	arr.Push("1")
	arr.Push("2")
	arr.Push("3")
	arr.Push("4")
	arr.Push("5")
	t.Log("容量：", arr.Size())
	t.Log("元素：", arr.Length())
	arr.Pop()
	t.Log("容量：", arr.Size())
	t.Log("元素：", arr.Length())
	arr.Pop()
	arr.Pop()
	arr.Pop()
	t.Log("容量：", arr.Size())
	t.Log("元素：", arr.Length())
	t.Log(arr)
}

func TestDynamicArray_Index(t *testing.T) {
	arr := NewDynamicArray()
	arr.Push("0")
	arr.Push("1")
	arr.Push("2")
	arr.Push("3")
	arr.Push("4")
	arr.Push("5")
	t.Log(arr.Index(5))
}