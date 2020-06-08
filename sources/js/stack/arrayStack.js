/*
 * @Name: 栈
 * @Description: 数组实现
 * @Author: yaya
 * @Date: 2020-01-15 11:25:07
 * @Editors: yaya
 */

class ArrayStack{
    constructor(){
        this.data = [];
    }

    length(){
        return this.data.length;
    }

    // 压栈
    push(el){
        this.data.push(el);
    }

    // 弹栈
    pop(){

        if(this.length() == 0){
            return null;
        }

        this.data.pop();
    }

    // 获取栈顶元素
    top(){
        if(this.length() == 0){
            return null;
        }
        return this.data[this.length()-1];
    }
}

module.exports = ArrayStack