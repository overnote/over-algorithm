function Node(element) {
    this.element = element;     //节点当前元素
    this.next = null;           //节点后继元素
    this.prev = null;           //节点前驱元素
}


function TWLinkedList() {
    this.head = new Node("head");   
}
TWLinkedList.prototype = {            
    constructor: LinkedList, 
    find(item) {               
        var currentNode = this.head;
        while(currentNode.element != item) {
            currentNode = currentNode.next;
        }
        return currentNode;
    },
    insert(newElement, item) {                 
        var currentNode = this.find(item);     
        var newNode = new Node(newElement);        
        newNode.prev = currentNode;                 //设置双向链表的前驱节点      
        newNode.next = currentNode.next;       
        currentNode.next = newNode;
    },
    remove(item) {                      //双向链表删除元素不再需要查找前驱节点
        var currentNode = this.find(item);
        if ( currentNode.next != null) {
            currentNode.prev.next = currentNode.next;
            currentNode.next.prev = currentNode.prev;
            currentNode.next = null;
            currentNode.prev = null;
        }
    },
    display() {           //链表正向输出元素方法                     
        var currentNode = this.head;
        while(currentNode.next != null) {
            console.log(currentNode.next.element);
            currentNode = currentNode.next;
        }
    },
    displayReverse() {      //双向链表反向输出元素方法
        var currentNode = this.head;
        while(currentNode.next != null) {
            currentNode = currentNode.next;
        }
        while(currentNode.prev != null) {
            console.log(currentNode.element);
            currentNode = currentNode.prev;
        }
    }

}

// 使用案例：

var cities = new TWLinkedList();
cities.insert("BeiJing", "head");
cities.insert("ShangHai", "BeiJing");
cities.insert("GuangZhou","ShangHai");
console.log(cities.displayReverse());  //输出 GuangZhou ShangHai  BeiJing undefined
console.log("----------"); 
cities.remove("ShangHai");
console.log(cities.displayReverse()); //输出 GuangZhou BeiJing undefined
