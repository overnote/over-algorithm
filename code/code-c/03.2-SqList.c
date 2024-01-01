/**
 * 顺序表
 *  为了简便，这里不考虑数组扩容，缩容问题，简单申请一个数组即可
 */

#include "SqList.h"
#include <stdio.h>
#include <stdlib.h>

typedef int ElementType;
#define MAX_SIZE 5 // 存储空间的初始分配量，取小值便于测试

// 顺序表结构体
typedef struct {
    ElementType *data; // 存储空间地址，即一个动态数组
    int size;          // 当前长度，即数据元素个数
    int capacity; // 当前分配的存储容量，以 sizeof(ElemType)为单位
} SqList;

// 构造空顺序表
SqList *newSqList() {
    // 申请内存
    SqList *L = malloc(sizeof(SqList));
    if (L == NULL) {
        printf("SqList 内存申请未成功");
        exit(1); // 1 表示异常退出
    }
    ElementType *newData =
        (ElementType *)malloc(MAX_SIZE * sizeof(ElementType));
    if (newData == NULL) {
        printf("SqList 内部空间内存申请未成功");
        exit(1); // 1 表示异常退出
    }
    L->data = newData;

    // 初始化数据
    L->size = 0;
    L->capacity = MAX_SIZE;
    return L;
}

// 增：按照一个位置插入数据
int insert(SqList *L, ElementType e, int index) {
    if (index < 1 || index > L->size + 1) {
        printf("插入位置不合法\n");
        return -1;
    }
    if (L->size == L->capacity) {
        printf("顺序表容量已满\n");
        return -1;
    }

    // 移动元素
    for (int i = L->size; i >= index; i--) {
        L->data[i] = L->data[i - 1];
    }

    // 插入位置值改变
    L->data[index - 1] = e;
    L->size++;
    return 1;
}

// 删：删除位置上的元素，并获取到删除元素的值
int delete(SqList *L, int index, ElementType *e) {
    if (index < 1 || index > L->size) {
        printf("删除位置不合法\n");
        return -1;
    }
    if (L->size == 0) {
        printf("空表无元素可移除\n");
        return -1;
    }

    *e = L->data[index - 1];

    for (int i = index; i < L->size; i++) {
        L->data[i - 1] = L->data[i];
    }
    L->size--;
    return 1;
}

// 改
void update(SqList *L, int index, ElementType e) {
    if (index < 1 || index > L->size) {
        printf("修改位置不合法\n");
        return;
    }
    L->data[index - 1] = e;
}

// 查：根据位置查询值
int search(SqList *L, ElementType *e, int index) {
    if (index < 1 || index > L->size) {
        printf("参数位置不合法\n");
        return -1; // 查找失败，返回-1
    }
    *e = L->data[index - 1];
    return 1;
}

// 查：根据值查询位置
int locate(SqList *L, ElementType e) {
    for (int i = 0; i < L->size; i++) {
        if (L->data[i] == e) {
            return i;
        }
    }
    printf("未找到该元素\n");
    return -1; // 未找到
}

// 获取顺序表最大容量
int capacity(SqList *L) { return L->capacity; }

// 获取顺序表长度
int length(SqList *L) { return L->size; }

// 清空
void clear(SqList *L) { L->size = 0; }

// 销毁
void destroy(SqList *L) {
    if (L->data != NULL) {
        free(L->data);
    }
}

// 显示动态数组数据
void display(SqList *L) {
    if (L->size == 0) {
        printf("[]\n");
        return;
    }

    printf("[");
    for (int i = 0; i < L->size; i++) {
        if (i == L->size - 1) {
            printf("%d", L->data[i]);
        } else {
            printf("%d, ", L->data[i]);
        }
    }
    printf("]\n");
}