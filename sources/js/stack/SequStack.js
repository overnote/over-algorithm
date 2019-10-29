// ES6方式

let StackES6 = (function () {
    const store = new WeakMap();
    class Stack {
        constructor(){
            store.set(this, [])
        }
        push(e){
            store.get(this).push(e);
        }
        pop(){
            return store.get(this).pop();
        }
        size(){
            return store.get(this).length;
        }
        //其他方法
    }
    return Stack;
})();

// ES5方式

function StackES5() {
    this.store = [];        //栈内部使用一个数据结构来保存栈里的元素，这里使用数组：
}

//增:直接使用数组的push方法保存
StackES5.prototype.push = function(e) {
    this.store.push(e);

}
//删:只能删除栈顶元素，即最后加入数组中的元素 pop方法正合适
StackES5.prototype.pop = function() {
    return this.store.pop();
}

/* 通过上述的增删，已经实现了Stack对象的 LIFO原则 */

//获取栈长度
StackES5.prototype.size = function() {
    return this.store.length;
}
//获取栈顶元素
StackES5.prototype.peek = function() {
    return this.store[this.store.length - 1]
}
//清空栈
StackES5.prototype.clear = function() {
    arr.splice(0, arr.length);
}