/*
 *  链栈
 */

package linkstack

import(
	"errors"
	"fmt"
)

// 链栈节点
type node struct {
	data 		interface{}
	next		*node
}

// 链栈结构体
type LinkStack struct {
	top 		*node				// 栈顶指针
	length		int					// 栈的长度
}

func NewLinkStack() *LinkStack{
	return &LinkStack{
		top:	nil,
		length:	0,
	}
}

func (s *LinkStack)Display() {

	if s.top == nil {
		fmt.Println("空栈")
		return
	}

	fmt.Println("栈顶指针数据：", s.top.data)
	fmt.Printf("栈数据：")
	currentNode := s.top
	for i := 1; i <= s.length; i++ {
		fmt.Printf("%d ", currentNode.data)
		currentNode = currentNode.next
	}
	fmt.Println()
}

// 进栈操作
func (s *LinkStack)Push(e interface{}) {
	node := &node{
		data: e,
		next: s.top,
	}
	s.top = node
	s.length++
}

// 出栈操作：返回出栈的值
func (s *LinkStack)Pop() (interface{}, error){

	if s.top == nil  {
		fmt.Println("空栈无法pop")
		return nil, errors.New("空栈无法pop")
	}

	e := s.top.data
	node := s.top
	s.top = node.next
	s.length--
	return e, nil

}