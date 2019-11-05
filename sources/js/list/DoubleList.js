/**
 * 单链表
 */

class Node {
    constructor(){
        this.data = null;
        this.next = null;
        this.prev = null;
    }
}

class DoubleList {
    constructor(){
        this.len = 0;
        this.first = null;
    }
    length(){
        return this.len;
    }
    display(){
        if (this.len == 0) {
            console.log("数据结构内无元素");
            return;
        }
        let currentNode = this.first;
        for (let i = 1; i <= this.len; i++) {
            console.log("位置 " + i + "元素为 " + currentNode.data);
            if (i == this.len) {
                break
            }
            currentNode = currentNode.next;
        }
    }
    append(e){

        let insertNode = new Node();
        insertNode.data = e;

        // 第一次追加
        if (this.len == 0) {
            insertNode.next = insertNode;
            insertNode.prev = insertNode;
            this.first = insertNode;
            this.len++;
            return;
        }

        let currentNode = this.first;
        while(currentNode.next != this.first){
            currentNode = currentNode.next;
        }
        currentNode.next = insertNode;
        insertNode.prev = currentNode;
        insertNode.next = this.first;
        this.first.prev = insertNode;
        this.len++;
    }
    insert(index, e){

        if (index < 1 || index > this.len) {
            console.log("位序不正确");
            return
        }

        // 构造要插入的节点
        let insertNode = new Node();
        insertNode.data = e;

        // 找到插入位置的前一个节点
        let i = 1;
        let currentNode = this.first;
        while(currentNode.next != this.first) {
            if (i == index) {
                break;
            }
            i++;
            currentNode = currentNode.next;
        }

        // 插入
        let prevNode = currentNode.prev;
        insertNode.prev = prevNode;
        insertNode.next = currentNode;

        prevNode.next = insertNode;
        currentNode.prev = insertNode;
        this.len++;

        // 考虑头指针改变
        if (index == 1) {
            this.first = insertNode;
        }
    }
    delete(index){      // 按照位序删除

        if (index < 1 || index > this.len + 1) {
            console.log("位序不正确");
            return
        }

        // 如果要删除的是第一个元素
        if (index == 1 && this.len == 1) {
            this.first = null;
            this.len = 0;
            return
        }

        // 找到要删除元素的元素
        let i = 1;
        let currentNode = this.first;
        while(currentNode.next != this.first) {
            if (i == index) {
                break;
            }
            i++;
            currentNode = currentNode.next;
        }

        // 删除
        let prevNode = currentNode.prev;
        let nextNode = currentNode.next;
        prevNode.next = nextNode;
        nextNode.prev = prevNode;
        this.len--;

        // 考虑头指针改变
        if (index == 1) {
            this.first = nextNode;
        }
    }
    update(index, e){

        if(index < 1 || index > this.len) {
            console.log("位序不正确");
            return
        }

        let i = 1;
        let currentNode = this.first;
        while(currentNode.next != this.first) {
            if (i == index) {
                break;
            }
            i++;
            currentNode = currentNode.next;
        }

        currentNode.data = e;
    }
    getElem(index){          // 按照位序查询
        if (index < 1 || index > this.len) {
            console.len("位序不合法");
            return
        }
        let i = 1;
        let currentNode = this.first;
        while(currentNode.next != this.first) {
            if (i == index) {
                break;
            }
            i++;
            currentNode = currentNode.next;
        }
        return currentNode.data
    }
    locateElem(e){      // 按照值查询位序
        let i = 1;
        let currentNode = this.first;
        while(currentNode.next  != this.first) {
            if (currentNode.data == e) {
                break;
            }
            i++;
            currentNode = currentNode.next;
        }
        if (i == this.len && currentNode.data != e) {
            console.log("未找到元素");
            return
        }
        return i
    }
    priorElem(e) {      // 查找前驱

        if (this.len <= 1) {
            console.log("数据结构元素不足，无法查询");
            return
        }

        if (this.first.data == e) {
            console.log("首元素无前驱");
            return
        }

        let i = 1;
        let currentNode = this.first;
        while(currentNode.next != this.first) {
            if (currentNode.next.data == e) {
                return currentNode.data;
            }
            i++;
            currentNode = currentNode.next;
        }
        console.log("未找到元素");
    }
    nextElem(e){             // 查询后继

        if (this.len <= 1) {
            console.log("数据元素不足，无法查询");
            return
        }

        let i = 1;
        let currentNode = this.first;
        while(currentNode.next != this.first) {
            if (currentNode.data == e) {
                break;
            }
            i++;
            currentNode = currentNode.next;
        }

        if (i == this.len && currentNode.data != e) {
            console.log("未找到传入元素");
            return
        }

        if (i == this.len && currentNode.data == e) {
            console.log("最后元素无后继");
            return
        }

        return currentNode.next.data;

       
    }
    clear(){
        this.first = null;
        this.len = 0;
    }
}

export default DoubleList