/**
 * 顺序表
 */

class SqList {
    // 构造顺序表
    constructor() {
        this.data = [] // JS数组不限制长度，无需考虑容量
        this.size = 0 // 该顺序表元素个数
    }

    // 增：按一个位置插入数据
    insert(e, index) {
        if (index < 1 || index > this.size + 1) {
            console.log('插入位置不合法')
            return
        }
        // 参数二表示要删除多少个数据，0表示不删除
        this.data.splice(index - 1, 0, e)
        this.size++
    }

    // 删：delete
    delete(index) {
        if (index < 1 || index > this.size) {
            console.log('删除位置不合法')
            return
        }
        if (this.size == 0) {
            console.log('空表无元素可删除')
            return
        }
        let e = this.data[index - 1]
        this.data.splice(index - 1, 1)
        this.size--
        return e
    }

    // 改
    update(index, e) {
        if (index < 1 || index > this.size) {
            console.log('修改位置不合法')
            return
        }
        this.data[index - 1] = e
    }

    // 查：根据值查询位置
    search(index) {
        if (index < 1 || index > this.size) {
            console.log('参数位置不合法')
            return
        }
        return this.data[index - 1]
    }

    // 查：根据值查询位置
    locate(e) {
        return this.data.indexOf(e)
    }

    // 获取顺序表长度
    length() {
        return this.size
    }

    // 清空顺序表
    clear() {
        this.size = 0
    }

    // 销毁
    // js自动内存管理无需销毁

    // 显示顺序表
    display() {
        console.log(this.data.slice(0, this.size))
    }
}
