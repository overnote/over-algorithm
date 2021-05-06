/**
 * 双向循环链表
 */

#include <stdio.h>
#include <stdlib.h>
#include "DblList.h"

// 创建节点
Node* newNode(DataType data){
    Node *node = malloc(sizeof(Node));
    if(node == NULL){
        printf("申请结点内存失败\n");
        exit(1);       
    }

    node->data = data;
    node->prev = NULL;
    node->next = NULL;
    return node;
}

// 创建双向循环链表
DblList* newDblList(){
    // 创建头节点，返回头指针
    Node *head = newNode(0);
    head->prev = head;
    head->next = head;

    // 创建双向循环链表
    DblList *L = malloc(sizeof(DblList));
    if(L == NULL){
        printf("申请单链表内存失败\n");
        exit(1);
    }

    L->head = head;
    L->size = 0;
    return L;
}

// 增：插入节点。约定只能在头节点之后插入，且不超过最大元素个数位置
int insert(DblList *L, DataType e, int index){

    if(index < 1 || index > L->size + 1){
        printf("插入位置不合法\n");
        return -1;
    }

    // 找到插入位置前一个位置：也可以使用 locate函数
    Node *p = L->head;
    int k = 0;
    while(p->next != L->head && k < index - 1){
        p = p->next;
        k++;
    }

    // 创建要插入的结点
    Node *q = newNode(e);
    q->prev = p;
    q->next = p->next;

    // 执行插入
    p->next = q;

    L->size++;  // 不要忘记存储的长度+1
    return 1;
}

// 删：根据位置删除，返回被删除的元素
int delete(DblList *L, int index, DataType *e){
    
    if(index < 1 || index > L->size){
        printf("删除位置非法\n");
        return -1;
    }

    // 找到删除位置前一个位置：也可以使用 locate函数
    Node *p = L->head;
    int k = 0;
    while(p->next != L->head && k < index - 1){
        p = p->next;
        k++;
    }

    // 获取删除元素数据
    Node *q = p->next;
    *e = q->data;

    // 执行删除
    q->next->prev = p;
    p->next = q->next;
    free(q);

    L->size--;  // 不要忘记存储的长度-1
    return 0;
}

// 改
void update(DblList *L, int index, DataType e){
    Node *p = locate(L, index);
    if(p == NULL){
        return;
    }
    p->data = e;
}

// 查：根据值查询结点地址
Node* search(DblList *L, DataType e){
    Node *p = L->head;
    while(p->next != L->head){
        if(p->data == e){
            break;
        }
        p = p->next;
    }

    if(p->data == e){
        return p;
    } else {
        return NULL;
    }
}

// 定位
Node* locate(DblList *L, int index){
    if(index < 0 || index > L->size + 1){
        printf("获取位置不合法\n");
        return NULL;
    }
    Node *p = L->head;
    int k = 0;

    while(p->next != L->head && k < index){
        p = p->next;
        k++;
    }

    return p;
}

// 根据值快速找到前驱
Node* prevElem(DblList *L, DataType e){
    Node *p = search(L, e);
    if(p == NULL){
        printf("未找到该元素\n");
        return NULL;
    }
    return p->prev;
}

// 根据值快速找到后继
Node* nextElem(DblList *L, DataType e){
    Node *p = search(L, e);
    if(p == NULL){
        printf("未找到该元素\n");
        return NULL;
    }
    return p->next;
}

// 获取表长度:没有头结点一般使用循环获取长度
int length(DblList *L){
    return L->size;
}

// 清空表:仅保留头结点
void clear(DblList *L){
    Node *p = L->head;
    while(p->next != L->head){
        Node *q = p->next;
        p->next = q->next;
        free(q);
    }

    L->head->prev = L->head;    // 重新指定头结点循环
    L->head->next = L->head;    // 重新指定头结点循环
    L->size = 0;
}

// 销毁表
void destroy(DblList *L){
    Node *p = L->head;
    while(p->next != NULL){
        Node *q = p->next;
        p->next = q->next;
        free(q);
    }
    free(L);
}

// 显示表
void display(DblList *L){
    if(L->size == 0){
        printf("空链表\n");
        return;
    }

    Node *p = L->head;
    int pos = 0;
    while(p->next != NULL){
        if(pos == L->size){
            printf("%d<->%d<->...\n", p->data, L->head->data);
            break;
        }
        printf("%d<->", p->data);
        p = p->next;
        pos++;
    }
}