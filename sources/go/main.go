package main

import (
	"algorithm/array"
	SequenList "algorithm/list/SequenList"
	LinkedList "algorithm/list/LinkedList"
	RingList "algorithm/list/RingList"
	DoublyList "algorithm/list/DoublyList"
 	"fmt"
)

func testSparseArray() {

	var primitArray [11][11]int
	primitArray[1][2] = 1
	primitArray[2][3] = 2
	array.ShowOrigin(primitArray)

	s := array.NewSparseArray(primitArray)
	array.ShowSparse(s)

	arr := array.TransToArray(s)
	array.ShowOrigin(arr)
}

func testSequenList() {

	sl := SequenList.New()
	sl.Show()

	sl.Append(7)
	sl.Append(9)
	sl.Show()

	sl.Insert(1,6)
	sl.Show()

	sl.Pop()
	sl.Show()

}

func testLinkedList() {

	ll := LinkedList.New()
	ll.Show()

	ll.Append(9)
	ll.Append(5)
	ll.Show()

	ll.Insert(1,7)
	ll.Insert(3,4)
	ll.Show()

	ll.Delete(3)
	ll.Show()
	
	fmt.Println(ll.Node(1))
}

func testRingList() {

	rl := RingList.New()
	rl.Show()

	rl.Append(9)
	rl.Append(5)
	rl.Show()

	rl.Insert(1,7)
	rl.Insert(3,4)
	rl.Show()

	rl.Delete(3)
	rl.Show()
	
}

func testDoublyList() {

	dl := DoublyList.New()
	dl.Show()

	dl.Append(7)
	dl.Append(3)
	dl.Append(9)
	dl.Show()

	dl.Insert(1,5)
	dl.Insert(3,6)
	dl.Show()

	fmt.Println(dl.Delete(3))
	dl.Show()

	fmt.Println(dl.Delete(4))		// 删除最后一个节点
	dl.Show()
}


func main() {

	fmt.Println("start run...")

	// testSparseArray()				// 测试稀疏数组

	// testSequenList()					// 测试顺序表	

	// testLinkedList()					// 测试单链表

	// testRingList()					// 测试循环链表

	// testDoublyList()					// 测试双向链表


}
