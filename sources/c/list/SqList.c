#ifndef _COMMON_
#define _COMMON_
#include <stdio.h>
#include <stdlib.h>
#define OK 1
#define ERROR -1
#define TRUE 1                  
#define FALSE 0
typedef int Status;
typedef int ElemType;           // 定义数据元素类型为int，不直接使用int的原因是可以随时更改
#define LIST_MAX_SIZE 100       // 存储空间的初始分配量
#define LIST_INCREMENT 10       // 存储空间的分配增量
#endif


// 顺序表结构体
typedef struct SqList{
    ElemType    *arr;           // 存储空间基址，即一个动态数组
    int         length;         // 当前长度
    int         size;           // 当前分配的存储容量，以 sizeof(ElemType)为单位
}SqList;

// 构造空顺序表
Status SqListInit(SqList *L) {

    L = malloc(sizeof(SqList));
    if (L == NULL) {
        printf("内存分配失败\n");
        return ERROR;   
    }

    L->arr = (ElemType *)malloc(LIST_MAX_SIZE * sizeof(ElemType));
    if (!L->arr)  {            // 存储分配失败
        printf("内存分配失败\n");
        return ERROR;         
    }
    L->length = 0;
    L->size = LIST_MAX_SIZE;
    return OK;
}

// 获取线性表长度
int SqListLength(SqList *L) {
    return L->length;
}

// 获取最大容量
int SqListSize(SqList *L) {
    return L->size;
}

// 显示表
void SqListDisplay(SqList *L){
    if (L->length == 0) {
        printf("为空表\n");
        
    } else {
        int i;
        for (i = 0; i < L->length; i++) {
            printf("元素[%d] = %d\n", i + 1, L->arr[i]);
        }
    }
}

// 添加：向末尾添加元素
Status SqListAppend(SqList *L, ElemType e) {

    // 如果存储空间已满，增加分配
    ElemType *newbase;      // 增加分配后的新地址
    if (L->length >= L->size) {

        newbase = (ElemType *)realloc(L->arr, (L->size + LIST_INCREMENT) * sizeof(ElemType));
        if ( !newbase ) {
            printf("内存分配失败\n");
            return ERROR; 
        }

        L->arr = newbase;
        L->size += LIST_INCREMENT;
    }

    // append
    L->arr[L->length] = e;
    L->length++;
    return OK;
}

// 插入 与数组不同，顺序表是从第1个位置开始
Status SqListInsert(SqList *L, int index, ElemType e) {

    // 位序越界
    if (index < 1 || index > L->length + 1) {
        printf("位序不合法\n");
        return ERROR;
    }

    // 存储空间已满，增加分配
    ElemType *newbase;      // 增加分配后的新地址
    if (L->length >= L->size) {

        newbase = (ElemType *)realloc(L->arr, (L->size + LIST_INCREMENT) * sizeof(ElemType));
        if ( !newbase ) {
            printf("内存分配失败\n");
            return ERROR;
        }

        L->arr = newbase;
        L->size += LIST_INCREMENT;
    }

    // 插入
    int i;      // 最后一个元素的位序的后一位，即为 length
    for (i =  L->length + 1; i > index; i--) {
        L->arr[i - 1] = L->arr[i - 2];
    }
    L->arr[index - 1] = e;
    L->length++;
    return OK;
}

// 删除  按照位序删除
Status SqListDelete(SqList *L, int index) {

    if (index < 1 || index > L->length) {
        printf("索引越界\n");
        return ERROR;
    }

    // 删除
    int i;
    for (i = index - 1; i < L->length - 1; i++) {
        L->arr[i] = L->arr[i + 1];
    }
    L->length--;
    return OK;
}

// 修改  按照位序修改
Status SqListUpdate(SqList *L, int index, ElemType e) {
    
    if (index < 1 || index > L->length - 1) {
        printf("位序不合法\n");
        return ERROR;
    }
    
    L->arr[index - 1] = e;
    return OK;
}

// 查询  按照位序查询值
Status SqListGetElem(SqList *L, int index, ElemType *e) {
    if (index < 1 || index > L->length) {
        printf("位序不合法\n");
        return ERROR;
    }
    *e = L->arr[index - 1];
    return OK;
}

// 查询  按照值查询位序
Status SqListLocateElem(SqList *L,  ElemType e, int *index) {
    int i;
    for (i = 0; i < L->length; i++) {
        if (L->arr[i] == e) {
            *index = i + 1;
            return TRUE;
        }
    }
    return FALSE;
}

// 查询前驱
Status SqListPriorElem(SqList *L, ElemType e, ElemType *pe){

    if (L->length <= 1) {
        printf("顺序表元素不足，无法查询\n");
        return ERROR;
    }

    if (L->arr[0] == e) {
        printf("第一个元素无前驱\n");
        return ERROR;
    }

    // 查找 e 所在位置
    int index;
    int loacateResult;
    loacateResult = SqListLocateElem(L, e, &index);
    if (loacateResult != TRUE) {
        printf("未找到该元素\n");
        return ERROR;
    }

    *pe = L->arr[index - 2];
    return OK;

}

// 查询后继
Status SqListNextElem(SqList *L, ElemType e, ElemType *ne){

    if (L->length <= 1) {
        printf("顺序表元素不足，无法查询\n");
        return ERROR;
    }

    if (L->arr[L->length - 1] == e) {
        printf("最后一个元素无后继\n");
        return ERROR;
    }

    // 查找 e 所在位置
    int index;
    int loacateResult;
    loacateResult = SqListLocateElem(L, e, &index);
    if (loacateResult != TRUE) {
        printf("未找到该元素\n");
        return ERROR;
    }

    *ne = L->arr[index];
    return OK;

}

// 销毁
void SqListDestroy(SqList *L) {
    free(L->arr);
    L->arr = NULL;
    L->length = 0;
    L->size = 0;
}

// 清空表：即重置为空表
Status SqListClear(SqList *L) {
    L->arr = (ElemType *)malloc(L->size * sizeof(ElemType));
    if (!L->arr)  {            // 存储分配失败
        printf("内存分配失败\n");
        return ERROR;     
    }
    L->length = 0;
    return OK;
}