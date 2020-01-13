package heap

import (
	"errors"
	"fmt"
)

const SIZE = 10
//const INCREMENT = 1.5

// 最大堆：简单实现，未考虑扩容，长度溢出错误
type BinaryHeap struct {
	data 		[]int
	size 		int
	length 		int
}

func New() *BinaryHeap{
	return &BinaryHeap{
		data:   make([]int, SIZE),
		size:   SIZE,
		length: 0,
	}
}

func (h *BinaryHeap) Add(e int){
	h.data[h.length - 1] = e
	h.sitUp(h.length - 1)
	h.length++
}

func (h *BinaryHeap) Remove() (root int, err error){

	if h.length == 0 {
		fmt.Println("数据结构为空，无法删除")
		err = errors.New("数据结构为空，无法删除")
		return
	}

	root = h.data[0]
	h.data[0] = h.data[h.length - 1]
	h.data[h.length - 1] = 0
	h.length--

	h.sitDown(0)
	return
}

// 传入数组批量建堆，可以对数组进行循环，每个元素执行堆的Add操作，但是该操作性能较低
/**
	做法一：自上而下的上滤
	for i := 1; i < h.length; i++ {
		siftUp(i)
	}
	做法二：自下而上的下滤
	for i := (size >> 1) - 1; i>= 0; i-- {
		siftDown(i)
	}

	效率上做法二更高
 */
func Heapify([]int) {

}

func (h *BinaryHeap) sitUp(index int) {

	e := h.data[index]						// 当前插入元素

	for index > 0 {

		pIndex := (index - 1) >> 1			// 获取父节点内容
		p := h.data[pIndex]

		// 可以在找到最终要交换的位置后再做交换
		if e <= p {
			break
		}
		h.data[index] = p		// 将父元素存储在index位置

		// 重新复制index
		index = pIndex
	}
	h.data[index] = e			// 优化性能代码
}

func (h *BinaryHeap) sitDown(index int) {

	half := h.length >> 1
	e := h.data[index]

	// 第一个叶子节点索引 == 非叶子节点的数量 循环的条件是 index位置必须是非叶子节点
	for index < half {

		// 节点有两种情况：只有左子节点，同时有左右子节点
		maxIndex := (index << 1) + 1
		maxChild := h.data[maxIndex]
		rightIndex := maxIndex + 1
		if rightIndex <= h.size && maxChild < h.data[rightIndex] {
			maxChild = h.data[rightIndex]
			maxIndex = rightIndex
		}

		if maxChild <= e {
			break
		}
		h.data[index] = maxChild
		index = maxIndex
	}
	h.data[index] = e
}