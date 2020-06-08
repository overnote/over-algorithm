/*
 * @Name: 单向循环链表
 * @Description: 带头结点的链表
 * 单向链表、单向循环链表、双向链表、双向循环链表
 * @Author: yaya
 * @Date: 2020-01-13 13:37:52
 * @Editors:yaya
 */


class Node{
    constructor(data){
        this.data = data ;
        this.next = null;
    }
}

class CircularList{
    constructor(){
        this.head = new Node(0);
        this.head.next = this.head;
    }

    length(){
        return this.head.data;
    }

    pushBack(el){

        let newNode = new Node(el);
        let currentNode = this.head;

        while(currentNode.next != this.head ){
            currentNode = currentNode.next
        }

        newNode.next = this.head;
        currentNode.next = newNode;
        this.head.data++;

    }

    pushFront(el){
        let newNode = new Node(el);

        if(this.head.next == this.head){
            this.head.next = newNode;
            newNode.next = this.head;
            this.head.data++;
            return;
        }

        let firstNode = this.head.next;

        while(firstNode.next != this.head){
            firstNode = firstNode.next;
        }

        newNode.next = firstNode;
        this.head.next = newNode;
        this.head.data++;

    }

    pushIndex(el,index){
        if(index<1 || index > this.length()+1){
            console.log('参数错误');
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

        let newNode = new Node(el);
        let preNode = this.head;
        let currentNode = null;

        for(let i=1;i<index; i++){
            preNode = preNode.next;
        }

        currentNode = preNode.next;
        newNode.next = currentNode;
        preNode.next = newNode;
        this.head.data++;

    }
}

module.exports = CircularList;