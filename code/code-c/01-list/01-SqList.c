/**
 * 顺序表
 * 为了简便，这里不考虑数组扩容，缩容问题，简单申请一个数组即可
 */
#include <stdio.h>
#include <stdlib.h>

typedef int ElemType;
#define MAX_SIZE 5 // 存储空间的初始分配量，取小值便于测试

// 顺序表结构体
typedef struct {
    ElemType data[MAX_SIZE]; // 静态分配空间
    // ElemType *data; // 动态分配
    int length; // 当前长度，即数据元素个数
} SqList;

// 初始化方式一 初始化数据
void InitSqList(SqList *L) {
    L->length = 0;

    // 动态分配时数组需要这样额外进行初始化
    // L->data = (ElemType *)malloc(sizeof(ElemType) * MAX_SIZE);
}

// 初始化方式二 初始化数据
// SqList InitSqList() {
//     SqList *L = (SqList *)malloc(sizeof(SqList));
//     L->length = 0;
//     // 动态分配时数组需要这样额外进行初始化
//     // L->data = (ElemType *)malloc(sizeof(ElemType) * MAX_SIZE);
//     return *L;
// }

int Length(SqList L) { return L.length; }

// 清空
void ClearList(SqList *L) { L->length = 0; }

// 显示顺序表
void Display(SqList L) {
    printf("[ ");
    for (int i = 0; i < L.length; i++) {
        printf("%d ", L.data[i]);
    }
    printf("]\n");
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
    for (int i = L->length - 1; i >= idx - 1; i--) {
        L->data[i + 1] = L->data[i];
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
int GetElem(SqList L, int idx, ElemType *e) {
    if (idx > MAX_SIZE || idx < 1) {
        printf("按位查找:越界\n");
        return 0;
    }
    *e = L.data[idx - 1];
    return 1;
}

// 按值查找
int LocateElem(SqList L, ElemType e, int *idx) {
    for (int i = 0; i < L.length; i++) {
        if (L.data[i] == e) {
            *idx = i + 1;
            return 1;
        }
    }
    printf("按值查找:未找到");
    return 0;
}

int main() {
    // 第一种初始化方式下的声明赋值操作
    SqList L;
    InitSqList(&L);
    //  SqList L = InitSqList();    // 第二种初始化方式下的声明赋值操作
    printf("----执行了初始化操作----\n");
    Display(L);

    InsertElem(&L, 1, 5);
    InsertElem(&L, 2, 6);
    InsertElem(&L, 3, 1);
    InsertElem(&L, 4, 2);
    InsertElem(&L, 5, 3);
    printf("----执行了新增操作----\n");
    Display(L);

    ElemType e1;
    DeleteElem(&L, 5, &e1);
    printf("----执行了删除操作----\n");
    printf("删除位置5 e1 = %d\n", e1);
    Display(L);

    UpdateElem(&L, 4, 10);
    printf("----执行了更新操作----\n");
    Display(L);

    ElemType e2;
    GetElem(L, 3, &e2);
    printf("----执行了查找操作----\n");
    printf("按位查找的e2 = %d\n", e2);

    int idx1;
    LocateElem(L, 1, &idx1);
    printf("----执行了查找操作----\n");
    printf("按位查找的idx1 = %d\n", idx1);

    ClearList(&L);
    printf("----执行了清空操作----\n");
    Display(L);
}
