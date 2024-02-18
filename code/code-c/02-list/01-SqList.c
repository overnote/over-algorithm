/**
 * 顺序表
 * 本例使用静态分配方式：ElemType data[MAX_SIZE]
 * 为了简便，这里不考虑数组扩容，缩容问题，简单申请一个数组即可
 */
#include <stdio.h>
#include <stdlib.h>

typedef int ElemType;
#define MAX_SIZE 5 // 存储空间的初始分配量，取小值便于测试

// 顺序表结构体
typedef struct {
    ElemType data[MAX_SIZE]; // 静态分配空间
    int length;              // 当前长度，即数据元素个数
} SqList;

// 初始化数据
void InitSqList(SqList *L) {
    L->length = 0;
    for (int i = 0; i < MAX_SIZE - 1; i++) {
        L->data[i] = 0;
    }
}

// 显示动态数组数据
void Display(SqList L) {
    printf("{ ");
    for (int i = 0; i < L.length; i++) {
        printf("%d ", L.data[i]);
    }
    printf("}\n");
}

int Length(SqList L) { return L.length; }

// 清空
void ClearList(SqList *L) { L->length = 0; }

// 销毁
void DestroyList(SqList *L) {
    if (L->data != NULL) {
        free(L->data);
    }
}

// 增：按照一个位置插入数据
int InsertElem(SqList *L, int idx, ElemType e) {
    if (idx < 1 || idx > L->length + 1) {
        printf("插入位置不合法\n");
        return 0;
    }
    if (L->length == MAX_SIZE) {
        printf("顺序表容量已满\n");
        return 0;
    }

    // 移动元素
    for (int i = L->length; i >= idx; i--) {
        L->data[i] = L->data[i - 1];
    }

    // 插入位置值改变
    L->data[idx - 1] = e;
    L->length++;
    return 1;
}

// 删：删除位置上的元素，并获取到删除元素的值
int DeleteElem(SqList *L, int idx, ElemType *e) {
    if (idx < 1 || idx > L->length) {
        printf("删除位置不合法\n");
        return -1;
    }
    if (L->length == 0) {
        printf("空表无元素可移除\n");
        return -1;
    }

    *e = L->data[idx - 1];

    for (int i = idx; i < L->length; i++) {
        L->data[i - 1] = L->data[i];
    }
    L->length--;
    return 1;
}

// 改
int UpdateElem(SqList *L, int idx, ElemType e) {
    if (idx < 1 || idx > L->length) {
        printf("修改位置不合法\n");
        return 0;
    }
    L->data[idx - 1] = e;
    return 1;
}

// 按位查找
ElemType GetElem(SqList L, int idx, ElemType *e) {
    if (idx > MAX_SIZE || idx < 1) {
        printf("越界\n");
        return 0;
    }
    e = L.data[idx - 1];
    return 1;
}

// 按值查找
int LocateElem(SqList L, int *idx, ElemType e) {
    for (int i = 0; i < L.length; i++) {
        if (L.data[i] == e) {
            idx = i + 1;
            return 1;
        }
    }
    printf("未找到");
    return 0;
}

int main() {
    SqList L;
    InitSqList(&L);
    Display(L);
    InsertElem(&L, 3, 1);
    InsertElem(&L, 6, 2);
    InsertElem(&L, 9, 3);
    Display(L);
}
