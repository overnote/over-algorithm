//先设计一个Node类表示一个一个的节点
function Node(element) {
    this.element = element;     //节点当前元素
    this.next = null;           //节点指向的元素
}

//LinkedList类：提供插入节点、删除节点、显示元素列表等方法
function LinkedList() {
    this.head = new Node("head");   //设置头节点，新建的头节点指向null
}
LinkedList.prototype = {            
    constructor: LinkedList,    //注意：本章所有的原型上的函数直接写在原型内，此时需要修改构造器指向
    find(item) {                //传入一个元素，查找他的下一个元素，这也是链表的训话迭代方式
        var currentNode = this.head;
        while(currentNode.element != item) {
            currentNode = currentNode.next;
        }
        return currentNode;
    },
    insert(newElement, item) {                  //newElement插入到item后
        var currentNode = this.find(item);      //查找传入元素所在的节点
        var newNode = new Node(newElement);     //创建新插入的节点
        newNode.next = currentNode.next;        //注意该句和下一句顺序不能相反
        currentNode.next = newNode;
    },
    remove(item) {
        var prevNode = this.head; //找到要删除的item的上一个节点
        while(prevNode != null && prevNode.next.element != item) {
            prevNode = prevNode.next;
        }
        if (prevNode.next != null) {
            prevNode.next = prevNode.next.next;
        }
    },
    display() {                                 //获取链表中的元素
        var currentNode = this.head;
        while(currentNode.next != null) {
            console.log(currentNode.next.element);
            currentNode = currentNode.next;
        }
    }

}


var cities = new LinkedList();
cities.insert("BeiJing", "head");
cities.insert("ShangHai", "BeiJing");
cities.insert("GuangZhou","ShangHai");
console.log(cities.display()); //BeiJing ShangHai GuangZhou undefined
console.log("----------"); 
cities.remove("ShangHai");
console.log(cities.display());//BeiJing GuangZhou undefined