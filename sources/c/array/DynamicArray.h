/**
 * 可变数组
 */

#define MaxSize 5           // 默认容量，取小值便于测试
typedef int DataType;       // 数据元素类型

// 可变数组结构体
typedef struct{ 
    DataType    *data;      // 可变数组数据元素存储地址，可变分配方式
    int         capacity;   // 可变数组容量
    int         length;     // 当前数组长度
} DynamicArray;

int newDynamicArray(DynamicArray *A);                   // 构造
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