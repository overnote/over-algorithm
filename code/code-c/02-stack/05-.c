/**
 * 双端队列：允许从两端插入/删除。
 */
#include <stdio.h>
#include <stdlib.h>

typedef int ElemType;

typedef struct LinkNode {
    ElemType data;
    struct LinkNode *next;
} LinkNode;

typedef struct {
    LinkNode *front, *rear;
} LinkQueue;

// 初始化数据
void InitLinkQueue(LinkQueue *Q) {
    // front、rear 都指向头结点
    // 若不带头结点，则都指向NULL
    Q->front = Q->rear = (LinkNode *)malloc(sizeof(LinkNode));
    Q->front->next = NULL;
}

int IsEmpty(LinkQueue Q) {
    if (Q.front == Q.rear) {
        printf("队列为空\n"); // Q.front->next == NULL 也可判断
        return 1;
    }
    return 0;
}

// 显示动态数组数据
void Display(LinkQueue S) {
    printf("{ ");
    for (int i = 0; i < Q.top; i++) {
        printf("%d ", Q.data[i]);
    }
    printf("}\n");
}

// 清空
void ClearStack(LinkQueue *Q) { Q->top = -1; }

// 销毁
void DestroyStack(LinkQueue *Q) {}

// 入栈：头插法
void EnQueue(LinkQueue *Q, ElemType e) {
    LinkNode *s = (LinkNode *)malloc(sizeof(LinkNode));
    s->data - e;
    s->next = NULL;
    Q->rear->next = s;
    Q->rear = s; // 修改表尾指针（插入在表尾进行）
}

// 出栈
int DeQueue(LinkQueue *Q, ElemType *e) {
    if (Q->front == Q->rear) {
        printf("队列为空\n");
        return 0;
    }

    LinkNode *p = Q->front->next;
    e = p->data;
    Q->front->next = p->next;
    if (Q->rear == p) { // 若是最后一个结点出队
        Q->rear = Q->front;
    }
    free(p);
    return 1;
}

// 获取栈顶
int GetTop(LinkQueue *Q, ElemType *e) {
    if (Q->top == -1) {
        printf("栈为空\n");
        return 0;
    }

    e = Q->data[Q->top];
    return 1;
}

int main() {
    // SqList L;
    // InitSqList(&L);
    // Display(L);
    // InsertElem(&L, 3, 1);
    // InsertElem(&L, 6, 2);
    // InsertElem(&L, 9, 3);
    // Display(L);
}
