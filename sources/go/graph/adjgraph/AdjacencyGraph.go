package adjgraph

import (
	"fmt"
	"structure/queue/cirqueue"
)

type VertexType string 			// 顶点类型
type EdgeType 	int   			// 权值类型
const INFINITY EdgeType = 65535		// 最大权

type Graph struct {
	vexs		[]VertexType	// 顶点表
	arc   		[][]EdgeType	// 邻接矩阵
	numVertexes	int				// 图的顶点数
	numEdges	int				// 边数
}

// 创建无向网图
func NewMGraphNoDirect(numVertexes int, numEdges int) *Graph{

	// 记录顶点数和边数
	g := &Graph{
		vexs:        make([]VertexType, numVertexes),
		arc:         make([][]EdgeType, numVertexes),
		numVertexes: numVertexes,
		numEdges:    numEdges,
	}
	for k,_ := range g.arc {
		tempS := make([]EdgeType, numVertexes)
		g.arc[k] = tempS
	}

	// 记录顶点表
	fmt.Printf("请输入 %v 个顶点数据\n", numVertexes)
	for k,_ := range g.vexs {
		fmt.Scanf("%v", &g.vexs[k])
	}
	fmt.Println("g.vexs = ", g.vexs)

	// 记录边表
	for i := 0; i < g.numVertexes; i++ {
		for j := 0; j < g.numVertexes; j++ {

			// 对称位置值为0
			if i == j {
				g.arc[i][j] = 0
			} else {

				// 有权的位置记录权
				var w EdgeType
				fmt.Printf("当前位置为：[%v, %v]，请输入权(输入0或回车表示无权)： \n", i, j)
				fmt.Scanf("%v", &w)
				if w != 0 {
					g.arc[i][j] = w					// 手动记录权
				} else {
					g.arc[i][j] = INFINITY			//  记录默认权，即最大值
				}

			}

		}
	}
	fmt.Println("g.arc = ", g.arc)


	return g
}

// 深度优先搜索
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

	fmt.Printf("%v ", g.vexs[i])
	visited[i] = true	// 该顶点遍历完毕

	for j := 0; j < g.numVertexes; j++ {
		if g.arc[i][j] == 1 && !visited[j] {
			g.dfs(visited, j)
		}
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

			fmt.Printf("%v ", g.vexs[i])	// 打印遍历到的顶点
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

				for j := 0; j < g.numVertexes; j++ {
					if g.arc[i][k.(int)] == 1 && !visited[k.(int)] {
						visited[k.(int)] = true
						fmt.Printf("%v ", g.vexs[k.(int)])
						queue.EnQueue(k)
					}
				}

			}

		}

	}

	fmt.Println()
}