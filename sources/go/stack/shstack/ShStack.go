/*
 *  顺序栈：共享栈
 */

package shstack

import(
	//"errors"
	"errors"
	"fmt"
 )
 
 // 共享栈结构体
type ShStack struct {
	data		[]interface{}		// 节点数组（切片）
	top1 		int					// 栈1栈顶指针
	top2		int 				// 栈2栈顶指针
	size		int					// 栈的最大容量
 }
 
func NewShStack(size int) *ShStack{
	return &ShStack{
		data: 	nil,
		top1:	-1,
		top2:	size,
		size: 	size,
	}
 }
 
func (s *ShStack)Display() {
 
	if s.top1 == -1 {
		fmt.Println("栈1为空栈")
	} else {
		fmt.Println("栈1栈顶指针：", s.top1)
		fmt.Println("栈1数据：", s.data[:s.top1+1])
	}

	if s.top2 == s.size {
		fmt.Println("栈2为空栈")
	} else {
		fmt.Println("栈2栈顶指针：", s.top2)
		fmt.Println("栈2数据：", s.data[s.top2:])
	}
}
 
// 进栈操作
func (s *ShStack)Push(e interface{}, num int) error{		// num 表示元素 e Push到栈1还是栈2

	if num < 1 || num > 2 {
		fmt.Println("栈选择错误")
		return errors.New("栈选择错误")
	}

	/*
	在第一次push操作时，有两种选择：
	方式一： 使用s.data[e]设置值，但是go的空接口类型在初始化时值为nil，无法进行s.data[e] 操作
	方式二 ：使用go的append函数，但是这样实现的Push方法后，在Pop方法执行删除时也需要删除数据，
	完美的栈应该是不需要删除数据的，只用移动栈顶指针，对对应数据进行修改即可
	*/

	// 笔者这里做了如下优化：如果是第一次push，则将所有元素都初始化一遍
	if len(s.data) == 0 {
		s.data = make([]interface{}, s.size)
		for i := 0; i < s.size; i++ {
			s.data = append(s.data, e)
		}
		s.data = s.data[s.size:]
	}

	if s.top1 + 1 == s.top2 {
		fmt.Println("空间已满")
		return errors.New("空间已满")
	}

	// 对第一个栈进行入栈
	if num == 1 {
		s.data[s.top1 + 1] = e
		s.top1++
	} else {		// 对第二个栈进行入栈
		s.data[s.top2 - 1] = e
		s.top2--
	}

	return nil
}

// 出栈操作：返回出栈的值
func (s *ShStack)Pop(num int) (interface{}, error){

	if (num == 1 && s.top1 == -1) || (num ==2 && s.top2 == s.size) {
		fmt.Println("空栈无法pop")
		return nil, errors.New("空栈无法pop")
	}

	var e interface{}

	if num == 1 {
		e = s.data[s.top1]
		s.top1--
	} else {
		e = s.data[s.top2]
		s.top2++
	}

	return e, nil
}