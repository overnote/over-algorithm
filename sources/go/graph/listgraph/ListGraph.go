package listgraph

import "fmt"

// 顶点
type vertex struct {
	data 		interface{}			// 顶点存储的数据
	head 		*node				// 顶点的边表头指针
}

// 边节点类型
type node struct {
	adjvex		int				// 邻接点域，即在顶点数组中的索引
	weight		int				// 权
	next 		*node			// 链表域
}

// 图
type Graph struct {
	vexs		[]vertex		// 顶点数组
	numVertexes	int 			// 当前定点数
	numEdges	int				// 当前边数
}

// 创建无向网图
func NewLGraph(numVertexes int, numEdges int) *Graph{

	// 记录顶点数和边数
	g := &Graph{
		vexs:        make([]vertex, numVertexes),
		numVertexes: numVertexes,
		numEdges:    numEdges,
	}

	// 记录顶点表
	fmt.Printf("请输入 %v 个顶点数据\n", numVertexes)
	for k,_ := range g.vexs {
		var e int
		fmt.Scanf("%v", &e)
		g.vexs[k].data = e
		g.vexs[k].head = nil
	}
	fmt.Println("g.vexs = ", g.vexs)

	// 记录边表：每条边对应2个顶点，用户手动输入这2个顶点序号
	for k := 0; k < numEdges; k++ {

		var w int
		fmt.Println("输入当前边的权")
		fmt.Scanf("%d", &w)

		fmt.Printf("输入第 %d 条边的两个顶点序号：\n", k + 1)
		var i,j int
		fmt.Scanf("%d %d", &i, &j)

		// 生成第一个链表节点
		nodeI := &node{
			adjvex:		j,
			weight: 	w,
			next:   	g.vexs[i].head,
		}
		// 将当前顶点的指针指向 node
		g.vexs[i].head = nodeI

		// 生成第二个链表节点
		nodeJ := &node{
			adjvex:		i,
			weight: 	w,
			next:   	g.vexs[j].head,
		}
		// 将当前顶点的指针指向 node
		g.vexs[j].head = nodeJ
	}

	return g
}

// 显示图
func (g *Graph)Display(){

	for i := 0; i < g.numVertexes; i++ {
		currentNode := g.vexs[i].head
		fmt.Printf("%v：", g.vexs[i].data)
		for currentNode.next != nil {
			fmt.Printf("%v ", g.vexs[currentNode.adjvex].data)
			currentNode = currentNode.next
		}
		fmt.Println(g.vexs[currentNode.adjvex].data)
	}

}