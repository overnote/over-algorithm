#include <stdio.h>
#include <stdlib.h>

typedef int DataType;
typedef struct{
    DataType        data;       // 结点存储的书
    struct  Node    *next;      // 结点指针
}Node, LinkList;

// 创建节点
Node newNode(DataType data){
    Node *node = malloc(sizeof(Node));
    node->data = data;
    return *node;
}

// 创建单链表
LinkList newLinkList(){
    // 申请链表内存：其本身其实就是创建了一个头节点
    LinkList *L = malloc(sizeof(LinkList));
    L->data = 0;
    return *L;
}

// 增：插入节点
int insert(LinkList *L, DataType e, int index){
    // 笔者这里规定：头节点位置为0号位，插入节点只能在头节点之后，也不允许插入超过最大元素个数的位置
    if(index < 1 || index > L->data + 1){
        printf("插入位置不合法");
        return -1;
    }
    
    // 创建要插入的节点
    Node insertNode = newNode(e);
    int pos = 1;    // 指针位置
    while(L->next != NULL){
        if(pos == index){
            L->next = &insertNode;
        } else {
            pos++;
        }
    }
}

// 显示单链表
void display(LinkList *L){
    if(L->data == 0){
        printf("空链表");
        return;
    }
}