/**
 * 单链表
 */

class Node {
    constructor(data) {
        this.data = data
        this.next = null
    }
}

class LinkList {
    constructor() {
        this.head = new Node(0)
    }
    // 增：插入节点
    // 约定：带头节点的链表，插入时，只能在头节点之后插入，也不允许插入超过最大元素个数的位置
    insert(e, index) {
        // 找到其前一个元素位置
        let p = this.locate(index - 1)
        if (!p) {
            return
        }
        // 创建要插入的节点
        let temp = new Node(e)
        temp.next = p.next
        p.next = temp

        this.head.data++
        return
    }
    // 删
    delete(index) {
        // 找到前一个元素位置
        let p = this.locate(index - 1)
        if (!p) {
            return
        }
        // 执行删除
        let tempData = p.next.data
        p.next = p.next.next
        this.head.data--
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
        let p = this.next
        while (p && p.data != e) {
            p = p.next
        }
        return p
    }
    //定位：根据位置查询节点地址
    locate(index) {
        if (index < 0 || index > this.head.data + 1) {
            console.log('获取位置不合法')
            return
        }
        let k = 0
        let p = this.head
        while (p && k < index) {
            p = p.next
            k++
        }
        return p
    }
    // 获取长度
    length() {
        return this.head.data
    }
    // 清空表
    clear() {
        this.head.data = 0
        this.head.next = null
    }
    // 显示单链表
    display() {
        if (this.head.data === 0) {
            console.log('空链表')
            return
        }

        let p = this.head
        let pos = 0
        let res = ''
        while (p) {
            if (pos === this.head.data) {
                res += p.data
                break
            } else {
                res += p.data
                res += '->'
                p = p.next
                pos++
            }
        }
        console.log(res)
    }
}

module.exports = LinkList
