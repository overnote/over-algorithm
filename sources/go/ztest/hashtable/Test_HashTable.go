package hashtable

import (
	"fmt"
	"structure/hashtable"
)

func TestHashTable() {

	fmt.Println("*********** 测试开始 ***********")

	fmt.Println("----------- 测试哈希值 -----------")

	c1 := hashtable.HashCode("li")
	c2 := hashtable.HashCode("zs")
	fmt.Println(c1)
	fmt.Println(c2)

	fmt.Println("*********** 测试结束 ***********")
}