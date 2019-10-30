import SqList from "../list/SqList";

function TestSqList(){
    console.log("********** 顺序表测试开始 **********");

    let l = new SqList();
    l.display()

    console.log("增加元素-------------");
    l.insert(1, 17)
    l.insert(2, 18)
    l.insert(3, 19)
    console.log("长度：", l.len())
    l.display()

    console.log("修改元素-------------");
    l.update(2, 22)
    l.display()

    let pe = l.priorElem(22)
    let ne = l.nextElem(22)
    console.log("查找第二个元素前驱：", pe);
    console.log("查找第二个元素后继：", ne);

    console.log("删除元素-------------");
    l.delete(2)
    l.display()

    console.log("********** 顺序表测试结束 **********");
}

export default TestSqList