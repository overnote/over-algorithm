/**
 * 循环链表
 */

 package list

 import (
	 "errors"
	 "fmt"
 )
 
 // 节点对象，存储循环链表上某个节点数据
 type node struct {
	 data interface{}			// 数据域
	 next *node					// 指针域
 }
 
 // 循环链表对象：存储头节点即可
 type RingList struct {
	 head *node					
 }
 
 // 创建循环链表
 func New() *RingList {
	 head := &node{0, nil}		// 头节点存储链表中元素的个数
	 head.next = head
	 return &RingList{
		 head,
	 }
 }
 
 // 判断空
 func (ll *RingList)IsEmpty() bool {
	 var len int = 0
	 len = ll.head.data.(int)
	 return len == 0
 }
 
 // 增加：末尾添加
 func (rl *RingList) Append(data interface{}){
 
	 insertNode := &node{data, rl.head}					// 要插入的节点
	 var len int = 0
	 len = rl.head.data.(int)
 
	 // 查询最后一个节点
	 lastNode := rl.head.next

	 if len == 0 {								// 第一次添加
		
		insertNode.next = rl.head
		rl.head.next = insertNode

	 } else {									// 不是第一次添加
		for lastNode.next != rl.head {					
			lastNode = lastNode.next
		}
		lastNode.next = insertNode
	 }

	 len ++
	 rl.head.data = len

	 return
 }
 
 // 增加：任意位置插入结点
 func (rl *RingList) Insert(index int, data interface{}) error{
 
	 var len int = 0
	 len = rl.head.data.(int)
 
	 if index < 1 || index > len {
		 return errors.New("index overflow")
	 }

	 if index == len + 1 {					 // 如果是在末尾添加			
		rl.Append(data)
	 } else {								// 如果不是在末尾添加
	
		beforeNode := rl.head
		appendNode := &node{data, nil}
 
		for i := 0; i < index - 1; i++ {
			beforeNode = beforeNode.next		// 找到要插入的位置的前置元素
		}
 
		appendNode.next = beforeNode.next
		beforeNode.next = appendNode
	 }					
 
	 len ++
	 rl.head.data = len
 
	 return nil
 
 }

 // 删除：删除末尾节点
 func (rl *RingList) Pop() (interface{}, error) {

	var len int = 0
	len = rl.head.data.(int)

	 if len == 0 {
		 return nil, errors.New("list is empty")
	 }

	 currentNode := rl.head
	 for currentNode.next != rl.head {
		 currentNode = currentNode.next
	 }

	 len --
	 rl.head.data = len
	 currentNode.next = rl.head
	 return currentNode.data, nil

 }
 
 // 删除：删除指定位置结点
 func (rl *RingList) Delete(index int) (interface{}, error) {

	var len int = 0
	len = rl.head.data.(int)
 
	 if index <= 0 || index > len {
		 return nil, errors.New("index overflow")
	 }

	 if len == 0 {
		 return nil, errors.New("list is empty")
	 }

	 if index == len {
		 return rl.Pop()
	 }
 
	 beforeNode := rl.head
	 var delData interface{}					// 被删除的数据
 
	 for i := 0; i < index - 1; i++ {
		beforeNode = beforeNode.next
	 }
 
	 beforeNode.next = beforeNode.next.next
 
	 len--
	 rl.head.data = len
 
	 return delData, nil
 }
 
 // 查询： 获取指定位置结点
 func (rl *RingList) Node(index int) (interface{}, error) {
 
	 var len int = 0
	 len = rl.head.data.(int)
 
	 if index < 0 || index > len {
		 return nil, errors.New("index overflow")
	 }
	 currentNode := rl.head
	 for i := 0; i < index; i++ {
		 currentNode = currentNode.next
	 }
	 return currentNode.data, nil
 }
 
 // 打印链表
 func (rl *RingList) Show() {
	 var len int = 0
	 len = rl.head.data.(int)
	 currentNode := rl.head
	 for i := 0; i <= len; i++ {
		 fmt.Print(currentNode.data)
		 if i == len {
			 fmt.Print(" \n")
		 } else {
			 fmt.Print(" ")
		 }
 
		 currentNode = currentNode.next
	 }
 }

 // 获取长度
 func (rl *RingList)Length() int {
	 var len int = 0
	 len = rl.head.data.(int)
	 return len
 }

 // 获取头结点
 func (rl *RingList)GetHead() *node {
	return rl.head
 }

 // 获取尾节点
 func (rl *RingList)GetTail() *node {
	if rl.Length() == 0 {
		return nil
	}
	currentNode := rl.GetHead()
	for currentNode.next != rl.GetHead() {
		currentNode = currentNode.next
	} 
	return currentNode
 }