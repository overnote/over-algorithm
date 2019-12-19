package graph

import (
	"fmt"
	"structure/graph/adjgraph"
)

func TestMGraph() {

	fmt.Println("*********** 测试开始 ***********")

	adjgraph.NewMGraphNoDirect(4,6)

	fmt.Println("*********** 测试结束 ***********")

}

