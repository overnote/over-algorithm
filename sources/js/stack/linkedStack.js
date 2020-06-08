/*
 * @Name: 栈
 * @Description: 链表实现栈
 * @Author: yaya
 * @Date: 2020-01-15 11:36:26
 * @Editors: yaya
 */

class Node{
    constructor(data){
        this.data = data;
        this.next = null;
    }
}

class LinkedStack{
    constructor(){
        this.node = new Node(null);
        this.top = this.node;
        this.length = 0;
    }
    
    // 压栈
    push(el){
        let newNode = new Node(el);

        let currentNode = this.node;

        newNode.next = currentNode;
        this.node = newNode;
        this.top = newNode;
        this.length++;
    }

    // 出栈弹栈
    pop(){
        if(this.length == 0){
            return null;
        }

        // 取头部
        let firstNode  = this.node;

        if(this.length == 1){
            this.node = null;
            this.top = null;
            this.length--;
            return firstNode.data;
        }

        let secondNode = this.node.next;

        this.node = secondNode;
        this.top = secondNode;
        this.length --;

        return firstNode.data;

    }

    // 获取栈顶元素
    topEl(){
        if(this.length == 0){
            console.log('栈为空栈');
            return ;
        }
        
        return this.top.data; 
    }
}

module.exports = LinkedStack