/**
 * 顺序表
 */

#include <stdio.h>
#include <stdlib.h>
#include "SqList.h"

// 构造空顺序表
SqList newSqList() {
    // 申请内存
    SqList *L = malloc(sizeof(SqList));
    if(L == NULL){
        printf("SqList 内存申请未成功");
        exit(1);    // 1 表示异常退出
    }
    DataType *newData = (DataType *)malloc(MAX_SIZE * sizeof(DataType));
    if(newData == NULL){
        printf("SqList 内部空间内存申请未成功");
        exit(1);    // 1 表示异常退出
    }
    L->data = newData;

    // 初始化数据
    L->length = 0;
    L->size = MAX_SIZE;
    return *L;
}

// 增：按照一个位置插入数据
int insert(SqList *L, DataType e, int index){
    if(index < 1 || index > L->length + 1){
        printf("插入位置不合法\n");
        return -1;
    }
    if(L->length == L->size){
        printf("顺序表容量已满\n");
        return -1;
    }
    
    // 移动元素
    for(int i = L->length; i >= index; i--){
        L->data[i] = L->data[i - 1];
    }

    // 插入位置值改变
    L->data[index-1] = e;
    L->length++;
    return 1;
}

// 删：删除位置上的元素，并获取到删除元素的值
int delete(SqList *L, int index, DataType *e){
    if(index < 1 || index > L->length){
        printf("删除位置不合法\n");
        return -1;
    }
    if(L->length == 0){
        printf("空表无元素可移除\n");
        return -1;
    }

    *e = L->data[index - 1];

    for(int i = index; i < L->length; i++){
        L->data[i - 1] = L->data[i];
    }
    L->length--;
    return 1;
}

// 改
void update(SqList *L, int index, DataType e){
    if(index < 1 || index > L->length){
        printf("修改位置不合法\n");
        return;
    }
    L->data[index - 1] = e;
}

// 查：根据位置查询值
int search(SqList *L, DataType *e, int index){
    if(index < 1 || index > L->length){
        printf("参数位置不合法\n");
        return -1;  // 查找失败，返回-1
    }
    *e = L->data[index - 1];
    return 1;
}

// 查：根据值查询位置
int locate(SqList *L, DataType e){
    for(int i = 0; i < L->length; i++){
        if(L->data[i] == e){
            return i;
        }
    }
    printf("未找到该元素\n");
    return -1;  // 未找到
}

// 获取顺序表最大容量
int size(SqList *L) {
    return L->size;
}

// 获取顺序表长度
int length(SqList *L) {
    return L->length;
}

// 清空
void clear(SqList *L){
    L->length = 0;
}

// 销毁
void destroy(SqList *L){
    if(L->data != NULL){
        free(L->data);
    }
}

// 显示动态数组数据
void display(SqList *L){
    if(L->length == 0){
        printf("[]\n");
        return;
    }

    printf("[");
    for(int i = 0; i < L->length; i++){ 
        if(i == L->length - 1){
            printf("%d", L->data[i]);
        } else {
            printf("%d, ", L->data[i]);
        }
    }
    printf("]\n");
}
