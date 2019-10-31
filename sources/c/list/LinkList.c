#ifndef _COMMON_
#define _COMMON_
#include <stdio.h>
#include <stdlib.h>
#define OK 1
#define ERROR -1
#define TRUE 1                  
#define FALSE 0
typedef int Status;
typedef int ElemType;           // 定义数据元素类型为int，不直接使用int的原因是可以随时更改
#define LIST_MAX_SIZE 100       // 存储空间的初始分配量
#define LIST_INCREMENT 10       // 存储空间的分配增量
#endif

// 链表节点结构体
typedef struct LinkNode {
    ElemType            data;       // 数据域
    struct  LinkNode    *next;      // 指针域
}LinkNode;

// 链表自身结构体
typedef struct LinkList{
    struct LinkNode     *first;      // 首元节点    
    int                 length;             
}LinkList;

// 构造空链表
Status LinkListInit(LinkList *L) {

    L = malloc(sizeof(LinkList));
    if (L == NULL) {
        printf("内存分配失败\n");
        return ERROR;
    }
    L->first = NULL;
    L->length = 0;
    return OK;
}

// 获取长度
int LinkListLength(LinkList *L){
    return L->length;
}

// 显示表
void LinkListDisplay(LinkList *L){

    if (L->length == 0) {
        printf("链表数据元素个数为0\n");
        return;
    }

    int count = 1;
    struct LinkNode *currentNode = L->first;
    while(count <= L->length) {
        printf("位置 %d 元素为：%d \n", count, currentNode->data);
        currentNode = currentNode->next;
        count++;
     }
}

// 追加元素
Status LinkListAppend(LinkList *L, ElemType e){

    // 构造要插入的节点
    struct LinkNode *insertNode;
    insertNode = malloc(sizeof(LinkNode));
    if (!insertNode) {
        printf("内存分配失败\n");
        return ERROR;
    }
    insertNode->data = e;
    insertNode->next = NULL;

    // 第一次append
    if (L->length == 0) {
        L->first = insertNode;
        L->length++;
        return OK;
    }

    // 当前循环的节点
    struct LinkNode *currentNode = L->first;
    while(currentNode->next != NULL){
        currentNode = currentNode->next;
    }
    currentNode->next = insertNode;
    L->length++;
    return OK;
}

// 插入元素
Status LinkListInsert(LinkList *L, int index, ElemType e){

    // 越界
    if (index < 1 || index > L->length + 1) {
        printf("插入位序不正确\n");
        return FALSE;
    }

    // 构造要插入的节点
    struct LinkNode *insertNode;
    insertNode = malloc(sizeof(LinkNode));
    if (!insertNode) {
        printf("内存分配失败\n");
        return FALSE;
    }
    insertNode->data = e;

    // 如果在第一个结点插入
    if (index == 1) {
        insertNode->next = L->first;
        L->first = insertNode;
        L->length++;
        return TRUE;
    }
    
    // 如果在末尾插入：找到插入位置的前一个节点
    struct LinkNode *currentNode = L->first;
    while ( index > 2){
        currentNode = currentNode->next;
        index--;
    }

    // 切键
    insertNode->next = currentNode->next;
    currentNode->next =insertNode;
    L->length++;
    return TRUE;

}