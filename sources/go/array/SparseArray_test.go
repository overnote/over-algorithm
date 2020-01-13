package array

import (
	"testing"
)

func TestSparseArray_TransToArray(t *testing.T) {
	// 生成一个二维原始数组
	originArr := [][]interface{}{
		{nil, nil, 7, nil},
		{nil, 6, nil, nil},
		{nil, nil, nil, nil},
	}
	t.Log("原始数组：", originArr)

	sparseArr := NewSparseArray(originArr)
	t.Log("压缩数组：", sparseArr)

	newOriginArr := sparseArr.TransToArray()
	t.Log("返回原始数组：", newOriginArr)
}