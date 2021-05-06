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
        this.size = 0
    }
    // 增：插入结点
    // 约定：带头结点的链表，插入时，只能在头结点之后插入，也不允许插入超过最大元素个数的位置
    insert(e, index) {
        if (index < 1 || index > this.size + 1) {
            console.log('插入位置非法')
            return
        }

        // 找到插入位置前一个位置：也可以使用 locate函数
        let p = this.head
        let k = 0
        while (p.next != null && k < index - 1) {
            p = p.next
            k++
        }

        // 创建要插入的结点
        let q = new Node(e)
        q.next = p.next
        p.next = q

        this.size++
    }
    // 删
    delete(index) {
        if (index < 1 || index > this.size) {
            console.log('删除位置非法')
            return
        }

        // 找到插入位置前一个位置：也可以使用 locate函数
        let p = this.head
        let k = 0
        while (p.next != null && k < index - 1) {
            p = p.next
            k++
        }

        // 执行删除
        let tempData = p.next.data
        p.next = p.next.next

        this.size--
        return tempData
    }
    // 改
    update(index, e) {
        // 找到修改位置：这里使用 locate 函数
        let p = this.locate(index)
        if (!p) {
            return
        }
        p.data = e
    }
    // 查
    search(e) {
        let p = this.head
        while (p.next) {
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
    //定位：根据位置查询结点地址
    locate(index) {
        if (index < 0 || index > this.size + 1) {
            console.log('获取位置不合法')
            return
        }

        let p = this.head
        let k = 0
        while (p.next && k < index) {
            p = p.next
            k++
        }
        return p
    }
    // 获取长度
    length() {
        return this.size
    }
    // 清空表
    clear() {
        this.size = 0
        this.head.next = null
    }
    // 显示单链表
    display() {
        if (this.size === 0) {
            console.log('空链表')
            return
        }

        let p = this.head
        let pos = 0
        let res = ''
        while (p) {
            if (pos === this.size) {
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
