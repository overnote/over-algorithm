/**
 * 可变数组
 */

const MaxSize = 5 // 默认容量，取小值便于测试

class DynamicArray {
    constructor() {
        this.data = new Array(5) // 可变数组内的真实数组
        this.capacity = MaxSize // 可变数组容量
        this.length = 0 // 可变数组元素个数
    }
    // 增：根据索引插入数据
    insert(e, index) {
        if (index < 0 || index > this.length) {
            console.log('插入位置不合法')
            return
        }
        // 每次插入前执行扩容机制
        this.expandCap()
        // 执行插入
        for (let i = this.length; i >= index; i--) {
            if (i == index) {
                this.data[index] = e
            } else {
                this.data[this.length] = this.data[this.length - 1]
            }
        }
        // 更新长度
        this.length++
    }
    // 删：按索引位置删除元素
    delete(index) {
        if (index < 0 || index > this.length - 1) {
            console.log('删除位置不合法')
            return
        }
        // 每次删除钱执行缩容
        this.reduceCap()
        // 执行删除
        for (let i = 0; i < this.length; i++) {
            this.data[i] = this.data[i + 1]
        }
        this.length--
    }
    // 改：根据索引修改元素
    update(e, index) {
        if (index < 0 || index > this.length - 1) {
            console.log('数组索引越界')
            return
        }
        this.data[index] = e
    }
    // 查：根据值查询索引
    locate(e) {
        for (let i = 0; i < this.length; i++) {
            if (this.data[i] == e) {
                return i
            }
        }
        return -1
    }
    // 取：根据索引值进行数据存取
    get(index) {
        if (index < 0 || index > this.length - 1) {
            console.log('数组索引越界')
            return null
        }
        return this.data[index]
    }
    // 获取容量
    cap() {
        return this.capacity
    }
    // 获取长度
    len() {
        return this.length
    }
    // 清空
    clear() {
        this.length = 0
    }
    // 销毁：自动内存管理，无需该方法
    // 扩容方法：超过数组容量，按照当前容量的 2 倍扩容
    expandCap() {
        if (this.length < this.capacity) {
            return
        }
        // 拷贝数据
        let newCap = this.capacity * 2
        let newData = [] // 只是模拟，其实完全没这个必要
        for (let i = 0; i < this.length; i++) {
            newData[i] = this.data[i]
        }
        this.data = newData
        // 变更容量
        this.capacity = newCap
    }
    // 缩容方法：数组元素为当前容量的 1/4，缩容为当前容量的一半
    reduceCap() {
        if (this.length > this.capacity / 4) {
            return
        }
        // 拷贝数据
        let newCap = this.capacity / 2
        let newData = [] // 只是模拟，其实完全没这个必要
        for (let i = 0; i < this.length; i++) {
            newData[i] = this.data[i]
        }
        this.data = newData
        // 变更容量
        this.capacity = newCap
    }
    // 显示可变数组数据
    display() {
        console.log(this.data.slice(0, this.length))
    }
}

module.exports = DynamicArray
