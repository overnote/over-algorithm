/**
	并查集： Quick Union方式实现
 */

package union

import "fmt"

type Union struct {
	parents 	[]int	// 存储其父元素的切片
	cap			int
}

// 初始化：每个元素各自属于一个单元素集合
func NewUnion(cap int) *Union{

	if cap <= 0 {
		fmt.Println("参数非法")
		return nil
	}

	parents := make([]int, cap)
	for i := 0; i < cap; i++{
		parents[i] = i
	}
	return &Union{
		parents: parents,
		cap:     cap,
	}
}

// 查找：查找元素所在集合，即v的根节点，即所在集合
func (u *Union)Find(v int) int{
	if v < 0 || v >= u.cap {
		fmt.Println("参数非法")
		return 0
	}
	for v != u.parents[v] {
		v = u.parents[v]			// v不断的向上查找
	}
	return v
}

// 合并: 约定v1的父节点改为v2的父节点
func (u *Union)Union(v1 int, v2 int) {
	p1 := u.Find(v1)
	p2 := u.Find(v2)
	if p1 == p2 {
		return
	}
	u.parents[p1] = p2
}