## 笔记说明

本算法笔记采用`Go`语言、`JavaScript`语言同时描述，内部引用了一些专用函数，如Go的`append()`，这样做屏蔽了很多底层数据元素的操作细节，但是快速实现了数据结构，适合解题与实际应用，不再是学院派风格。  

源码中并未采取面向接口、继承等思想来实现工具方法的复用，笔者认为不应该将变成语言的学习带入数据结构，这样能够有效降低学习成本。笔者的意愿是：学习者拿到每个数据结构，都可以做到信手拈来，直接使用，而不是还要查询其继承结构。  

Go语法极简，读者即使不具备Go语言基础也能看懂代码，如果对Go、JavaScript语言有兴趣，也可以查看笔者的笔记进行详细学习：
- [Golang笔记](https://github.com/overnote/over-golang)   
- [JavaScript笔记](https://github.com/overnote/over-javascript) 

为了能够详细了解底层操作信息，避免Go和JavaScrit带来细节上缺失的遗憾，笔者在 [源码](https://github.com/overnote/over-algorithm/tree/master/sources) 目录也提供了 `C` 语言版本，严格按照学院派风格（如严蔚敏版）对数据结构、算法进行描述，适合学习考研。      

## 目录

- [第01章：概述](https://github.com/overnote/over-algorithm/tree/master/Overview/)
  - [数据结构概述]()
  - [算法概述]()
  - [大O表示法]()
- [第02章：数组](https://github.com/overnote/over-algorithm/tree/master/Array)
- [第03章：线性表](https://github.com/overnote/over-algorithm/tree/master/LinearList)
- [第04章：栈与队列](https://github.com/overnote/over-algorithm/tree/master/)
- [第05章：串与广义表](https://github.com/overnote/over-algorithm/tree/master/)
- [第06章：树](https://github.com/overnote/over-algorithm/tree/master/)
- [第07章：图](https://github.com/overnote/over-algorithm/tree/master/)
- [第08章：查找与排序](https://github.com/overnote/over-algorithm/tree/master/)
- [第09章：复杂数据结构](https://github.com/overnote/over-algorithm/tree/master/)
- [第10章：算法设计思想](https://github.com/overnote/over-algorithm/tree/master/)
- [第11章：面试题](https://github.com/overnote/over-algorithm/tree/master/)


## 数据结构与算法书籍

**入门**：
- [《大话数据结构》](https://book.douban.com/subject/6424904/)：适合0基础时了解数据结构，但是书中的小错误，啰嗦的地方较多
- [《漫画算法》](https://book.douban.com/subject/33420587/)：全彩配图，有趣的基础入门书
- [《算法图解》](https://book.douban.com/subject/26979890/)：适合算法基础薄弱时入门算法

**基础**
- [《算法》第4版](https://book.douban.com/subject/10432347/)：经典书籍，最好的算法书籍之一，Java编写
- [《数据结构》(邓俊辉)](https://book.douban.com/subject/25859528/)：数据结构的集大成者，深入浅出
- [《趣学算法》](https://book.douban.com/subject/27109832/)：覆盖了市面大部分算法

**提升**
- [《算法设计与分析基础》](https://book.douban.com/subject/26337727/)
- [《数据结构与算法分析》-C语言描述](https://book.douban.com/subject/4924153/)：数据结构与算法进阶的巨匠，也拥有Java版本：[传送门](https://book.douban.com/subject/26745780/)
- [《算法新解》](https://book.douban.com/subject/26931430/)：大量复杂数据结构的深入，难度较大

**思想**
- [《算法设计手册》](https://book.douban.com/subject/4048566/)
- [《算法技术手册(原书第2版)》](https://book.douban.com/subject/27123062/)
- [《算法设计》](https://book.douban.com/subject/2035809/)
- [《算法引论》](https://book.douban.com/subject/4178907/)

**成神**
- [《算法导论》](https://book.douban.com/subject/1885170/)：算法领域的代表作
- [《计算机程序设计艺术》](https://book.douban.com/subject/1130500/)：恢弘巨作，算法领域的里程碑

**刷题**
- [《程序员代码面试指南》](https://book.douban.com/subject/26638586/)：面试指南之一
- [《剑指offer》](https://book.douban.com/subject/27008702/)：面试指南之一
- [《编程珠玑》](https://book.douban.com/subject/3227098/)：为算法提供了精辟的解题思路，是算法思想学习的瑰宝
- [《编程之美》](https://book.douban.com/subject/3004255/)：微软面试指南集合

**一些网站**
- [力扣](https://leetcode.com/)：著名的算法题网站
- [牛客网](https://www.nowcoder.com/)：面向基础与面试的算法题库

## 附录：笔记汇总

**OverNote**全系列地址：https://github.com/overnote   

欢迎关注up主：https://github.com/ruyuejun

**OverNote分类**：  
- [Go笔记](https://github.com/overnote/over-golang)：详尽的Go领域笔记：Go语法、Go并发编程、GoWeb编程、Go微服务等
- [大前端](https://github.com/overnote/over-javascript)：包含JavaScript、Node.js、vue/react、微信开发、Flutter等大前端技术
- [数据结构与算法](https://github.com/overnote/over-algorithm)：以Go实现的数据结构与算法的笔记，附C，JavaScript版本
- [分布式与微服务架构](https://github.com/overnote/over-server)：分布式与微服务等架构笔记，附mysql、redis、nginx、docker、k8s等笔记
- [Linux](https://github.com/overnote/over-linux)：计算机组成原理、操作系统、计算机网络、编译原理基础学科笔记
- [大数据](https://github.com/overnote/over-bigdata)：大数据笔记，完善中
- [Python](https://github.com/overnote/over-python)：Python相关笔记，完善中
