/**
 * 双向循环链表
 */

typedef int DataType;

// 经典教材写法
// typedef struct Node {
//     DataType        data;
//     struct Node     *prev;
//     struct Node     *next;
// }  Node, *DblList;

// 笔者的写法
typedef struct Node {
    DataType        data;       // 数据域
    struct Node     *prev;      // 前驱指针
    struct Node     *next;      // 后继指针
} Node;

typedef struct {
    Node            *head;      // 头指针
    int             size;       // 元素个数：约定不包括头节点
} DblList;

Node* newNode(DataType data);
DblList* newDblList();
int insert(DblList *L, DataType e, int index);
int delete(DblList *L, int index, DataType *e);
void update(DblList *L, int index, DataType e);
Node* search(DblList *L, DataType e);
Node* locate(DblList *L, int index);
Node* prevElem(DblList *L, DataType e); // 根据值快速找到前驱
Node* nextElem(DblList *L, DataType e); // 根据值快速找到后继
int length(DblList *L);
void clear(DblList *L);
void destroy(DblList *L);
void display(DblList *L);