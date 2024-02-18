/**
 * 顺序队列
 */
#include <stdio.h>
#include <stdlib.h>

typedef int ElemType;
#define MAX_SIZE 5

// 顺序表结构体
typedef struct {
    ElemType data[MAX_SIZE]; // 静态分配空间
    int front, rear;         // 队头与队尾指针
} SqQueue;

// 初始化数据
void InitSqQueue(SqQueue *Q) {
    Q->front = 0; // 队头指针指向队头元素
    Q->rear = 0; // 队尾指针指向接下来应该插入的位置，即队尾的后一个元素位置
}

int IsEmpty(SqQueue Q) {
    if (Q.rear == Q.front) {
        printf("队列为空\n");
        return 1;
    }
    return 0;
}

// 显示动态数组数据
void Display(SqQueue Q) {
    printf("{ ");
    // for (int i = 0; i < S.top; i++) {
    //     printf("%d ", S.data[i]);
    // }
    printf("}\n");
}

// 清空
void ClearStack(SqQueue *Q) { Q->top = -1; }

// 销毁
void DestroyStack(SqQueue *Q) {}

// 入队
int EnQueue(SqQueue *Q, ElemType e) {
    // 初始化时队首队尾都指向0，判断是否为空时使用队首队尾是否相等，那么判断队列是否满则需要放弃一个元素
    // 如果不想浪费该空间，队列可以加入size属性，通过size来判断队空、队满！
    if ((Q->rear + 1) % MAX_SIZE == Q->front) {
        printf("队列已满\n");
        return 0;
    }

    Q->data[Q->rear] = e;
    Q->rear = (Q->rear + 1) % MAX_SIZE; // 队尾指针+1取模
    return 1;
}

// 出队
int DeQueue(SqQueue *Q, ElemType *e) {
    if (Q->rear == Q->front) {
        printf("队列为空\n");
        return 0;
    }

    e = Q->data[Q->front];
    Q->front = (Q->front + 1) % MAX_SIZE;
    return 1;
}

// 获取队头
int GetHead(SqQueue Q, ElemType *e) {
    if (Q.rear == Q.front) {
        printf("队列为空\n");
        return 0;
    }

    e = Q.data[Q.front];
    return 1;
}

// 获取队列数据元素个数
int GetLength(SqQueue Q) { return (Q.rear + MAX_SIZE - Q.front) % MAX_SIZE; }

int main() {
    // SqList L;
    // InitSqList(&L);
    // Display(L);
    // InsertElem(&L, 3, 1);
    // InsertElem(&L, 6, 2);
    // InsertElem(&L, 9, 3);
    // Display(L);
}
