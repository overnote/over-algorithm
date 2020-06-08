const LinkList = require("./linkList");

let ll = new LinkList();
ll.pushBack('e');
ll.pushFront('b');
console.log(ll.length());
ll.pushIndex('c',2);
console.log(ll.head.next);
console.log(ll.length());