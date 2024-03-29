/**
 * 单链表
 */

#include <stdio.h>
#include <stdlib.h>

typedef int ElemType;

typedef struct LNode {
    ElemType data;
    struct LNode *next;
} LNode, *LinkList;

// 初始化方式一：因为带头节点， LinkList 声明后需要先初始化
int InitLinkList(LinkList *L) {
    *L = (LNode *)malloc(sizeof(LNode));
    if (L == NULL) {
        printf("申请L失败\n");
        exit(1);
        return 0;
    }
    (*L)->data = 0;
    (*L)->next = NULL;
    return 1;
}

// 初始化方式二：直接生成一个 LinkList
LinkList NewList() {
    LinkList L = (LNode *)malloc(sizeof(LNode));
    if (L == NULL) {
        printf("申请L失败\n");
        exit(1);
        return 0;
    }

    L->data = 0;
    L->next = NULL;
    return L;
}

void Dsiplay(LinkList L) {
    printf("{ ");

    LNode *p = L;
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

// 按位查找节点
LNode *GetElem(LinkList L, int idx) {
    if (idx < 0) {
        printf("查找位置不合法\n");
        return NULL;
    }
    LNode *p = (LNode *)L; // 当前循环到的节点
    int j = 0;             // 循环到 idx - 1 的位置
    while (p != NULL && j < idx) {
        p = p->next; // 遇到边界时 p 也为NULL
        j++;
    }
    return p;
}

// 按值查找
LNode *LocateElem(LinkList L, ElemType e) {
    LNode *p = L->next;
    while (p != NULL && p->data != e) {
        p = p->next;
    }
    return p; // 否则返回NULL
}

// 带头节点的按位插入：将头节点视为第0个节点
int InsertElem(LinkList L, int idx, ElemType e) {

    if (idx < 1) {
        printf("插入位置不合法\n");
        return 0;
    }

    // 不带头节点时，这里需要额外判断 idx = 1 的情况

    // LNode *p = (LNode *)L; // 当前循环到的节点
    // int j = 0; // 当前循环到的节点是第几个节点，最终循环到插入位置idx的前一位
    // while (p != NULL && j < idx - 1) {
    //     p = p->next;
    //     j++;
    // }
    LNode *p = GetElem(L, idx - 1); // 化简为一步

    if (p == NULL) {
        printf("插入位置不合法\n");
        return 0;
    }

    LNode *s = (LNode *)malloc(sizeof(LNode));
    s->data = e;
    s->next = p->next;
    p->next = s;
    return 1;
}

int DeleteElem(LinkList L, int idx) {
    if (idx < 1) {
        printf("删除位置不合法\n");
        return 0;
    }

    // 不带头节点时，这里需要额外判断 idx = 1 的情况

    // LNode *p = (LNode *)L; // 当前循环到的节点
    // int j = 0;             // 循环到 idx - 1 的位置
    // while (p != NULL && j < idx - 1) {
    //     p = p->next;
    //     j++;
    // }
    LNode *p = GetElem(L, idx - 1); // 化简为一步

    if (p == NULL) {
        printf("删除位置不合法\n");
        return 0;
    }
    if (p->next == NULL) {
        printf("删除位置不合法\n"); // idx为最末尾的后一位
        return 0;
    }

    LNode *q = p->next; // 令 q 指向被删除节点
    p->next = q->next;
    free(q);
    return 1;
}

int UpdateList(LinkList L, int idx, ElemType e) {
    if (idx < 1) {
        printf("查找位置不合法\n");
        return 0;
    }

    LNode *p = GetElem(L, idx);
    if (p == NULL) {
        printf("查找位置不合法\n");
        return 0;
    }
    p->data = e;
    return 1;
}

// 后插操作：在 p 节点之后插入元素e
int InsertNextNode(LNode *p, ElemType e) {
    if (p == NULL) {
        printf("p 为空\n");
        return 0;
    }

    LNode *s = (LNode *)malloc(sizeof(LNode));
    if (s == NULL) {
        printf("申请节点内存失败\n");
        return 0;
    }

    s->data = e;
    s->next = p->next;
    p->next = s;
    return 1;
}

// 前插操作：在 p 节点之前插入元素e
int InsertPriorNode(LNode *p, ElemType e) {
    if (p == NULL) {
        printf("p 为空\n");
        return 0;
    }

    LNode *s = (LNode *)malloc(sizeof(LNode));
    if (s == NULL) {
        printf("申请节点内存失败\n");
        return 0;
    }

    s->next = p->next;
    p->next = s;       // 新节点s连到p之后
    s->data = p->data; // p中的元素复制到s中
    p->data = e;       // p 中的元素覆盖为e
    return 1;
}

// 尾插法
LinkList TailInsert() {
    LinkList L = (LinkList)malloc(sizeof(LNode));

    int x;
    LNode *s, *r = L;
    scanf("%d", &x);
    while (x != 9999) {
        s = (LNode *)malloc(sizeof(LNode));
        s->data = x;
        r->next = s;
        r = s;
        scanf("%d", &x);
    }
    r->next = NULL;
    return L;
}

// 头插法
LinkList HeadInsert() {
    LinkList L = (LinkList)malloc(sizeof(LNode));
    L->next = NULL;

    int x;
    LNode *s;
    scanf("%d", &x);
    while (x != 9999) {
        s = (LNode *)malloc(sizeof(LNode));
        s->data = x;
        s->next = L->next;
        L->next = s;
        scanf("%d", &x);
    }
    return L;
}

int main() {

    LinkList L;
    InitLinkList(&L); // 传入LinkList地址
    // LinkList L = NewList(); // 也可以将上述2步使用创建方式一步写完

    Dsiplay(L);
    InsertElem(L, 1, 5);
    Dsiplay(L);
    InsertElem(L, 2, 6);
    InsertElem(L, 3, 7);
    Dsiplay(L);
    InsertElem(L, 8, 100);
    DeleteElem(L, 2);
    Dsiplay(L);

    LNode *n1 = GetElem(L, 1);
    printf("%d\n", n1->data);
    LNode *n2 = GetElem(L, 2);
    printf("%d\n", n2->data);
    // LNode *n3 = GetElem(L, 3);
    // printf("%d\n", n3->data);

    UpdateList(L, 2, 10);
    Dsiplay(L);
    return 0;
}