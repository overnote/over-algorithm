package testStruct

import (
	"fmt"
	"structure/array/sparsearray"
)

func TestSparseArray() {

	fmt.Println("*********** 测试开始 ***********")

	// 生成一个二维原始数组
	originArr := [][]interface{}{
		{nil, nil, 7, nil},
		{nil, 6, nil, nil},
		{nil, nil, nil, nil},
	}
	fmt.Println("原始数组：", originArr)

	sparseArr := sparsearray.NewSparseArray(originArr)
	fmt.Println("压缩数组：", sparseArr)

	newOriginArr := sparseArr.TransToArray()
	fmt.Println("返回原始数组：", newOriginArr)

	fmt.Println("*********** 测试结束 ***********")

}
