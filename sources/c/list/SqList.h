/**
 * 顺序表
 */

typedef int DataType;           // 定义数据元素类型为int，不直接使用int的原因是可以随时更改
#define MAX_SIZE 5              // 存储空间的初始分配量，取小值便于测试

// 顺序表结构体
typedef struct{
    DataType    *data;          // 存储空间地址，即一个动态数组
    int         length;         // 当前长度，即数据元素个数
    int         size;           // 当前分配的存储容量，以 sizeof(ElemType)为单位
}SqList;

SqList newSqList();
int insert(SqList *L, DataType e, int index);
int delete(SqList *L, int index, DataType *e);
void update(SqList *L, DataType e, int index);
int search(SqList *L, DataType *e, int index);
int locate(SqList *L, DataType e);
int size(SqList *L);
int length(SqList *L);
void clear(SqList *L);
void destroy(SqList *L);
void diplay(SqList *L);

