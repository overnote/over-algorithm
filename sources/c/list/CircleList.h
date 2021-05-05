/**
 * 循环链表
 */

typedef int DataType;

// 经典教材写法
// typedef struct Node {
//     DataType            data;   // 结点数据
//     struct Node         *next;  // 指向下一个结点的指针
// } Node, *CircleList;

typedef struct Node{
    DataType        data;       // 结点存储的书
    struct Node     *next;      // 结点指针
} Node;

typedef struct{
    Node            *head;      // 头指针
    int             length;     // 元素个数：约定这里排除头结点
} CircleList;


Node* newNode(DataType data);
CircleList* newCircleList();
int insert(CircleList *L, DataType e, int index);
int delete(CircleList *L, int index, DataType *e);
void update(CircleList *L, int index, DataType e);
Node* search(CircleList *L, DataType e);
Node* locate(CircleList *L, int index);
int length(CircleList *L);
void clear(CircleList *L);
void destroy(CircleList *L);
void display(CircleList *L);