/**
 * 顺序表
 */
#define MAX_SIZE 5 // 存储空间的初始分配量，取小值便于测试
typedef int
    ElemType; // 定义数据元素类型为int，不直接使用int的原因是可以随时更改

// 顺序表结构体
typedef struct {
    ElemType *data; // 存储空间地址，即一个动态数组
    int size;       // 当前长度，即数据元素个数
    int capacity; // 当前分配的存储容量，以 sizeof(ElemType)为单位
} SqList;

SqList *newSqList();
int insert(SqList *L, ElemType e, int index);
int delete(SqList *L, int index, ElemType *e);
void update(SqList *L, ElemType e, int index);
int search(SqList *L, ElemType *e, int index);
int locate(SqList *L, ElemType e);
int size(SqList *L);
int length(SqList *L);
void clear(SqList *L);
void destroy(SqList *L);
void display(SqList *L);
