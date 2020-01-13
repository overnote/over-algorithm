/*
 *  稀疏数组
 */
package array

type SparseNode struct {
	row 	int
	col 	int
	val 	interface{}
}

type SparseArray struct {
	data 			[]SparseNode
	lengthRow		int
	lengthCol 		int
}

func NewSparseArray(originArr [][]interface{}) *SparseArray{

	var sparseArr []SparseNode

	for i1, v1 := range originArr {
        for i2, v2 := range v1 {
            if v2 != nil {
                tempNode := SparseNode{
                    row : i1,
                    col : i2,
                    val : v2,
                }
				sparseArr = append(sparseArr , tempNode)
            }
        }
	}

	return &SparseArray{
		data: 			sparseArr,
		lengthRow:  	len(originArr),
		lengthCol: 		len(originArr[0]),
	}
}

// 稀疏数组转普通数组
func (sa *SparseArray)TransToArray() [][]interface{}{

	// 构建一个二维切片
	originArr := make([][]interface{},  sa.lengthRow)
	for k, _ := range originArr {
		resultArr := make([]interface{}, sa.lengthCol)
		originArr[k] = resultArr
	}

    for _, v := range sa.data {
		originArr[v.row][v.col] = v.val
	}
	return originArr
}
