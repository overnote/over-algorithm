## 笔记说明

本算法笔记采用`C`、`Go`、`JavaScript`三种语言同时解题，既满足了大部分算法书籍、大学课程以C为基础的需要，也能满足大家学习利用新兴语言Go等解决算法题的期望。  

Go语法极简，读者即使不具备Go语言基础也能看懂代码，如果对Go语言有兴趣，也可以查看笔者的[Golang笔记](https://github.com/overnote/over-golang)进行详细学习。    

笔记的一些贴士：
- 在理论描述上，均使用基于C语言的最基础代码展示，读者不必担忧因为对C的不了解产生阅读障碍
- 在代码实现上，笔记中所有的数据结构与算法均使用考研标准C、Go、JS三种语言同时描述，并在source目录中上传了全部源码

**C**版本数据结构与算法以标准C为主，适合考研党，食用方式：
```
cd sources/c
cc main.c                       // 在 main.c 中打开想要的测试方法即可
```

**Go**版本数据结构与算法以实际应用为准则，基于go mod，环境，笔者建议直接安装go1.13以上版本即可，会默认开启 go mod。食用方式：
```
cd sources/go
go run main.go                  // 在 main.go 中打开想要的测试方法即可
```

**JavaScript**版数据结构与算法实现完全采用ES6、ES Module方式书写，需要编译支持。笔者已经配置好了编译模块，按照下列命令直接食用即可： 
```
# 在sources/js 目录中  贴士：需要提前安装node.js。
npm install
npx babel-node index.js         // 在 index.js 中打开想要的测试方法即可
```

## 数据结构与算法书籍

### 1.1 入门书籍

- [《大话数据结构》](https://book.douban.com/subject/6424904/)：适合0基础时了解数据结构
- [《漫画算法》](https://book.douban.com/subject/33420587/)：全彩配图，有趣的基础入门书
- [《算法图解》](https://book.douban.com/subject/26979890/)：适合算法基础薄弱时入门算法

### 1.2 基础书籍

- [《算法》第4版](https://book.douban.com/subject/10432347/)：经典书籍，最好的算法书籍之一，Java编写
- [《趣学算法》](https://book.douban.com/subject/27109832/)：覆盖了市面大部分算法

### 1.3 算法加强

- [《算法设计与分析基础》](https://book.douban.com/subject/26337727/)
- [《数据结构》(邓俊辉)](https://book.douban.com/subject/25859528/)：数据结构的集大成者，深入浅出
- [《数据结构与算法分析》-C语言描述](https://book.douban.com/subject/4924153/)：数据结构与算法进阶的巨匠，也拥有Java版本：[传送门](https://book.douban.com/subject/26745780/)

### 1.4 算法进阶

- [《算法新解》](https://book.douban.com/subject/26931430/)：大量复杂数据结构的深入，难度较大
- [《算法导论》](https://book.douban.com/subject/1885170/)：算法领域的代表作
- [《计算机程序设计艺术》](https://book.douban.com/subject/1130500/)：恢弘巨作，算法领域的里程碑

### 1.5 面试方向

- [《编程珠玑》](https://book.douban.com/subject/3227098/)：为算法提供了精辟的解题思路，是算法思想学习的瑰宝
- [《编程之美》](https://book.douban.com/subject/3004255/)：微软面试指南集合
- [《程序员代码面试指南》](https://book.douban.com/subject/26638586/)：面试指南之一
- [《剑指offer》](https://book.douban.com/subject/27008702/)：面试指南之一

## 算法资料

- [力扣](https://leetcode.com/)：著名的算法题网站
- [牛客网](https://www.nowcoder.com/)：面向基础与面试的算法题库
- [浙江大学-陈越姥姥-数据结构](https://www.icourse163.org/course/ZJU-93001)：经典数据结构视频
- [清华大学-邓俊辉-数据结构与算法](https://www.bilibili.com/video/av49361421)：经典数据结构与算法视频

## 附录：笔记汇总

**OverNote**全系列地址：https://github.com/overnote   

欢迎关注up主：https://github.com/ruyuejun

**OverNote分类**：  
- [Go](https://github.com/overnote/over-golang)：详尽的Go领域笔记：Go语法、Go并发编程、GoWeb编程、Go微服务等
- [大前端](https://github.com/overnote/over-front-end)：包含JavaScript、Node.js、vue/react、微信开发、Flutter等大前端技术
- [数据结构与算法](https://github.com/overnote/over-algorithm)：以C/Go实现为主记录数据结构与算法的笔记
- [分布式与微服务架构](https://github.com/overnote/over-architecture/)：分布式与微服务等架构相关笔记
- [Linux](https://github.com/overnote/over-linux)：计算机组成原理、操作系统、计算机网络、编译原理基础学科笔记
- [服务端常用技术](https://github.com/overnote/over-server)：nginx、mysql、redis、mongodb、linux系统基础等服务端常用技术汇总笔记
- [大数据](https://github.com/overnote/over-bigdata)：大数据笔记，完善中
- [Python](https://github.com/overnote/over-python)：Python相关笔记，完善中
