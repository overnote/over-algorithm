/**
 * 单链表
 */

#include <stdio.h>
#include <stdlib.h>
#include "LinkList.h"

// 创建结点
Node* newNode(DataType data){
    Node *node = malloc(sizeof(Node));
    if(node == NULL){
        printf("申请结点内存失败\n");
        exit(1);
    }
    node->data = data;
    node->next = NULL;
    return node;
}

// 创建单链表
LinkList* newLinkList(){
    // 创建头结点，返回头指针
    Node *head = newNode(0);

    LinkList *L = malloc(sizeof(LinkList));
    L->head = head;
    L->length = 0;
    return L;
}

// 增：插入结点
// 约定：带头结点的链表，插入时，只能在头结点之后插入，也不允许插入超过最大元素个数的位置
int insert(LinkList *L, DataType e, int index){

    if(index < 1 || index > L->length + 1){
        printf("插入位置非法\n");
        return -1;
    }

    // 找到插入位置前一个位置：也可以使用 locate函数
    Node *p = L->head;
    int k = 0;
    while(p->next != NULL && k < index - 1){
        p = p->next;
        k++;
    }

    // 创建要插入的结点
    Node *q = newNode(e);
    q->next = p->next;
    p->next = q;

    L->length++;  // 不要忘记存储的长度+1
    return 1;
}

// 删：根据位置删除，返回被删除的元素
int delete(LinkList *L, int index, DataType *e){

    if(index < 1 || index > L->length){
        printf("删除位置非法\n");
        return -1;
    }

    // 找到删除位置前一个位置：也可以使用 locate函数
    Node *p = L->head;
    int k = 0;
    while(p->next != NULL && k < index - 1){
        p = p->next;
        k++;
    }

    // 执行删除
    Node *q = p->next;
    *e = q->data;
    p->next = q->next;
    free(q);

    L->length--;  // 不要忘记存储的长度-1
    return 0;
}

// 改
void update(LinkList *L, int index, DataType e){

    // 找到修改位置：这里使用 locate 函数
    Node *p = locate(L, index);
    if(p == NULL){
        return;
    }
    p->data = e;
}

// 查：根据值查询结点地址
Node* search(LinkList *L, DataType e){
    Node *p = L->head;
    while(p->next != NULL){
        if(p->data == e){
            return p;
        }
        p = p->next;
    }
    return NULL;
}

// 查：根据位置查询结点地址
Node* locate(LinkList *L, int index){
    if(index < 0 || index > L->length + 1 ){
        printf("获取位置不合法\n");
        return NULL;
    }

    Node *p = L->head;
    int k = 0;
    while(p->next != NULL && k < index){
        p = p->next;
        k++;
    }
    return p;
}

// 获取表长度：若未保存该值，则可以通过循环获得
int length(LinkList *L){
    return L->length;
}

// 清空表:仅保留头结点
void clear(LinkList *L){
    Node *p = L->head;
    while(p->next != NULL){
        Node *q = p->next;
        p->next = q->next;
        free(q);
    }
    L->length = 0;
}

// 销毁表
void destroy(LinkList *L){
    Node *p = L->head;
    while(p->next != NULL){
        Node *q = p->next;
        p->next = q->next;
        free(q);
    }
    free(L);
}

// 显示单链表
void display(LinkList *L){
    if(L->length == 0){
        printf("空链表\n");
        return;
    }

    Node *p = L->head;
    int pos = 0;
    while(p != NULL){
        if(pos == L->length){    // 最后一位
            printf("%d\n", p->data);
            break;
        }
        printf("%d->", p->data);
        p = p->next;
        pos++;
    }
}