/*
 * @Name: 双向链表
 * @Description: 带头结点的链表
 * 单向链表、单向循环链表、双向链表、双向循环链表
 * @Author: yaya
 * @Date: 2020-01-13 19:47:05
 * @Editors: yaya
 */
class Node{
    constructor(data){
        this.data = data;
        this.next = null;
        this.prev = null
    }
}

class DoubleLinkList{
    constructor(){
        this.head = new Node(0);
    }

    length(){
        return this.head.data;
    }

    pushBack(el){
        let newNode = new Node(el);

        if(this.head.next == null){
            this.head.next = newNode;
            newNode.prev = this.head;
            this.head.data++;
            return;
        }

        let currentNode = this.head.next;
        
        while(currentNode.next != null){
            let prevNode = currentNode;
            currentNode = currentNode.next;
            currentNode.prev = prevNode;
        }

        currentNode.next = newNode;
        newNode.prev = currentNode;
        this.head.data++;

    }

    pushFront(el){
        let newNode = new Node(el);
        newNode.prev = this.head;

        if(this.head.next == null){
            this.head.next = newNode;
            this.head.data++;
            return;
        }

        let firstNode = this.head.next;
        
        newNode.next = firstNode;
        firstNode.prev = newNode;
        this.head.next = newNode;
        this.head.data++;

    }

    pushIndex(el,index){
        let newNode = new Node(el);

        if(index<1 || index>this.length()+1){
            console.log('参数不符');
            return;
        }
        if(index==1){
            this.pushFront(el);
            return;
        }
        if(index==this.length()+1){
            this.pushBack(el);
            return;
        }

        let prevNode = this.head;
        let currentNode = null;

        for(let i=1; i<index; i++){
            let node = prevNode;
            prevNode = prevNode.next;
            prevNode.prev = node;
        }

        currentNode = prevNode.next;

        newNode.next = currentNode;
        currentNode.prev = newNode;
        prevNode.next = newNode;
        newNode.prev = prevNode;
        this.head.data++;

    }
}

module.exports = DoubleLinkList;