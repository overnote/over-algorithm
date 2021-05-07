/**
 * 静态链表
 */

typedef int DataType;           
#define MAX_SIZE 5             

// 静态链表结点
typedef struct Node{
    DataType            data;  
    int                 cur;    // 游标：类似于next指针
} Node;

// 顺序表结构体
typedef struct{
    Node        *head;          // 整个数组，也是备用链表的表头          
    int         size;           // 数组元素个数         
    int         capacity;       // 数组容量
} StaticList;

Node* newNode(DataType data, int cur);
StaticList* newStaticList();
int insert(StaticList *L, DataType e, int index);
int delete(StaticList *L, int index);
void display(StaticList *L);
