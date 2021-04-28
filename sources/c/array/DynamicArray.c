/**
 * 动态数组
 */

#include <stdio.h>
#include <stdlib.h>

#include "DynamicArray.h"

// 构造空动态数组
DynamicArray newDynamicArray(){
    // 申请内存
    DynamicArray *A = malloc(sizeof(DynamicArray));
    if(A == NULL){
        printf("DynamicArray 内存申请失败");
        exit(1);
    }
    DataType *newData = (DataType *)malloc(MaxSize * sizeof(DataType));
    if(newData == NULL){
        printf("DynamicArray 内部结构内存申请失败");
        exit(1);
    }
    A->data = newData;

    // 修正属性
    A->capacity = MaxSize;
    A->length = 0;

    return *A;
}

// 扩容工具：注意该函数放置顺序
void expandCap(DynamicArray *A){
    if(A->length == A->capacity){
        // 执行扩容：扩充为原容量的二倍
        int newCap = A->capacity * 2;

        // 申请新数组空间，并拷贝数据：使用最原始的for循环，尽量简单易理解
        // realloc 方式：(DataType *)realloc(A->data, newCap * sizeof(DataType)); 
        DataType *newData = malloc(newCap * sizeof(DataType));
        if(newData == NULL){
            printf("DynamicArray 内部结构内存申请失败");
            exit(1);
        }
        for(int i = 0; i < A->length; i++){
            newData[i] = A->data[i];
        }
        free(A->data);
        A->data = newData;

        // 变更新容量
        A->capacity = newCap;
    }
}

// 缩容工具：注意该函数放置顺序，数组元素为当前容量的 1/4，缩容为当前容量的一半
void reduceCap(DynamicArray *A){
    if(A->length <= A->capacity / 4){
        int newCap = A->capacity / 2;

        // 缩容
        DataType *newData = malloc(newCap * sizeof(DataType));
        if(newData == NULL){
            printf("DynamicArray 内部结构内存申请失败");
            exit(1);
        }
        for(int i = 0; i < A->length; i++){
            newData[i] = A->data[i];
        }
        free(A->data);
        A->data = newData;

        // 变更新容量
        A->capacity = newCap;
    }
}

// 增：增加最后一位
void pushElem(DynamicArray *A, DataType e){
    if(A->data == NULL){
        printf("数组不存在\n");
    } else {
        expandCap(A);
        A->length++;
        A->data[A->length - 1] = e;
    }
}

// 删:移除最后一位
void popElem(DynamicArray *A){
    if(A->length == 0){
        printf("数组元素个数为0\n");
    } else {
        reduceCap(A);
        A->length--;
    }
}

// 查：根据索引获取数据
DataType indexElem(DynamicArray *A, int index){
    if(index < 0 || index > A->length - 1){
        printf("数组索引越界\n");
        return (DataType)0;
    } else {
        return A->data[index];
    }
}

// 获取容量：
int Capacity(DynamicArray *A){
    return A->capacity;
}

// 获取长度：
int Length(DynamicArray *A){
    return A->length;
}

// 清空
void clear(DynamicArray *A){
    A->length = 0;
}

// 销毁
void destroy(DynamicArray *A){
    if(A->data != NULL){
        free(A->data);
    }
}

// 显示动态数组数据
void display(DynamicArray *A){
    if(A->length == 0){
        printf("[]\n");
    } else {
        printf("[");
        for(int i = 0; i < A->length; i++){ 
            if(i == A->length - 1){
                printf("%d", A->data[i]);
            } else {
                printf("%d, ", A->data[i]);
            }
        }
        printf("]\n");
    }
}

// 测试
int main(){

    // 初始化一个数组
    DynamicArray arr = newDynamicArray();
    display(&arr);

    pushElem(&arr, 1);
    pushElem(&arr, 2);
    pushElem(&arr, 3);
    pushElem(&arr, 4);
    pushElem(&arr, 5);
    printf("arr[4]=%d\n", indexElem(&arr, 4));
    display(&arr);

    pushElem(&arr, 6);
    pushElem(&arr, 7);
    display(&arr);

    popElem(&arr);
    popElem(&arr);
    popElem(&arr);
    popElem(&arr);
    popElem(&arr);
    display(&arr);

    destroy(&arr);
    return 0;
}