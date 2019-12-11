package teststruct

import (
	"fmt"
	"structure/graph/listgraph"
)

func TestLGraph() {

	fmt.Println("*********** 测试开始 ***********")

	g := listgraph.NewLGraph(4, 5)
	g.Display()

	fmt.Println("*********** 测试结束 ***********")

}

