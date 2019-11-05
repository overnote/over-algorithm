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

// 双向循环链表

// 链表节点结构体
typedef struct DoubleNode {
    ElemType            data;       // 数据域
    struct  DoubleNode  *next;      // 后继指针域
    struct  DoubleNode  *prev;      // 前驱指针域
}DoubleNode;

// 链表自身结构体
typedef struct DoubleList{
    int                 length;   
    struct  DoubleNode  *first;  // 首元节点指针             
}DoubleList;

// 构造空链表
Status DoubleListInit(DoubleList *L) {

    L = malloc(sizeof(DoubleList));
    if (L == NULL) {
        printf("内存分配失败\n");
        return ERROR;
    }
    L->first = NULL;
    L->length = 0;
    return OK;
}

// 获取长度
int DoubleListLength(DoubleList *L){
    return L->length;
}

// 显示表
void DoubleListDisplay(DoubleList *L){
    if (L->length == 0) {
        printf("链表数据元素个数为0\n");
        return;
    }
    struct DoubleNode *currentNode = L->first;
    printf("链表数据元素为：%d ", currentNode->data);
    while (currentNode->next != L->first){
        currentNode = currentNode->next;
        printf(" %d ", currentNode->data);
    }
    printf("\n");
}

// 追加元素
Status DoubleListAppend(DoubleList *L, ElemType e){

    // 构造要插入的节点
    struct DoubleNode *insertNode;
    insertNode = malloc(sizeof(DoubleNode));
    if (!insertNode) {
        printf("内存分配失败\n");
        return ERROR;
    }
    insertNode->data = e;

    // 第一次append
    if (L->length == 0) {
        insertNode->next = insertNode;
        insertNode->prev = insertNode;
        L->first = insertNode;
        L->length++;
        return OK;
    }

    // 查找到最后一个节点执行拼接
    struct DoubleNode *currentNode = L->first;
    while (currentNode->next != L->first){
        currentNode = currentNode->next;
    }

    currentNode->next = insertNode;
    insertNode->prev = currentNode;
    insertNode->next = L->first;
    L->first->prev = insertNode; 
    L->length++;
    return OK;
}

// 插入元素
Status DoubleListInsert(DoubleList *L, int index, ElemType e){

    // 越界
    if (index < 1 || index > L->length + 1) {
        printf("插入位序不正确\n");
        return FALSE;
    }

    // 构造要插入的节点
    struct DoubleNode *insertNode;
    insertNode = malloc(sizeof(DoubleNode));
    if (!insertNode) {
        printf("内存分配失败\n");
        return FALSE;
    }
    insertNode->data = e;
    
    // 查找插入位置
    int i = 1;                                 
    struct DoubleNode *currentNode = L->first;
    while(currentNode->next != L->first){
        if (i == index) {
            break;
        }
        i++;
        currentNode = currentNode->next;
    }

    // 执行插入
    struct DoubleNode *prevNode = currentNode->prev;
    insertNode->prev = prevNode;
    insertNode->next = currentNode;
    prevNode->next = insertNode;
    currentNode->prev = insertNode;
    L->length++;

    // 考虑头指针是否会改变
    if (index == 1) {
        L->first = insertNode;
    }

    return TRUE;

}

// 删除 按照位序删除
Status DoubleListDelete(DoubleList *L, int index){

    // 越界
    if (index < 1 || index > L->length) {
        printf("索引越界\n");
        return ERROR;
    }

    // 如果只有一个节点
    if (index == 1 && L->length == 1) {
        L->first = NULL;
        L->length = 0;
        return TRUE;        
    }

    // 找到要删除的元素
    int i = 1;                                 
    struct DoubleNode *currentNode = L->first;
    while(currentNode->next != L->first){
        if (i == index) {
            break;
        }
        i++;
        currentNode = currentNode->next;
    }
    
    // 执行删除
    struct DoubleNode *prevNode = currentNode->prev;
    struct DoubleNode *nextNode = currentNode->next;
    prevNode->next = nextNode;
    nextNode->prev = prevNode;
    L->length--;

    // 考虑头指针是否会改变
    if (index == 1) {
        L->first = nextNode;
    }

    return OK;
}

// 修改 按照位序修改
Status DoubleListUpdate(DoubleList *L, int index, ElemType e){

    // 越界
    if (index < 1 || index > L->length) {
        printf("索引越界\n");
        return ERROR;
    }

    // 找到要修改的位置
    int i = 1;
    struct DoubleNode *currentNode = L->first;
    while(currentNode->next != L->first){
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
Status DoubleListGetElem(DoubleList *L, int index, ElemType *e){

    if (index < 1 || index > L->length) {
        printf("位序不合法\n");
        return ERROR;
    }

    int i = 1;
    struct DoubleNode *currentNode = L->first;
    while(currentNode->next != L->first){
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
Status DoubleListLocateElem(DoubleList *L,  ElemType e, int *index) {

    int i = 1;
    struct DoubleNode *currentNode = L->first;
    while(currentNode->next != L->first){
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
Status DoubleListPriorElem(DoubleList *L, ElemType e, ElemType *pe){

    if (L->length <= 1) {
        printf("数据结构元素不足，无法查询\n");
        return ERROR;
    }

    if (L->first->data == e) {
        printf("第一个元素无前驱\n");
        return ERROR;
    }

    int i = 1;
    struct DoubleNode *currentNode = L->first;
    while(currentNode->next != L->first){
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
Status DoubleListNextElem(DoubleList *L, ElemType e, ElemType *ne){

    if (L->length <= 1) {
        printf("数据结构元素不足，无法查询\n");
        return ERROR;
    }

    int i = 1;
    struct DoubleNode *currentNode = L->first;
    while(currentNode->next != L->first){
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
void DoubleListClear(DoubleList *L) {

    int i;
    struct DoubleNode *currentNode = L->first;
    struct DoubleNode *tempNode;

    for(i = 1; i <= L->length; i++){
        tempNode = currentNode;
        currentNode = currentNode->next;
        free(tempNode);
    }
    L->length = 0;
}

// 销毁
void DoubleListDestroy(DoubleList *L) {
    int i;
    struct DoubleNode *currentNode = L->first;
    struct DoubleNode *tempNode;
    for(i = 1; i <= L->length; i++){
        tempNode = currentNode;
        currentNode = currentNode->next;
        free(tempNode);
    }
    L = NULL;
}