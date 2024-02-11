#include <stdio.h>
#include <stdlib.h>

typedef int ElemType;

typedef struct LNode {
    ElemType data;
    struct LNode *next;
} LNode, *LinkList;

// 如果需要带头节点，则 LinkList 声明后需要先初始化
int InitLinkList(LinkList &L) { 
    L = (LNode *)malloc(sizeof(LNode));
    if(L == NULL){
        printf("申请L失败\n");
        return 0;
    }

    L->next = NULL;
    return 1;
}

void Dsiplay(LinkList &L) {
    printf("{ ");

    LNode *p = L;   
    int j =  0; 
    while (p != NULL) {
        p = p->next;
        if(p != NULL){
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
    while (p != NULL && p->data != e){
        p = p->next;
    }
    return p;   // 否则返回NULL
}

// 带头节点的按位插入：将头节点视为第0个节点
int InsertElem(LinkList &L, int idx, ElemType e){

    if(idx < 1){
        printf("插入位置不合法\n");
        return 0;
    }

    // 不带头节点时，这里需要额外判断 idx = 1 的情况

    LNode *p = GetElem(L, idx - 1); // 找到 idx - 1 位置
    if(p == NULL){
        printf("插入位置不合法\n");
        return 0;
    }

    LNode *s = (LNode *)malloc(sizeof(LNode));
    s->data = e;
    s->next = p->next;
    p->next = s;
    return 1;
}


int DeleteElem(LinkList &L, int idx){
        if(idx < 1){
        printf("删除位置不合法\n");
        return 0;
    }

    // 不带头节点时，这里需要额外判断 idx = 1 的情况

    LNode *p = GetElem(L, idx - 1); // 找到 idx - 1 位置
    if(p == NULL){
        printf("删除位置不合法\n");
        return 0;
    }
    if(p->next == NULL){
        printf("删除位置不合法\n"); // idx为最末尾的后一位
        return 0;
    }

    LNode *q = p->next;     // 令 q 指向被删除节点
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

// 在 p 节点之后插入元素e
int InsertNextNode(LNode *p, ElemType e){
    if(p == NULL){
        printf("p 为空\n");
        return 0;
    }

    LNode *s = (LNode *)malloc(sizeof(LNode));
    if(s == NULL){
        printf("申请节点内存失败\n");
        return 0;
    }

    s->data = e;
    s->next = p->next;
    p->next = s;
    return 1;
}

// 在 p 节点之前插入元素e
int InsertPriorNode(LNode *p, ElemType e){
    if(p == NULL){
        printf("p 为空\n");
        return 0;
    }

    LNode *s = (LNode *)malloc(sizeof(LNode));
    if(s == NULL){
        printf("申请节点内存失败\n");
        return 0;
    }

    s->next = p->next; 
    p->next = s;    // 新节点s连到p之后 
    s->data = p->data; // p中的元素复制到s中
    p->data = e;    // p 中的元素覆盖为e
    return 1;
}

// 尾插法
LinkList TailInsert() { 
    LinkList L = (LinkList)malloc(sizeof(LNode));

    int x;
    LNode *s, *r = L;
    scanf("%d", &x);
    while (x != 9999){
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
    while (x != 9999){
        s = (LNode *)malloc(sizeof(LNode));
        s->data = x;
        s->next = L->next;
        L->next = s;
        scanf("%d", &x);
    }
    return L;
}

int main() {

    LinkList L; // 声明一个空的单链表指针
    InitLinkList(L);
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
    return 0;
}