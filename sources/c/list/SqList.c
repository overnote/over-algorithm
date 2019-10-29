#include <stdio.h>
#include <stdlib.h>

#define TRUE 1                  
#define FALSE 0
#define ERROR -1
#define NOTFOUND -2
#define UNEXEC -3
typedef int Status;

#define LIST_MAX_SIZE 100       // 顺序表存储空间的初始分配量
#define LIST_INCREMENT 10       // 顺序表存储空间的分配增量

typedef int ElemType;           // 定义数据元素类型为int，不直接使用int的原因是可以随时更改
typedef struct SqList{
    ElemType    *elem;           // 存储空间基址，即一个动态数组
    int         length;         // 当前长度
    int         size;           // 当前分配的存储容量，以 sizeof(ElemType)为单位
}SqList;

// 构造一个空的顺序表L
void SqListInit(SqList *L) {
    L->elem = (ElemType *)malloc(LIST_MAX_SIZE * sizeof(ElemType));
    if (!L->elem)  {            // 存储分配失败
        printf("内存分配失败");
        exit(ERROR);         
    }
    L->length = 0;
    L->size = LIST_MAX_SIZE;
}

// 获取线性表长度
int SqListLength(SqList *L) {
    return L->length;
}

// 显示
void SqListDisplay(SqList *L) {
    if (L->length == 0) {
        printf("数据结构无长度");
        exit(UNEXEC);
    }
    int i;
    for (i = 0; i < L->length; i++) {
        prinft("item[%d]=%d\n", i, L->elem[i]);
    }
}

// 插入
void SqListInsert(SqList *L, int index, ElemType e) {
    int i ;
    for (i = 0; i < L->length; i++) {
        if
    }
}

// 删除  按照位序删除

// 修改  按照位序修改

// 查询  按照位序查询值
ElemType SqListGetElem(SqList *L, int index) {
    if (index < 1 || index > L->length) {
        printf("索引越界");
        exit(ERROR);
    }
    return L->elem[index];
}

// 查询  按照值查询位序
int SqListLocateElem(SqList *L,){

}


// 查询前驱
// ElemType Pro


// 查询后继



// 销毁
void SqListDestroy(SqList *L) {
    free(L->elem);
    L->elem = NULL;
    L->length = 0;
    L->size = 0;
}

// 清空表：即重置为空表
void ClearList(SqList *L) {
    L->length = 0;
}

// 判断 是否是空表
Status SqListEmpty(SqList *L) {
    if (L->length == 0) {
        return TRUE;
    }
    return FALSE;
}

