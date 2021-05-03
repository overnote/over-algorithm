const DynamicArray = require('./DynamicArray')

const arr = new DynamicArray()
arr.display()

// 测试插入
arr.insert(10, 0)
arr.insert(11, 1)
arr.insert(12, 2)
arr.insert(13, 3)
arr.insert(14, 4)
arr.display()

// 测试扩容
arr.insert(15, 5)
arr.insert(16, 6)
arr.display()
console.log('容量：', arr.cap())

// 测试删除
arr.delete(0)
arr.display()
arr.delete(2)
arr.display()

// 测试缩容
console.log('容量：', arr.cap())
arr.delete(4)
arr.delete(0)
arr.delete(0)
arr.delete(0)
arr.display()
console.log('容量：', arr.cap())

// 测试更新
arr.update(100, 0)

// 测试查找
console.log('位置：', arr.locate(100))

// 测试取值
console.log('取值：', arr.get(0))
