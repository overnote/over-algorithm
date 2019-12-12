package listgraph

import (
	"fmt"
	"structure/queue/cirqueue"
)

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
func NewLGraphNoDirect(numVertexes int, numEdges int) *Graph{

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
		//fmt.Println("输入当前边的权")		// 需要加权则打开这两行代码
		//fmt.Scanf("%d", &w)

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

// 深度优先遍历
func (g *Graph)DFSTraverse() {

	// 设置一个数组，记录所有顶点的访问状态
	visited := make([]bool, g.numVertexes)
	for k, _ := range visited {
		visited[k] = false		// 所有顶点都未被访问
	}

	// 对所有未访问的顶点调用 深度优先算法 dfs。如果是连通图，只会执行一次
	for i := 0; i < g.numVertexes; i++{
		if !visited[i] {
			fmt.Printf("本次遍历顶点：")
			g.dfs(visited, i)			// 调用深度优先算法
			fmt.Println()
		}
	}

}

// 深度优先算法
func (g *Graph)dfs(visited []bool, i int){

	fmt.Printf("%v ", g.vexs[i].data)
	visited[i] = true	// 该顶点遍历完毕

	currentNode := g.vexs[i].head
	for currentNode.next != nil {
		// 对未访问的邻接点进行递归调用
		if !visited[currentNode.adjvex] {
			g.dfs(visited, currentNode.adjvex)
		}
		currentNode = currentNode.next
	}
}

// 广度优先搜索
func (g *Graph)BFSTraverse() {

	// 设置一个数组，记录所有顶点的访问状态
	visited := make([]bool, g.numVertexes)
	for k, _ := range visited {
		visited[k] = false		// 所有顶点都未被访问
	}

	queue := cirqueue.NewCirQueue(50)

	for i := 0; i < g.numVertexes; i++ {

		if !visited[i] {

			visited[i] = true							// 当前遍历到了该节点

			fmt.Printf("%v ", g.vexs[i].data)	// 打印遍历到的顶点
			err := queue.EnQueue(i)
			if err != nil {
				fmt.Println("入队发生错误：", err)
				return
			}

			for queue.Length() != 0{

				k, err := queue.DeQueue()
				if err != nil {
					fmt.Println("出队发生错误：", err)
					return
				}

				currentNode := g.vexs[k.(int)].head
				for currentNode != nil {

					if !visited[currentNode.adjvex] {
						visited[currentNode.adjvex] = true
						fmt.Printf("%v ", g.vexs[currentNode.adjvex].data)
						err = queue.EnQueue(currentNode.adjvex)
						if err != nil {
							fmt.Println("入队发生错误：", err)
							return
						}
					}
					currentNode = currentNode.next

				}
			}

		}

	}

	fmt.Println()
}
