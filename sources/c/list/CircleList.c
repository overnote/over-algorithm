/**
 * 循环链表
 */

#include <stdio.h>
#include <stdlib.h>
#include "CircleList.h"

// 创建结点
Node* newNode(DataType data) {
    Node *node = malloc(sizeof(Node));
    if(node == NULL){
        printf("申请结点内存失败\n");
        exit(1);
    }
    node->data = data;
    node->next = NULL;
    return node;
}

// 创建循环链表
CircleList* newCircleList(){
    // 创建头节点，返回头指针
    Node *head = newNode(0);
    head->next = head;  // 循环链表！！！！！

    CircleList *L = malloc(sizeof(CircleList));
    if(L == NULL){
        printf("申请单链表内存失败\n");
        exit(1);
    }

    L->head = head;
    L->length = 0;
    return L;
}

// 增：插入结点
// 约定：插入时，只能在头结点之后插入，也不允许插入超过最大元素个数的位置
int insert(CircleList *L, DataType e, int index){
    if(index < 1 || index > L->length + 1){
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

    // 创建要插入的节点
    Node *q = newNode(e);
    q->next = p->next;
    p->next = q;

    L->length++;
    return 1;
}

// 删：根据位置删除，返回被删除的元素
int delete(CircleList *L, int index, DataType *e){
    
    if(index < 1 || index > L->length){
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

    // 执行删除
    Node *q = p->next;
    *e = q->data;
    p->next = q->next;
    free(q);

    L->length--;  // 不要忘记存储的长度-1
    return 0;
}

// 改
void update(CircleList *L, int index, DataType e){
    Node *p = locate(L, index);
    if(p == NULL){
        return;
    }
    p->data = e;
}

// 查：根据值查询结点地址
Node* search(CircleList *L, DataType e){
    Node *p = L->head;
    while(p->next != L->head){
        if(p->data == e){
            return p;
        }
        p = p->next;
    }
    return NULL;
}

// 定位
Node* locate(CircleList *L, int index){
    if(index < 0 || index > L->length + 1){
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

// 获取表长度:没有头结点一般使用循环获取长度
int length(CircleList *L){
    return L->length;
}

// 清空表:仅保留头结点
void clear(CircleList *L){
    Node *p = L->head;
    while(p->next != L->head){
        Node *q = p->next;
        p->next = q->next;
        free(q);
    }
    L->head->next = L->head;    // 重新指定头结点循环
    L->length = 0;
}

// 销毁表
void destroy(CircleList *L){
    Node *p = L->head;
    while(p->next != NULL){
        Node *q = p->next;
        p->next = q->next;
        free(q);
    }
    free(L);
}

// 显示循环链表
void display(CircleList *L){
    if(L->length == 0){
        printf("空链表\n");
        return;
    }

    Node *p = L->head;
    int pos = 0;
    while(p->next != NULL){
        if(pos == L->length){
            printf("%d->%d->...\n", p->data, L->head->data);
            break;
        }
        printf("%d->", p->data);
        p = p->next;
        pos++;
    }
}