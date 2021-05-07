/**
 * 静态链表
 */
#include <stdio.h>
#include <stdlib.h>
#include "StaticList.h"

// 构造列表节点
Node* newNode(DataType data, int cur){
    Node *node = malloc(sizeof(Node));
    if(node == NULL){
        printf("申请结点内存失败\n");
        exit(1);
    }

    node->data = data;
    node->cur = cur;
    return node;
}

// 构造静态链表：其实是构建一个完整的备用链表
StaticList* newStaticList(){
    // 申请内存
    StaticList *L = malloc(sizeof(StaticList));
    if(L == NULL){
        printf("StaticList 内存申请未成功");
        exit(1);
    }

    // 初始化数组
    for(int i = 0; i < MAX_SIZE - 1; i++){
        L->head[i].cur = i + 1; // 第i个元素的游标指向i+1
    }
    L->head[MAX_SIZE - 1].cur = 0;   // 链表最后一个节点的游标为0

    // 初始化数据
    L->head[MAX_SIZE - 1].cur = 0;
    L->size = 0;
    L->capacity = MAX_SIZE;
    return L;  
}

// 显示表
void display(StaticList *L){
    // 显示数据
    if(L->head[L->capacity - 1].data == 0){
        printf("空链表\n");
        return;
    }
}