/**
 * 双链表
 */

#include <stdio.h>
#include <stdlib.h>

typedef int ElemType;

typedef struct DNode {
    ElemType data;
    struct DNode *prior, *next;
} DNode, *DLinkList;

// 初始化方式一：因为带头节点， LinkList 声明后需要先初始化
int InitDLinkList(DLinkList *L) {
    *L = (DNode *)malloc(sizeof(DNode));
    if (L == NULL) {
        printf("申请L失败\n");
        exit(1);
        return 0;
    }
    (*L)->data = 0;
    (*L)->next = NULL;

    (*L)->prior = NULL; // 头节点的prior永远指向NULL
    (*L)->next = NULL;  // 刚初始化之后暂时还没next节点
    return 1;
}

// 初始化方式二：直接生成一个 LinkList
DLinkList NewList() {
    DLinkList L = (DNode *)malloc(sizeof(DNode));
    if (L == NULL) {
        printf("申请L失败\n");
        exit(1);
        return 0;
    }

    L->data = 0;
    L->next = NULL;

    L->prior = NULL; // 头节点的prior永远指向NULL
    L->next = NULL;  // 刚初始化之后暂时还没next节点
    return L;
}

void Dsiplay(DLinkList L) {
    if (L == NULL) {
        printf("数据结构为NULL\n");
        return;
    }

    printf("{ ");

    DNode *p = L;
    int j = 0;
    while (p != NULL) {
        p = p->next;
        if (p != NULL) {
            printf("%d ", p->data);
        }
        j++;
    }
    printf("}\n");
}

int IsEmpty(DLinkList L) {
    if (L->next == NULL) {
        printf("为空链表\n");
        return 1;
    }
    printf("链表不为空\n");
    return 0;
}

// 删除p的后继节点
int DeleteNextNode(DNode *p) {
    if (p == NULL) {
        printf("节点参数为空\n");
        return 0;
    }

    DNode *q = p->next; // 找到p的后继节点
    if (q == NULL) {
        printf("p没有后继节点\n");
        return 0;
    }
    p->next = q->next;
    if (q->next != NULL) {
        q->next->prior = p;
    }
    free(q);
    return 1;
}

void DestroyList(DLinkList *L) {
    while ((*L)->next != NULL) {
        DeleteNextNode(*L);
    }
    free(*L);
    L = NULL; // 头指针指向NULL
}

// 后插操作：在节点p之后插入节点s
// 通过后插操作，可以轻松实现按位序插入，前插操作
int InsertNextDNode(DNode *p, DNode *s) {
    if (p == NULL || s == NULL) {
        printf("传入节点参数为空\n");
        return 0;
    }

    s->next = p->next;
    if (p->next != NULL) {
        p->next->prior = s;
    }

    s->prior = p;
    p->next = s;
    return 1;
}

// 前插操作：利用双链表特性，找到给定节点的前驱节点，转换为后插操作
// int InsertPriorDNode(DNode *p, DNode *s) {}

// 按位序插入：从头节点开始找到位序的前驱节点，对该前驱节点执行后插操作

int main() {

    DLinkList L;
    InitDLinkList(&L); // 传入LinkList地址
    // DLinkList L = NewList(); // 也可以将上述2步使用创建方式一步写完

    Dsiplay(L);
    DestroyList(&L); // 销毁
    Dsiplay(L);
    printf("销毁后:%p\n", L);
    // InsertElem(L, 1, 5);
    // Dsiplay(L);
    // InsertElem(L, 2, 6);
    // InsertElem(L, 3, 7);
    // Dsiplay(L);
    // InsertElem(L, 8, 100);
    // DeleteElem(L, 2);
    // Dsiplay(L);

    // LNode *n1 = GetElem(L, 1);
    // printf("%d\n", n1->data);
    // LNode *n2 = GetElem(L, 2);
    // printf("%d\n", n2->data);
    // // LNode *n3 = GetElem(L, 3);
    // // printf("%d\n", n3->data);

    // UpdateList(L, 2, 10);
    // Dsiplay(L);

    return 0;
}