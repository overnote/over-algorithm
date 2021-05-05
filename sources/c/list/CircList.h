/**
 * 循环链表
 */

typedef int DataType;

typedef struct{
    DataType        data;       // 结点存储的书
    struct Node     *next;      // 结点指针
}Node;

typedef struct{
    Node            *head;      // 头指针
    int             length;     // 元素个数：约定这里排除头结点
}CircList;


Node newNode(DataType data);
CircList newCircList();
int insert(CircList *L, DataType e, int index);
int delete(CircList *L, int index, DataType *e);
void update(CircList *L, int index, DataType e);
Node search(CircList *L, DataType e);
Node locate(CircList *L, int index);
int length(CircList *L);
void clear(CircList *L);
void destroy(CircList *L);
void display(CircList *L);