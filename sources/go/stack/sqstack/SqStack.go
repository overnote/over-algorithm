/*
 *  顺序栈
 */

package sqstack

import(
	"errors"
	"fmt"
)

// 顺序栈结构体
type SqStack struct {
	data		[]interface{}		// 节点数组（切片）
	top 		int					// 栈顶指针，其实就是当前索引位置
	size		int					// 栈的最大容量
}

func NewSqStack(size int) *SqStack{
	return &SqStack{
		data: 	nil,
		top:	-1,					// 默认-1，表示是空栈，到达0则是数组的第一个索引成为栈顶
		size: 	size,
	}
}

func (s *SqStack)Display() {

	if s.top == -1 {
		fmt.Println("空栈")
		return
	}

	fmt.Println("栈顶指针：", s.top)
	fmt.Println("栈数据：", s.data[:s.top+1])
}

// 进栈操作
func (s *SqStack)Push(e interface{}) error{
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

	// 正常情况
	if s.top == s.size - 1 {
		fmt.Println("栈已满")
		return errors.New("栈已满")
	}

	s.data[s.top + 1] = e
	s.top++
	return nil
}

// 出栈操作：返回出栈的值
func (s *SqStack)Pop() (interface{}, error){

	if s.top == -1 {
		fmt.Println("空栈无法pop")
		return nil, errors.New("空栈无法pop")
	}

	e := s.data[s.top]
	s.top--
	return e, nil

}