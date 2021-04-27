/**
 * 顺序表
 */

class SqList {
    // 构造顺序表
    constructor() {
        this.data = [] // JS数组不限制长度，无需考虑容量
        this.length = 0 // 该顺序表元素个数
    }

    // 增：按一个位置插入数据
    insert(e, index) {
        if (index < 1 || index > this.length + 1) {
            console.log('插入位置不合法')
            return
        }
        // 参数二表示要删除多少个数据，0表示不删除
        this.data.splice(index - 1, 0, e)
        this.length++
    }

    // 删：delete
    delete(index) {
        if (index < 1 || index > this.length) {
            console.log('删除位置不合法')
            return
        }
        if (l.length == 0) {
            console.log('空表无元素可删除')
            return
        }
        this.data.splice(index - 1, 1)
        this.length--
    }

    // 查：根据值查询位置
    searchValue(index) {
        if (index < 1 || index > this.length) {
            console.log('参数位置不合法')
            return
        }
        return this.data[index - 1]
    }

    // 查：根据值查询位置
    searchIndex(e) {
        return this.data.indexOf(e)
    }

    // 获取顺序表长度
    length() {
        return this.length
    }

    // 清空顺序表
    clear() {
        this.length = 0
    }

    // 销毁
    // js自动内存管理无需销毁

    // 显示顺序表
    display() {
        console.log(this.data.slice(0, this.length))
    }
}

// 测试
const l = new SqList()
l.insert(11, 1)
l.insert(12, 2)
l.insert(17, 7)
l.display()

l.insert(13, 3)
l.insert(14, 4)
l.insert(15, 5)
l.display()

l.insert(16, 6)
l.insert(5, 5)
l.display()

l.delete(3)
l.display()
