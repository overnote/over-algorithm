/**
 * 栈:顺序栈
 */

package SequenStack

import "fmt"

type SequenStack struct {
	data   []interface{}
}

func New() *SequenStack {
	s := make([]interface{}, 0)
	return &SequenStack{
		s,
	}
}

// 压栈
func (ss *SequenStack) Push(data interface{}) {
	ss.data = append(ss.data, data)
}

// 出栈
func (ss *SequenStack) Pop() interface{} {

	if ss.Length() == 0 {
		fmt.Println("stack is empty")
		return nil
	}

	index := ss.Length() - 1
	value := ss.data[index]
	ss.data = append(ss.data[:index])

	return value
}

// 获取栈长度
func (ss *SequenStack) Length() int {
	return len(ss.data)
}

// 展示栈
func (ss *SequenStack) Show() {

	if ss.Length() == 0 {
		fmt.Println("stack is empty")
		return
	}

	for i := 0; i < ss.Length(); i++ {
		if i == ss.Length() - 1 {
			fmt.Println(ss.data[i])
		} else {
			fmt.Print(ss.data[i], " ")
		}
		
	}

}