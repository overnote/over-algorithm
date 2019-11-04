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
    struct LinkNode *currentNode = L->first;
    printf("链表数据元素为：%d ", currentNode->data);
    while (currentNode->next != NULL){
        currentNode = currentNode->next;
        printf(" %d ", currentNode->data);
    }
    printf("\n");
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

    // 查找到最后一个节点执行拼接
    struct LinkNode *currentNode = L->first;
    while (currentNode->next != NULL){
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
    
    // 从第一位first开始查找，找到插入位置的前一个节点
    int i = 1;                                 
    struct LinkNode *currentNode = L->first;
    while(currentNode->next != NULL){
        if (i == index - 1) {
            break;
        }
        i++;
        currentNode = currentNode->next;
    }

    // 执行插入
    insertNode->next = currentNode->next;
    currentNode->next =insertNode;
    L->length++;
    return TRUE;

}

// 删除 按照位序删除
Status LinkListDelete(LinkList *L, int index){

    // 越界
    if (index < 1 || index > L->length) {
        printf("索引越界\n");
        return ERROR;
    }

    // 如果删除的是第一个元素
    if (index == 1 && L->length >= 1) {
        if (L->length == 1) {
            L->first->data = 0;
        } else {
            L->first = L->first->next;
        }
        L->length--;
        return TRUE;
    }

    // 找到要删除元素的前一个元素
    int i = 1;                                 
    struct LinkNode *currentNode = L->first;
    while(currentNode->next != NULL){
        if (i == index - 1) {
            break;
        }
        i++;
        currentNode = currentNode->next;
    }
    
    // 执行删除
    currentNode->next = currentNode->next->next;
    L->length--;
    return OK;
}

// 修改 按照位序修改
Status LinkListUpdate(LinkList *L, int index, ElemType e){

    // 越界
    if (index < 1 || index > L->length) {
        printf("索引越界\n");
        return ERROR;
    }

    // 找到要修改的位置
    int i = 1;
    struct LinkNode *currentNode = L->first;
    while(currentNode->next != NULL){
        if (i == index) {
            break;
        }
        i++;
        currentNode = currentNode->next;
    }
    currentNode->data = e;
    return OK;
}

// 查询 按照位序查询值
Status LinkListGetElem(LinkList *L, int index, ElemType *e){

    if (index < 1 || index > L->length) {
        printf("位序不合法\n");
        return ERROR;
    }

    int i = 1;
    struct LinkNode *currentNode = L->first;
    while(currentNode->next != NULL){
        if (i == index) {   
            break;
        }
        i++;
        currentNode = currentNode->next;
    }
    *e = currentNode->data;
    return OK;
}

// 查询  按照值查询位序
Status LinkListLocateElem(LinkList *L,  ElemType e, int *index) {
    int i = 1;
    struct LinkNode *currentNode = L->first;

    while(currentNode->next != NULL){
        if (currentNode->data == e) {
            break;
        }
        i++;
        currentNode = currentNode->next;
    }

    if (i == L->length && currentNode->data != e) {
        printf("未找到该元素\n");
        return FALSE;
    }

    *index = i;
    return TRUE;
}

// 查询前驱
Status LinkListPriorElem(LinkList *L, ElemType e, ElemType *pe){

    if (L->length <= 1) {
        printf("数据结构元素不足，无法查询\n");
        return ERROR;
    }

    if (L->first->data == e) {
        printf("第一个元素无前驱\n");
        return ERROR;
    }

    int i = 1;
    struct LinkNode *currentNode = L->first;
    while(currentNode->next != NULL){
        if (currentNode->next->data == e) {
            *pe = currentNode->data;
            return TRUE;      
        }
        i++;
        currentNode = currentNode->next;
    }

    printf("未找到元素\n");
    return FALSE;      
}

// 查询后继
Status LinkListNextElem(LinkList *L, ElemType e, ElemType *ne){

    if (L->length <= 1) {
        printf("数据结构元素不足，无法查询\n");
        return ERROR;
    }

    int i = 1;
    struct LinkNode *currentNode = L->first;
    while(currentNode->next != NULL){
        if (currentNode->data == e) {
            break;
        }
        i++;
        currentNode = currentNode->next;
    }

    if (i == L->length && currentNode->data != e) {
        printf("未找到该元素\n");
        return FALSE;
    }

    if (i == L->length && currentNode->data == e) {
        printf("最后一个元素无后继\n");
        return FALSE;
    }

    *ne = currentNode->next->data;
    return TRUE;      

}

// 清空表：即重置为空表
void LinkListClear(LinkList *L) {
    int i;
    struct LinkNode *currentNode = L->first;
    struct LinkNode *tempNode;

    for(i = 1; i <= L->length; i++){
        tempNode = currentNode;
        currentNode = currentNode->next;
        free(tempNode);
    }
    L->length = 0;
}

// 销毁
void LinkListDestroy(LinkList *L) {
    int i;
    struct LinkNode *currentNode = L->first;
    struct LinkNode *tempNode;
    for(i = 1; i <= L->length; i++){
        tempNode = currentNode;
        currentNode = currentNode->next;
        free(tempNode);
    }
    L = NULL;
}