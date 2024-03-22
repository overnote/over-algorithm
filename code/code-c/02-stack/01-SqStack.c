/**
 * 顺序栈
 */
#include <stdio.h>
#include <stdlib.h>

typedef int ElemType;
#define MAX_SIZE 5

// 顺序表结构体
typedef struct {
    ElemType data[MAX_SIZE]; // 静态分配空间
    int top;                 // 栈顶指针：数组下标
} SqStack;

// 初始化数据
void InitSqStack(SqStack *S) {
    S->top = -1; // 默认top没有指向-1(指向0也是可行的)
}

int IsEmpty(SqStack S) {
    if (S.top == -1) {
        printf("栈为空\n");
        return 1;
    }
    return 0;
}

// 显示动态数组数据
void Display(SqStack S) {
    printf("{ ");
    for (int i = 0; i < S.top; i++) {
        printf("%d ", S.data[i]);
    }
    printf("}\n");
}

// 清空
void ClearStack(SqStack *S) { S->top = -1; }

// 销毁
void DestroyStack(SqStack *S) {}

// 入栈
int Push(SqStack *S, ElemType e) {
    if (S->top == MAX_SIZE - 1) {
        printf("栈已满\n");
        return 0;
    }

    // S.data[++S.top] = e;
    S->top = S->top + 1;
    S->data[S->top] = e;
    return 1;
}

// 出栈
int Pop(SqStack *S, ElemType *e) {
    if (S->top == -1) {
        printf("栈为空\n");
        return 0;
    }

    // e = S.data[S.top --];
    e = S->data[S->top];
    S->top = S->top - 1;
    return 1;
}

// 获取栈顶
int GetTop(SqStack *S, ElemType *e) {
    if (S->top == -1) {
        printf("栈为空\n");
        return 0;
    }

    e = S->data[S->top];
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
