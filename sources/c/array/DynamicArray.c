/**
 * 可变数组
 */

#include <stdio.h>
#include <stdlib.h>
#include "DynamicArray.h"

// 构造空可变数组
int newDynamicArray(DynamicArray *A){
    // 申请内部动态数组内存
    A->data = (DataType *)malloc(MaxSize * sizeof(DataType));
    if(A->data == NULL){
        printf("内存申请失败");
        return -1;
    }

    // 修正属性
    A->capacity = MaxSize;
    A->length = 0;

    return 1;
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
int search(DynamicArray *A, DataType e){
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

// 测试
int main(){

    // 初始化一个数组
    DynamicArray arr;
    int res = newDynamicArray(&arr);
    if(res == -1){
        return 0;
    }
    display(&arr);

    // 测试插入
    insert(&arr, 10, 0);
    insert(&arr, 11, 1);
    insert(&arr, 12, 2);
    insert(&arr, 13, 3);
    insert(&arr, 14, 4);
    display(&arr);

    // 测试扩容
    insert(&arr, 15, 5);
    insert(&arr, 16, 6);
    display(&arr);
    printf("容量:%d\n", capacity(&arr));

    // 测试删除
    delete(&arr, 0);
    display(&arr);
    delete(&arr, 2);
    display(&arr);

    // 测试缩容
    printf("容量:%d\n", capacity(&arr));
    delete(&arr, 4);
    delete(&arr, 0);
    delete(&arr, 0);
    delete(&arr, 0);
    display(&arr);
    printf("容量:%d\n", capacity(&arr));

    // 测试更新
    update(&arr, 100, 0);

    // 测试查找
    printf("位置：%d\n", search(&arr, 100));

    // 测试取值
    int e = 100;
    get(&arr, 0, &e);
    printf("e:%d\n", e);

    return 0;
}