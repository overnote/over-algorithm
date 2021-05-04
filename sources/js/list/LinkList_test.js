const Linklist = require('./LinkList')

const l = new Linklist()
l.display()

l.insert(11, 1)
l.insert(12, 2)
l.insert(13, 3)
l.insert(14, 4)
l.insert(15, 5)
l.display()

console.log(l.delete(5))
l.display()
console.log(l.delete(3))
l.display()

l.update(3, 33)
l.display()

l.clear()
l.display()
