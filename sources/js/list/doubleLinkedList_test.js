const DoubleLinkList = require("./doubleLinkedList")

let cl = new DoubleLinkList();
cl.pushBack('a');
cl.pushFront('b');
console.log(cl.length());
cl.pushIndex('c',2)
console.log(cl.head);
console.log(cl.head.next);
console.log(cl.head.next.next);
console.log(cl.head.next.next.next);