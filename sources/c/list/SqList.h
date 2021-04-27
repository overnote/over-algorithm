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

// 构造空顺序表
SqList newSqList();

// 增：按照一个位置插入数据
int insert(SqList *L, DataType e, int index);

// 删：delete
int delete(SqList *L, int index);

// 查：根据位置查询值
DataType searchValue(SqList *L, int index);

// 查：根据值查询位置
int searchIndex(SqList *L, DataType e);

// 获取顺序表最大容量
int SqListSize(SqList *L);

// 清空
void clear(SqList *L);

// 销毁
void destroy(SqList *L);

