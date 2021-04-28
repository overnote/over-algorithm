/**
 * 动态数组
 */

#define MaxSize 5               // 默认容量，取小值便于测试
typedef int DataType;           // 数据元素类型

// 动态数组结构体
typedef struct {                // 定义的类型
    // DataType data[MaxSize];  // 静态分配方式，即定义时已经确定了数组空间（长度）
    DataType    *data;          // 动态数组数据元素存储地址，动态分配方式
    int         capacity;       // 动态数组容量
    int         length;         // 当前数组长度
} DynamicArray;                 // 给该类型起的名字

// 构造空动态数组
DynamicArray newDynamicArray();

// 扩容工具：注意该函数放置顺序
void expandCap(DynamicArray *A);

// 缩容工具：注意该函数放置顺序，数组元素为当前容量的 1/4，缩容为当前容量的一半
void reduceCap(DynamicArray *A);

// 删:移除最后一位
void popElem(DynamicArray *A);

// 查：根据索引获取数据
DataType indexElem(DynamicArray *A, int index);

// 获取容量：
int Capacity(DynamicArray *A);

// 获取长度：
int Length(DynamicArray *A);

// 清空
void clear(DynamicArray *A);

// 销毁
void destroy(DynamicArray *A);