# README

## 笔记说明

本算法笔记采用 `C` 语言进行严谨描述，并提供`Go` 、`JavaScript`等符合实际的快捷简便版本，这是为了符合考研、工作双重需要，在基础原理部分使用严谨的学院派风格的`C` 语言描述，再附带上工作中常用的 `Go` 语言、`JavaScript` 等语言来进行简化（未来会补全`C++` `Java`、`Python` 版本）。

贴士：

- code 文件夹中，`code-c` 为纯 C 实现，文件后缀为 `.c`
- `code-c-c++` 则是使用 C++环境（以 `.cpp` 为文件后缀），以 C 语言风格书写，这是为了额外对应国内一些常见教材中的写法，非作者本意
- `code-c++` 文件夹则是纯正的 C++语法书写

源码中并未采取面向接口、继承等思想来实现工具方法的复用，笔者认为不应该将编程语言的学习成本带入数据结构，笔者的意愿是：学习者拿到每个数据结构，都可以做到信手使用，而不是还要反复查询其继承、实现结构。

[点击查看源码](https://github.com/overnote/over-algorithm/tree/master/sources)

学习的推荐：

- 如果完全零基础，笔者建议通过视频来学习，B 站上拥有大量零基础数据结构教学视频，搜索即可。
- 拥有一点基础后，笔者推荐考研使用《算法 4》，后续通过 LeetCode、《剑指 Offer》大量做题就 OK 了。
- 切记笔者总结编程学习经验：**阅万卷，不若作一页之书；操千道，求解方不由分说**。

## 数据结构与算法书籍推荐

### 1.1 基础书籍

- [](https://book.douban.com/subject/27109832/)：入门书籍
- [《算法》第 4 版](https://book.douban.com/subject/19952400/)：经典书籍，最好的算法书籍之一，Java 编写
- [《算法详解》](https://book.douban.com/subject/30424415/)：共 2 卷，弥补《算法 4》分治算法缺失，[点击进入卷 2 地址](https://book.douban.com/subject/35093209/)
- [《数据结构与算法分析》-C 语言描述](https://book.douban.com/subject/4924153/)：数据结构与算法进阶的巨匠，也拥有 Java 版本：[传送门](https://book.douban.com/subject/26745780/)
- [《算法新解》](https://book.douban.com/subject/26931430/)：大量复杂数据结构的深入，难度较大

其他书籍：

```txt
《趣学数据结构》：适合零基础入门
《大话数据结构》：适合零基础入门
《趣学算法》：适合零基础入门
《算法图解》：适合零基础入门
《漫画算法：小灰的算法之旅》：适合零基础入门
《啊哈!算法》：适合零基础入门
《我的第一本算法书》：适合零基础入门
《算法精解 C 语言描述》：源码完整，适合参考
```

### 1.2 算法设计书籍

- [《算法设计》Jon Kleinberg](https://book.douban.com/subject/35391618/)
- [《算法引论 一种创造性方法》](https://book.douban.com/subject/4178907/)
- [《算法设计指南》斯基恩纳](https://book.douban.com/subject/27092363/)：共 2 卷，原书位于：[《算法设计手册》](https://book.douban.com/subject/4048566/)
- [《算法之道(第 2 版)》](https://book.douban.com/subject/10564644/)
- [《算法设计与分析基础》](https://book.douban.com/subject/26337727/)

### 1.3 刷题书籍

- [《程序员面试金典》（第 6 版）](https://book.douban.com/subject/34813624/)：面试指南之一
- [《程序员代码面试指南》](https://book.douban.com/subject/26638586/)：面试指南之一
- [《剑指 offer》](https://book.douban.com/subject/27008702/)：面试指南之一
- [《编程珠玑》](https://book.douban.com/subject/3227098/)：为算法提供了精辟的解题思路，是算法思想学习的瑰宝
- [《编程之美》](https://book.douban.com/subject/3004255/)：微软面试指南集合
- [《算法笔记》](https://book.douban.com/subject/26827295/)：考研类刷题书籍

### 1.4 成神书籍

- [《算法心得：高效算法的奥秘（原书第 2 版）》](https://book.douban.com/subject/25837031/)
- [《算法导论》](https://book.douban.com/subject/1885170/)：算法领域的代表作
- [《计算机程序设计艺术》](https://book.douban.com/subject/1130500/)：恢弘巨作，算法领域的里程碑

### 1.5 竞赛类书籍

- [《算法竞赛入门经典（第 2 版）》](https://book.douban.com/subject/25902102/)
- [《挑战程序设计竞赛》](https://book.douban.com/subject/24749842/)
- [《算法竞赛入门经典》](https://book.douban.com/subject/20254543)
- [《算法竞赛进阶指南》](https://book.douban.com/subject/30136932/)
- [《ACM 国际大学生程序设计竞赛》](https://book.douban.com/subject/20505391/)

## 附录：笔记汇总

**OverNote**全系列地址：<https://github.com/overnote>

欢迎关注 up 主：<https://github.com/ruyuejun>

**OverNote 分类**：

- [Go 笔记](https://github.com/overnote/over-golang)：详尽的 Go 领域笔记：Go 语法、Go 并发编程、GoWeb 编程、Go 微服务等
- [大前端](https://github.com/overnote/over-javascript)：包含 JavaScript、Node.js、vue/react、微信开发、Flutter 等大前端技术
- [数据结构与算法](https://github.com/overnote/over-algorithm)：以 Go 实现的数据结构与算法的笔记，附 C，JavaScript 版本
- [分布式与微服务架构](https://github.com/overnote/over-server)：分布式与微服务等架构笔记，附 mysql、redis、nginx、docker、k8s 等笔记
- [cs](https://github.com/overnote/over-cs)：计算机组成原理、操作系统、计算机网络、编译原理基础学科笔记
- [大数据](https://github.com/overnote/over-bigdata)：大数据笔记，完善中
- [Python](https://github.com/overnote/over-python)：Python 相关笔记，完善中
