package adjgraph

import "fmt"

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
func NewMGraph(numVertexes int, numEdges int) *Graph{

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