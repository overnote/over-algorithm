package array

import(
	"fmt"
)

type Node struct {
	row int
	col int
	val int
}

func NewSparseArray(arr [11][11]int) []Node{

	var sa []Node
	n := Node{
		row: len(arr),
		col: len(arr[0]),
		val: 0,
	}
	sa = append(sa, n)

	for i1, v1 := range arr {
        for i2, v2 := range v1 {
            if v2 != 0 {
                n := Node{
                    row : i1,
                    col : i2,
                    val : v2,
                }
                sa = append(sa ,n)
            }
        }
	}

	return sa
}

// 打印原始数组
func ShowOrigin(arr [11][11]int) {
	for i1, v1 := range arr {
        for i2, v2 := range v1 {
            fmt.Printf("%d|%d=%d\t", i1, i2, v2)
        }
        fmt.Println()
    }
}

// 打印稀疏数组
func ShowSparse(n []Node) {
	fmt.Println(n)
}

// 稀疏数组转普通数组
func TransToArray(n []Node) (arr [11][11]int){
    for i, v := range n {
        if i == 0 {
            continue
        }
        arr[v.row][v.col] = v.val
	}
	return
}