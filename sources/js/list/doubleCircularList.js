/*
 * @Name: 双向环形链表
 * @Description: 带头结点的链表
 * 单向链表、单向循环链表、双向链表、双向循环链表
 * @Author: yaya
 * @Date: 2020-01-14 10:47:49
 * @Editors: yaya
 */
class Node{
    constructor(data){
        this.data = data;
        this.prev = null;
        this.next = null;
    }
}

class DoubleCircularList{
    constructor(){
        this.head = new Node(0);
        this.head.next = this.head;
        this.head.prev = this.head;
    }

    length(){
        return this.head.data;
    }

    pushBack(el){
        let newNode = new Node(el);

        let currentNode = this.head;

        while(currentNode.next != this.head){
            let prevNode = currentNode;
            currentNode = currentNode.next;
            currentNode.prev = prevNode;    
        }

        currentNode.next = newNode;
        newNode.next = this.head;
        newNode.prev = currentNode;
        this.head.prev = newNode;
        this.head.data++;
    }

    pushFront(el) {
        let newNode = new Node(el);

        if(this.length() == 0){
            newNode.next = this.head
            newNode.prev = this.head;
            this.head.prev = newNode ; 
            this.head.next = newNode;
            this.head.data++;
            return;
        } 

        let firstNode = this.head.next;
        let currentNode = this.head;

        while(currentNode.next != this.head){
            let prevNode = currentNode;
            currentNode = currentNode.next;
            currentNode.prev = prevNode;
        }

        currentNode.next = this.head;
        this.head.prev = currentNode;

        newNode.next = firstNode;
        newNode.prev = this.head;
        firstNode.prev = newNode;
        this.head.next = newNode;
        this.head.data++;

    }

    pushIndex(el,index){
        if(index<1 || index>this.length()+1){
            console.log('参数错误');
            return;
        }
        if(index == 1){
            this.pushFront(el);
            return;
        }
        if(index==this.length()+1){
            this.pushBack(el);
            return;
        }

        let newNode = new Node(el);
        let prevNode = this.head;
        let cunrrentNode = null;

        for(let i=1; i<index; i++){
            let recordPrev = prevNode;
            prevNode = prevNode.next;
            prevNode.prev = recordPrev;
        }

        cunrrentNode = prevNode.next;
        cunrrentNode.prev = prevNode;

        this.head.prev = cunrrentNode;
        cunrrentNode.next = this.head;

        newNode.prev = prevNode;
        newNode.next = cunrrentNode;
        prevNode.next = newNode;
        cunrrentNode.prev = newNode;
        this.head.data++;

    }
}

module.exports = DoubleCircularList;