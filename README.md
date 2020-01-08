## 笔记说明

本算法笔记采用`Go`语言描述，同时内部引用了Go的一些专有函数，如`append()`，这样做屏蔽了很多底层数据元素的操作细节，但是快速实现了数据结构，适合解题与实际应用，不是学院派风格。  

源码中并未采取面向接口、继承等思想来实现工具方法的复用，尤其是在二叉搜索树、红黑树等章节，每个章节都具备自己独立的代码。因为笔者认为这些笔记是用来学习数据结构的，而不是用来学习面向对象思想的，过渡抽象会给初学者造成困惑。笔者的意愿是：学习者拿到每个数据结构，都可以做到信手拈来，直接使用，而不是还要查询其继承结构。  

笔者使用Go语言记录数据结构与算法笔记的原因：
- C：非面向对象，笔者认为不适合快速学习，会耗费一定的时间在内存管理之上
- C++：拥有面向对象功能，但是同样需要进行内存管理，会耗费一定的时间在此之上
- js/py：脚本语言写法变换太多，且很多结构在底层与大多语言结构特性不符，比如js的数组其实是哈希表
- java：具备面向对象，且语法严谨规范，适合学习与教学，但是市面上Java相关的数据结构教程层出不穷，已经做得很好了
- go：与Java一样严谨，完备，且语法上更加简洁，市面上很好的go算法书籍、教程极少

Go语法极简，读者即使不具备Go语言基础也能看懂代码，如果对Go语言有兴趣，也可以查看笔者的[Golang笔记](https://github.com/overnote/over-golang)进行详细学习。   

当然为了能够详细了解底层操作信息，笔者在sources文件夹也提供了`C`语言版本，严格按照学院派风格（如严蔚敏版）对数据结构、算法进行描述，适合学习考研。     
同时笔记也将提供`JavaScript`作为动态语言版本的示例学习数据结构与算法，未来可能考虑引入C++、Python、Java版本。    

**Go**版本数据结构与算法以实际应用为准则，基于go mod，环境，笔者建议直接安装go1.13以上版本即可，会默认开启 go mod。食用方式：
```
cd sources/go
go run main.go                  // 在 main.go 中打开想要的测试方法即可
```

**C**版本数据结构与算法以标准C为主，适合考研党，食用方式：
```
cd sources/c
cc main.c                       // 在 main.c 中打开想要的测试方法即可
```

**JavaScript**版数据结构与算法实现完全采用ES6、ES Module方式书写，需要编译支持。笔者已经配置好了编译模块，按照下列命令直接食用即可： 
```
# 在sources/js 目录中  贴士：需要提前安装node.js。
npm install
npx babel-node index.js         // 在 index.js 中打开想要的测试方法即可
```

## 数据结构与算法书籍

### 1.1 入门书籍

- [《大话数据结构》](https://book.douban.com/subject/6424904/)：适合0基础时了解数据结构，但是书中的小错误，啰嗦的地方，也有点让人发指
- [《漫画算法》](https://book.douban.com/subject/33420587/)：全彩配图，有趣的基础入门书
- [《算法图解》](https://book.douban.com/subject/26979890/)：适合算法基础薄弱时入门算法

### 1.2 基础书籍

- [《算法》第4版](https://book.douban.com/subject/10432347/)：经典书籍，最好的算法书籍之一，Java编写
- [《数据结构》(邓俊辉)](https://book.douban.com/subject/25859528/)：数据结构的集大成者，深入浅出
- [《趣学算法》](https://book.douban.com/subject/27109832/)：覆盖了市面大部分算法

### 1.3 算法加强

- [《算法设计与分析基础》](https://book.douban.com/subject/26337727/)
- [《数据结构与算法分析》-C语言描述](https://book.douban.com/subject/4924153/)：数据结构与算法进阶的巨匠，也拥有Java版本：[传送门](https://book.douban.com/subject/26745780/)
- [《算法新解》](https://book.douban.com/subject/26931430/)：大量复杂数据结构的深入，难度较大

### 1.4 算法进阶

- [《算法导论》](https://book.douban.com/subject/1885170/)：算法领域的代表作
- [《计算机程序设计艺术》](https://book.douban.com/subject/1130500/)：恢弘巨作，算法领域的里程碑

### 1.5 算法题

- [《程序员代码面试指南》](https://book.douban.com/subject/26638586/)：面试指南之一
- [《剑指offer》](https://book.douban.com/subject/27008702/)：面试指南之一
- [《编程珠玑》](https://book.douban.com/subject/3227098/)：为算法提供了精辟的解题思路，是算法思想学习的瑰宝
- [《编程之美》](https://book.douban.com/subject/3004255/)：微软面试指南集合

### 1.6 算法设计

- [《算法设计手册》](https://book.douban.com/subject/4048566/)
- [《算法技术手册(原书第2版)》](https://book.douban.com/subject/27123062/)
- [《算法设计》](https://book.douban.com/subject/2035809/)
- [《算法引论》](https://book.douban.com/subject/4178907/)

## 算法资料

- [力扣](https://leetcode.com/)：著名的算法题网站
- [牛客网](https://www.nowcoder.com/)：面向基础与面试的算法题库

## 附录：笔记汇总

**OverNote**全系列地址：https://github.com/overnote   

欢迎关注up主：https://github.com/ruyuejun

**OverNote分类**：  
- [Go](https://github.com/overnote/over-golang)：详尽的Go领域笔记：Go语法、Go并发编程、GoWeb编程、Go微服务等
- [大前端](https://github.com/overnote/over-javascript)：包含JavaScript、Node.js、vue/react、微信开发、Flutter等大前端技术
- [数据结构与算法](https://github.com/overnote/over-algorithm)：以Go实现的数据结构与算法的笔记，附C，JavaScript版本
- [分布式与微服务架构](https://github.com/overnote/over-server)：分布式与微服务等架构笔记，附mysql、redis、nginx、docker、k8s等笔记
- [Linux](https://github.com/overnote/over-linux)：计算机组成原理、操作系统、计算机网络、编译原理基础学科笔记
- [大数据](https://github.com/overnote/over-bigdata)：大数据笔记，完善中
- [Python](https://github.com/overnote/over-python)：Python相关笔记，完善中
