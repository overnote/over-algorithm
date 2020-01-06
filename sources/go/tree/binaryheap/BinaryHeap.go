package binaryheap

/**
	实现一个二叉最大堆
 */

type BinaryHeap struct {
	arr 	[]int
	size 	int
	length	int
}

func NewBinaryHeap(size int) *BinaryHeap{

	if size == 0 {
		size = 10			// 默认给出一个容量值为10
	}

	return &BinaryHeap{
		arr:  	make([]int, size),
		size: 	size,
		length: 0,
	}
}

// 添加元素：直接在数组最后添加。考虑情况：如果插入的值数据很大，则需要让其与父节点比较，交换，然后依次上浮
// 退出循环的条件：当前循环到的节点值小于等于父节点的值，或者没有父节点
//func (bh *BinaryHeap)Insert(num int) error{
//
//	if bh.length >= bh.size {
//		fmt.Println("容量达到最大")
//		return errors.New("容量达到最大")
//	}
//
//	// 在最后位置直接插入元素
//	bh.arr[bh.length] = num		// 已经有length个，新插入的索引为 length -1 + 1
//
//	// 插入后元素开始上浮
//	pIndex := (bh.length) >> 1	// 右移1位即为除以2，这是计算父节点的公式
//	for
//}