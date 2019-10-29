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

### 规范三：数组作为函数参数

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