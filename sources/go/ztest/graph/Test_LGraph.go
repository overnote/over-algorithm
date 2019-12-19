package graph

import (
	"fmt"
	"structure/graph/listgraph"
)

func TestLGraph() {

	fmt.Println("*********** 测试开始 ***********")

	g := listgraph.NewLGraphNoDirect(4, 5)
	//g.Display()				// 遍历方法

	//g.DFSTraverse()			// 深度优先
	g.BFSTraverse()				// 广度优先
	fmt.Println("*********** 测试结束 ***********")

}

