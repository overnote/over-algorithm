## 笔记说明

本算法笔记采用`C`、`Go`两种语言解题，既满足了大部分算法书籍、大学课程以C为基础的需要，也能满足大家学习利用新兴语言Go解决算法题的期望。  

Go语法极简，读者即使不具备Go语言基础也能看懂代码，如果对Go语言有兴趣，也可以查看笔者的[Golang笔记](https://github.com/overnote/over-golang)进行详细学习。    

对应的数据结构与算法源码也提供了部分js、java等其他语言版本，文件夹位于：[source](https://github.com/overnote/five-x/tree/master/sources)，有兴趣可以提交各种语言版本噢。   

笔记的一些贴士：
- `严本` 即 严蔚敏版本的意思

## C语言的一些规范

### 规范一：结点的表示方式
推荐链表结点的表示方式：
```c
typedef struct Node {
    int data;
    struct Node *next;
}Node;
```

推荐树结点的表示方式：
```c
typedef struct BTNode {
    int data;
    struct BTNode *lchild;
    struct BTNode *rchild;
}BTNode
```
上述结构体写法更加严谨，能够通过纯C编译器的编译（适合考研党）！ 

### 规范二：结点的申请方式
结点的申请有两种方式：
```c
// 方式一：不推荐
BTNode BT;      // 此时取值方式：x=BT.data

// 方式二：推荐。先定义一个节点指针BT，然后让BT指向这片内存空间。
BTNode *BT;
BT = （BTNode *）malloc(sizeof(BTNode));
```
方式一的BT只是某个节点的名字，定义好后无法脱离该节点，而方式二只是一个指针，可以随时指向其他结点。  

动态申请一组结点，如动态申请数组空间，推荐写法：
```c
int *arr;
arr = (int *)malloc(n * sizeof(int));       // 元素类型为int，长度为n的动态数组
```

### 规范三：函数的引用传参

值类型作为函数形参时，如果希望修改原本的值，则使用以下方式：
```c
void fn(int *&x) {

}
```

一维数组作为参数写法：
```c
//  参数一无需限定长度，参数二是个习惯，用来说明数组实际元素的个数，并不是总长度
void fn(int arr[], int n){
    
}
```

二维数组作为参数写法：
```c
// 第一个参数中的第一个[]无需写长度，第二个[]必须写数组长度的常量
void fn(int arrp[][maxSize], int n){

}
```


## 数据结构与算法书籍

### 1.1 入门书籍

- [《大话数据结构》](https://book.douban.com/subject/6424904/)：适合完全0基础读者了解数据结构，了解即可，因为内部有些许错误
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

- [《算法新解》](https://book.douban.com/subject/26931430/)
- [《算法导论》](https://book.douban.com/subject/1885170/)：恢弘巨著，算法领域的代表作
- [《计算机程序设计艺术》](https://book.douban.com/subject/1130500/)：恢弘巨作，算法领域的里程碑

### 1.5 算法习题

面试方向：
- [《编程珠玑》](https://book.douban.com/subject/3227098/)：为算法提供了精辟的解题思路，是算法思想学习的瑰宝
- [《编程之美》](https://book.douban.com/subject/3004255/)：微软面试指南集合
- [《程序员代码面试指南》](https://book.douban.com/subject/26638586/)：面试指南之一
- [《剑指offer》](https://book.douban.com/subject/27008702/)：面试指南之一

竞赛方向：
- [《算法竞赛进阶指南》李煜东](https://book.douban.com/subject/30136932/)：适合竞赛训练
- [《算法竞赛入门经典》刘汝佳·第2版](https://book.douban.com/subject/25902102/)：适合竞赛训练

考研方向：
- [《数据结构》严蔚敏](https://book.douban.com/subject/24699581/)：绝大多数的指定教材，结构严谨，但是没有代码实现
- [《数据结构算法解析（第2版）》](https://book.douban.com/subject/26823999/)：严版数据结构的实现
- [《算法笔记》](https://book.douban.com/subject/26827295/)：适合考研复习、PAT考试
- [《算法笔记上机训练实战指南》](https://book.douban.com/subject/30162908/)：算法笔记配套练习册

## 算法资料

网站：
- [力扣](https://leetcode.com/)：著名的算法题网站
- [牛客网](https://www.nowcoder.com/)：面向基础与面试的算法题库

视频：
- [浙江大学-陈越姥姥-数据结构](https://www.icourse163.org/course/ZJU-93001)
- [清华大学-邓俊辉-数据结构与算法](https://www.bilibili.com/video/av49361421)

## 附录：笔记汇总

**OverNote**地址：https://github.com/overnote   

**推荐书籍**地址：https://github.com/ruyuejun/polaris  

**OverNote分类**：  
- [Go](https://github.com/overnote/over-golang)：详尽的Go领域笔记：Go语法、Go并发编程、GoWeb编程、Go微服务等
- [大前端](https://github.com/overnote/over-front-end)：包含JavaScript、Node.js、vue/react、微信开发、Flutter等大前端技术
- [数据结构与算法](https://github.com/overnote/over-algorithm)：以C/Go实现为主记录数据结构与算法的笔记
- [分布式与微服务架构](https://github.com/overnote/over-architecture/)：分布式与微服务等架构相关笔记
- [Linux](https://github.com/overnote/over-linux)：计算机组成原理、操作系统、计算机网络、编译原理基础学科笔记
- [服务端常用技术](https://github.com/overnote/over-server)：nginx、mysql、redis、mongodb、linux系统基础等服务端常用技术汇总笔记
- [大数据](https://github.com/overnote/over-bigdata)：大数据笔记，完善中
- [Python](https://github.com/overnote/over-python)：Python相关笔记，完善中