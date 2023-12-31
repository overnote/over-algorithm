/**
 * 可变数组
 */

#include <stdio.h>
#include <stdlib.h>

#define MaxSize 5           // 默认容量，取小值便于测试
typedef int DataType;       // 数据元素类型

// 可变数组结构体
typedef struct{ 
    DataType    *data;      // 可变数组数据元素存储地址，可变分配方式
    int         capacity;   // 可变数组容量
    int         length;     // 当前数组长度
} DynamicArray;

DynamicArray* newDynamicArray();                        // 构造
void insert(DynamicArray *A, DataType e, int index);    // 增
void delete(DynamicArray *A, int index);                // 删
void update(DynamicArray *A, DataType e, int index);    // 改
int locate(DynamicArray *A, DataType e);                // 查
int get(DynamicArray *A, int index, DataType *e);       // 取
int capacity(DynamicArray *A);                          // 获取容量
int length(DynamicArray *A);                            // 获取长度
void clear(DynamicArray *A);                            // 清空
void destroy(DynamicArray *A);                          // 销毁

int expandCap(DynamicArray *A);                         // 扩容工具
int reduceCap(DynamicArray *A);                         // 缩容工具
void display(DynamicArray *A);                          // 显示工具

// 构造空可变数组
DynamicArray* newDynamicArray(){

    // 申请结构体内存
    DynamicArray *A = malloc(sizeof(DynamicArray));
    if(A == NULL){
        printf("内存申请失败");
        exit(1);    // 1 表示异常退出
    }

    // 申请内部动态数组内存
    A->data = (DataType *)malloc(MaxSize * sizeof(DataType));
    if(A->data == NULL){
        printf("内存申请失败");
        exit(1);    // 1 表示异常退出
    }

    // 修正属性
    A->capacity = MaxSize;
    A->length = 0;

    return A;
}

// 增：根据索引插入数据
void insert(DynamicArray *A, DataType e, int index){

    if(A->data == NULL){
        printf("可变数组结构不合法\n");
        return;
    }

    if(index < 0 || index > A->length){
        printf("插入位置不合法\n");
        return;
    }

    // 每次插入先执行扩容机制
    int res = expandCap(A);
    if(res != 1){
        return;
    }

    // 执行插入：从最后一位循环，循环到插入位置
    for(int i = A->length; i >= index; i--){
        if(i == index){
            A->data[index] = e;
        } else {
            A->data[A->length] = A->data[A->length - 1];
        }
    }

    // 更新长度
    A->length++;
}

// 删：按索引位置删除元素
void delete(DynamicArray *A, int index){
    if(index < 0 || index > A->length - 1){
        printf("删除位置不合法\n");
        return;
    }

    // 每次删除前执行缩容
    int res = reduceCap(A);
    if(res != 1){
        return;
    }

    for(int i = index; i < A->length; i++){
        A->data[i] = A->data[i + 1];
    }
    A->length--;
}

// 改：根据索引修改元素
void update(DynamicArray *A, DataType e, int index){
    if(index < 0 || index > A->length - 1){
        printf("数组索引越界\n");
        return;
    }
    A->data[index] = e;
} 

// 查：根据值查询索引
int locate(DynamicArray *A, DataType e){
    for(int i = 0; i < A->length; i++){
        if(A->data[i] == e){
            return i;
        }
    }
    return -1;
}

// 取：根据索引进行数据存取
int get(DynamicArray *A, int index, DataType *e){
    if(index < 0 || index > A->length - 1){
        printf("数组索引越界\n");
        return -1;
    } 
    
    *e = A->data[index];
    return 1;
}

// 获取容量
int capacity(DynamicArray *A){
    return A->capacity;
}

// 获取长度
int length(DynamicArray *A){
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

// 扩容工具：注意该函数放置顺序
int expandCap(DynamicArray *A){

    // 检查是否需要扩容
    if(A->length < A->capacity){
        return 1;
    }

    // 容量设置为原容量的二倍
    int newCap = A->capacity * 2;

    // 重新申请新的更大内存：可以使用for循环拷贝、realloc拷贝申请等方式
    DataType *newData = malloc(newCap * sizeof(DataType));
    if(newData == NULL){
        printf("内存申请失败");
        return -1;
    }
    for(int i = 0; i < A->length; i++){
        newData[i] = A->data[i];
    }

    // 释放旧数组内存，更换为新数组
    free(A->data);
    A->data = newData;

    // 更新容量
    A->capacity = newCap;
    return 1;
}

// 缩容工具：注意该函数放置顺序，数组元素为当前容量的 1/4，缩容为当前容量的一半
int reduceCap(DynamicArray *A){

    // 检查是否需要缩容
    if(A->length > A->capacity / 4){
        return 1;
    }

    // 容量设置为以前的一半
    int newCap = A->capacity / 2;

    // 重新申请更小的内存，并拷贝数据
    DataType *newData = malloc(newCap * sizeof(DataType));  // 
    if(newData == NULL){
        printf("内存申请失败");
        return -1;
    }
    for(int i = 0; i < A->length; i++){
        newData[i] = A->data[i];
    }

    // 释放旧数组内存，更换为新数组
    free(A->data);
    A->data = newData;

    // 更新容量
    A->capacity = newCap;
    return 1;
}

// 显示工具
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