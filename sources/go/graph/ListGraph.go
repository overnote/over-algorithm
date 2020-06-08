/**
	图：邻接表实现
 */

package graph

// 顶点
type ListVertex struct {
	value 	interface{}
}

// 边
type ListEdge struct {
	from	*ListVertex
	to 		*ListVertex
	weight	int					// 边的权
}