/**
 * 顺序表
 * 为了简便，这里不考虑数组扩容，缩容问题，简单申请一个数组即可
 */

class SqList {
    length
    data
    constructor() {
        this.length = 0
        this.data = []
    }

    length() {
        return this.length
    }

    clear() {
        this.length = 0
    }

    display() {
        if (this.length <= 0) {
            console.log([])
            return
        }
        console.log(this.data)
    }

    insertElem(idx, e) {
        if (idx < 1 || idx > this.length + 1) {
            console.log('插入位置不合法')
            return false
        }

        // 移动元素
        for (let i = this.length - 1; i >= idx - 1; i--) {
            this.data[i + 1] = this.data[i]
        }

        // 插入位置值改变
        this.data[idx - 1] = e
        this.length++
        return true
    }

    deleteElem(idx) {
        if (idx < 1 || idx > this.length) {
            console.log('删除位置不合法')
            return false
        }

        const e = this.data[idx - 1]
        this.data.splice(idx - 1, 1)
        this.length--
        return e
    }

    updateElem(idx, e) {
        if (idx < 1 || idx > this.length) {
            console.log('修改位置不合法')
            return false
        }
        this.data[idx - 1] = e
        return true
    }

    getElem(idx) {
        if (idx > this.length || idx < 1) {
            console.log('按位查找:越界')
            return false
        }
        return this.data[idx - 1]
    }

    locateElem(e) {
        for (let i = 0; i < this.length; i++) {
            if (this.data[i] == e) {
                return i + 1
            }
        }
        console.log('按值查找:未找到')
        return false
    }
}

function testSqList() {
    const l = new SqList()

    console.log('----执行了初始化操作----')
    l.display()

    l.insertElem(1, 5)
    l.insertElem(2, 6)
    l.insertElem(3, 1)
    l.insertElem(4, 2)
    l.insertElem(5, 3)
    console.log('----执行了新增操作----')
    l.display()

    let e1 = l.deleteElem(5)
    console.log('----执行了删除操作----')
    console.log('删除位置5 e1 = ', e1)
    l.display()

    l.updateElem(4, 10)
    console.log('----执行了更新操作----')
    l.display()

    let e2 = l.getElem(3)
    console.log('----执行了查找操作----')
    console.log('按位查找的e2 = ', e2)

    let idx1 = l.locateElem(1)
    console.log('----执行了查找操作----')
    console.log('按位查找的idx1 = ', idx1)

    l.clear()
    console.log('----执行了清空操作----')
    l.display()
}

testSqList()
