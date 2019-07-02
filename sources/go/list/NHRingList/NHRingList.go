/**
 * 循环链表：无头节点
 */

 package list

 import (
	 "fmt"
 )
 
 // 节点对象，存储循环链表上某个节点数据
 type node struct {
	 Data interface{}			// 数据域
	 Next *node					// 指针域
 }
 
 // 循环链表对象：存储头指针与长度
 type NHRingList struct {
	 Head *node	
	 Length int					
 }
 
 // 创建循环链表
 func New() *NHRingList {
	 head := &node{nil, nil}		// 头节点存储链表中元素的个数
	//  head.Next = head				// 循环指向自己
	 return &NHRingList{
		 head,
		 0,
	 }
 }
 
 // 增加：末尾添加
 func (rl *NHRingList) Append(data interface{}){
 
	 insertNode := &node{data, rl.Head}					// 要插入的节点

	 if rl.Length == 0 {								// 第一次添加
		rl.Head = insertNode
		rl.Head.Next = insertNode
	 } else {									// 不是第一次添加
		tailNode := rl.Head
		for tailNode.Next != rl.Head {					
			tailNode = tailNode.Next			// 找到最后一个节点
		}
		tailNode.Next = insertNode
	 }

	 rl.Length++
	 return
 }

 // 打印链表
 func (rl *NHRingList) Show() {

	if rl.Length == 0 {
		fmt.Println("list is empty")
		return
	}

	 currentNode := rl.Head
	 for i := 0; i <= rl.Length - 1; i++ {
		 fmt.Print(currentNode.Data)
		 if i == rl.Length - 1 {
			 fmt.Print(" \n")
		 } else {
			 fmt.Print(" ")
		 }
 
		 currentNode = currentNode.Next
	 }
 }