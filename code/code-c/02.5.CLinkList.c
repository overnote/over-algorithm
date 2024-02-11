/**
 * 循环链表：双向循环链表
 */

#include <stdio.h>
#include <stdlib.h>

typedef int ElemType;

typedef struct CNode {
    ElemType data;
    struct CNode *prior, *next;
} CNode, *CircleList;

// 初始化方式一：因为带头节点， LinkList 声明后需要先初始化
int InitCircleList(CircleList *L) {
    *L = (CNode *)malloc(sizeof(CNode));
    if (L == NULL) {
        printf("申请L失败\n");
        exit(1);
        return 0;
    }
    (*L)->data = 0;
    (*L)->prior = (*L);
    (*L)->next = (*L); // 头结点next指向头结点自己
    return 1;
}

int IsEmpty(CircleList L) {
    if (L->next == NULL) {
        printf("为空链表\n");
        return 1;
    }
    printf("链表不为空\n");
    return 0;
}

// 判断节点p是否为双向循环链表的表尾节点
int IsTail(CircleList L, CNode *p) {
    if (p->next == NULL) {
        printf("是表尾节点\n");
        return 1;
    }
    printf("不是表尾节点\n");
    return 0;
}

// void Dsiplay(CircleList L) {
//     printf("{ ");

//     LNode *p = L;
//     int j = 0;
//     while (p != NULL) {
//         p = p->next;
//         if (p != NULL) {
//             printf("%d ", p->data);
//         }
//         j++;
//     }
//     printf("}\n");
// }

// 后插操作：在节点p之后插入节点s  与 双链表相比无需做判断
int InsertNextDNode(CNode *p, CNode *s) {
    if (p == NULL || s == NULL) {
        printf("传入节点参数为空\n");
        return 0;
    }

    s->next = p->next;
    p->next->prior = s; // 无需做 p->next != NULL 判断
    s->prior = p;
    p->next = s;
    return 1;
}

int main() { return 0; }