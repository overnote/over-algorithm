const ArrayStack = require("./arrayStack")

let s = new ArrayStack();
s.push('a');
s.push('b');
s.push('c');
console.log(s);
s.pop();
console.log(s);
console.log(s.top());