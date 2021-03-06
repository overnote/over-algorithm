/**
 * 双向循环链表
 */

class Node {
    constructor(data) {
        this.data = data
        this.prev = null
        this.next = null
    }
}

class DblList {
    constructor() {
        let p = new Node(0)
        p.prev = p
        p.next = p
        this.head = p
        this.size = 0
    }
    // 增：插入节点。约定只能在头节点之后插入，且不超过最大元素个数位置
    insert(e, index) {
        if (index < 1 || index > this.size + 1) {
            console.log('插入位置不合法')
            return
        }

        // 找到插入位置前一个位置：也可以使用 Locate函数
        let p = this.head
        let k = 0
        while (p.next != this.head && k < index - 1) {
            p = p.next
            k++
        }

        // 创建要插入的节点
        let q = new Node(e)
        q.prev = p
        q.next = p.next

        // 执行插入
        p.next = q
        this.size++
    }
    // 删：根据位置删除，返回被删除的元素
    delete(index) {
        if (index < 1 || index > this.size) {
            console.log('删除位置非法')
            return
        }
        // 找到删除位置前一个位置：也可以使用 Locate函数
        let p = this.head
        let k = 0
        while (p.next != this.head && k < index - 1) {
            p = p.next
            k++
        }
        // 获取删除数据
        let q = p.next
        let tempData = q.data

        // 执行删除
        q.next.prev = p
        p.next = q.next
        this.size--
        return tempData
    }
    // 改
    update(index, e) {
        let p = this.locate(index)
        if (!p) {
            return
        }
        p.data = e
    }
    // 查
    search(e) {
        let p = this.head
        while (p.next != this.head) {
            if (p.data == e) {
                break
            }
            p = p.next
        }
        if (p.data == e) {
            return p
        } else {
            return null
        }
    }
    // 定位
    locate(index) {
        if (index < 0 || index > this.size + 1) {
            console.log('获取位置不合法')
            return
        }

        let p = this.head
        let k = 0
        while (p.next != this.head && k < index) {
            p = p.next
            k++
        }
        return p
    }
    // 根据值快速找到前驱
    prevElem(e) {
        let p = this.search(e)
        if (!p) {
            console.log('未找到该元素')
            return
        }
        return p.prev
    }
    // 根据值快速找到后继
    nextElem(e) {
        let p = this.search(e)
        if (!p) {
            console.log('未找到该元素')
            return
        }
        return p.next
    }
    // 获取长度
    length() {
        return this.size
    }
    // 清空表
    clear() {
        this.head.next = this.head
        this.head.prev = this.head
        this.size = 0
    }
    // 显示表
    display() {
        if (this.size == 0) {
            console.log('空链表')
            return
        }

        let p = this.head
        let pos = 0
        let res = ''
        while (p.next) {
            if (pos == this.size) {
                res += p.data
                res += '<->'
                res += this.head.data
                res += '<->...'
                break
            }
            res += p.data
            res += '<->'
            p = p.next
            pos++
        }
        console.log(res)
    }
}

module.exports = DblList
