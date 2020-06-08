const CircularList = require("./circularList")

let cl = new CircularList();

cl.pushBack('a');
cl.pushFront('b');
console.log(cl.length());
cl.pushIndex('c',2)
console.log(cl.head);
console.log(cl.head.next);
console.log(cl.head.next.next);
console.log(cl.head.next.next.next);
console.log(cl.head.next.next.next.next);
// console.log(cl);