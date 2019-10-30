/**
 * 顺序表
 */
class SqList {
    constructor(){              // 构造函数 相当于C中的 init，Go找那个的new
        this.length = 0;
        this.elem = [];
    }
    len(){                   // 获取长度
        return this.length;
    }
    display(){                  // 显示所有元素
        if(this.length == 0) {
            console.log("为空表");
        } else {
            for (let i = 0 ; i < this.length; i++) {
                console.log("元素[" + (i+1) + "] = " + this.elem[i]);
            }
        }
    }
    insert(index, e){        // 插入元素    
        // 如果是空表
        if (this.length == 0 && index == 1) {
            this.elem[0] = e;
            this.length++;
            return
        }
        // 如果索引越界
        if (index < 1 || index > this.length+1) {
            console.log("索引越界");
            return
        }
        // 执行插入 从最后一位开始一次后移
        for (let i = this.length + 1; i > index; i--) {
            this.elem[i - 1] = this.elem[i - 2];
        }
        this.elem[index - 1] = e;
        this.length++;
    }
    delete(index){          // 删除元素
        if(index < 1 || index > this.length) {
            console.log("索引越界");
            return
        }
        for (let i = index - 1; i < this.length - 1; i++) {
            this.elem[i] = this.elem[ i + 1 ];
        }
        this.length--
    }
    update(index, e){       // 修改元素
        if (index < 1 || index > this.length) {
            console.log("位序越界");
            return
        }
        this.elem[index - 1] = e;
    }        
    getElem(index) {        // 根据位序查询
        if (index < 1 || index > this.length) {
            console.log("位序不合法");
            return
        }
        return this.elem[index - 1]
    }
    locateElem(e) {     // 根据值查询
        let index = null;
        for (let i = 0; i < this.length; i++){
            if (this.elem[i] == e) {
                index = i + 1;
                break;
            }
        }
        return index;
    }  
    priorElem(e){       // 查询前驱
        if(this.length <= 1) {
            console.log("顺序表元素不足，无法查询");
            return
        }
        if (this.elem[0] == e) {
            console.log("第一个元素无前驱");
            return
        }
        let index = this.locateElem(e);
        if (!index) {
            console.log("未找到");
            return
        }
        return this.elem[index - 2]
    }
    nextElem(e){        // 查询后继
        if (this.length <= 1) {
            console.log("数据结构为空，无法查询");
            return
        }
        if (this.elem[this.length - 1] == e) {
            console.log("最后一个元素无后继");
            return
        }
        let index = this.locateElem(e)
        if (!index) {
            console.log("未找到");
            return
        }
        return this.elem[index]
    } 
    clear() {       // 清空
        this.length = 0;
        this.elem  = [];
    }
}

export default SqList