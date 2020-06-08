const LinkedStack = require("./linkedStack")

let ls = new LinkedStack();

ls.push('a');
ls.push('b');
console.log(ls);
console.log(ls.topEl());
console.log(ls.pop());
console.log(ls);