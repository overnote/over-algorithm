#include <stdio.h>
#include <stdlib.h>
#include "SqList.h"

// 构造空顺序表
SqList newSqList() {
    // 申请内存
    SqList *L = malloc(sizeof(SqList));
    L->data = (DataType *)malloc(MAX_SIZE * sizeof(DataType));

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

// 删：delete
int delete(SqList *L, int index){
    if(index < 1 || index > L->length){
        printf("删除位置不合法\n");
        return -1;
    }
    if(L->length == 0){
        printf("空表无元素可移除\n");
        return -1;
    }
    for(int i = index; i < L->length; i++){
        L->data[i - 1] = L->data[i];
    }
    L->length--;
    return 1;
}

// 查：根据位置查询值
DataType searchValue(SqList *L, int index){
    if(index < 1 || index > L->length){
        printf("参数位置不合法\n");
        return -1;  // 查找失败，返回-1，囧
    } 
    
    return L->data[index - 1];
}

// 查：根据值查询位置
int searchIndex(SqList *L, DataType e){
    for(int i = 0; i < L->length; i++){
        if(L->data[i] == e){
            return i;
        }
    }
    printf("未找到该元素\n");
    return -1;  // 未找到
}

// 获取顺序表长度
int SqListLength(SqList *L) {
    return L->length;
}

// 获取顺序表最大容量
int SqListSize(SqList *L) {
    return L->size;
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

// 测试
int main(){
    // 初始化一个顺序表
    SqList L = newSqList();
    display(&L);

    insert(&L, 11, 1);
    insert(&L, 12, 2);
    insert(&L, 17, 7);
    display(&L);
    
    insert(&L, 13, 3);
    insert(&L, 14, 4);
    insert(&L, 15, 5);
    display(&L);

    insert(&L, 16, 6);
    insert(&L, 15, 5);

    delete(&L, 3);
    display(&L);

    return 0;
}