#include <stdio.h>
#include <stdlib.h>

typedef int DataType;           // 定义数据元素类型为int，不直接使用int的原因是可以随时更改
#define MAX_SIZE 5              // 存储空间的初始分配量，取小值便于测试


// 顺序表结构体
typedef struct{
    DataType    *data;          // 存储空间地址，即一个动态数组
    int         length;         // 当前长度，即数据元素个数
    int         size;           // 当前分配的存储容量，以 sizeof(ElemType)为单位
}SqList;

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
int insert(SqList *L, int index, DataType e){
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
        printf("()\n");
    } else {
        printf("(");
        for(int i = 0; i < L->length; i++){ 
            if(i == L->length - 1){
                printf("%d", L->data[i]);
            } else {
                printf("%d, ", L->data[i]);
            }
        }
        printf(")\n");
    }
}

int main(){

    SqList L = newSqList();
    display(&L);

    insert(&L, 1, 11);
    insert(&L, 2, 12);
    insert(&L, 3, 13);
    insert(&L, 4, 14);
    insert(&L, 5, 15);
    display(&L);

    insert(&L, 6, 16);
    insert(&L, 5, 155);

    delete(&L, 3);
    display(&L);

    return 0;
}