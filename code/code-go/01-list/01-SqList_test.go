package list_test

import (
	list "code/01-list"
	"fmt"
	"testing"
)



func TestSqList(t *testing.T){
	l := list.NewSqList()

	fmt.Println("----执行了初始化操作----");
	l.Display()

	l.InsertElem(1, 5)
	l.InsertElem(2, 6)
	l.InsertElem(3, 1)
	l.InsertElem(4, 2)
	l.InsertElem(5, 3)
    fmt.Println("----执行了新增操作----");
	l.Display()

    e1,_ := l.DeleteElem(5)
    fmt.Println("----执行了删除操作----");
    fmt.Println("删除位置5 e1 = ", e1);
    l.Display()

    l.UpdateElem(4, 10);
    fmt.Println("----执行了更新操作----");
    l.Display()

   	e2,_ := l.GetElem(3);
    fmt.Println("----执行了查找操作----");
    fmt.Println("按位查找的e2 = ", e2);

    idx1,_ := l.LocateElem(1);
    fmt.Println("----执行了查找操作----");
    fmt.Println("按位查找的idx1 = ", idx1);

    l.Clear();
    fmt.Println("----执行了清空操作----");
    l.Display()

}