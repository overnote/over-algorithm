const DoubleCircularList = require("./doubleCircularList")

let dcl = new DoubleCircularList();
dcl.pushBack('a');
// dcl.pushBack('b');
dcl.pushFront('b');
console.log(dcl.length());
dcl.pushIndex('c',2)
console.log(dcl.head);
console.log(dcl.head.next);
console.log(dcl.head.next.next);
console.log(dcl.head.next.next.next);