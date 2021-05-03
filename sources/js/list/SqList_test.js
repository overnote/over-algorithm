const SqList = require('./SqList')

const l = new SqList()
l.display()

l.insert(11, 1)
l.insert(12, 2)
l.insert(13, 3)
l.insert(14, 4)
l.insert(15, 5)
l.insert(17, 7)
l.insert(15, 5)
l.display()

console.log('删除元素:', l.delete(3))
l.display()

l.update(1, 100)
l.display()

console.log('查询元素:', l.search(2))

console.log('查询元素:', l.locate(100))
