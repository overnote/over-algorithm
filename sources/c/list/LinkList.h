/**
 * 单链表
 */

typedef int DataType;

// 经典教材写法
// typedef struct Node {
//     DataType            data;   // 结点数据
//     struct Node         *next;  // 指向下一个结点的指针
// } Node, *LinkList;

// 笔者的写法：方便改造，如添加尾指针等等
// 额外的 length 为 int 类型，可以避免在头结点存储数据修改 DataType 造成的尴尬
typedef struct Node{
    DataType            data;   // 结点数据
    struct Node         *next;  // 指向下一个结点的指针
} Node;

typedef struct {
    struct Node         *head;
    int                 length; // 链表元素个数（不包括头结点）
} LinkList;


Node* newNode(DataType e);
LinkList* newLinkList();
int insert(LinkList *L, DataType e, int index);
int delete(LinkList *L, int index, DataType *e);
void update(LinkList *L, int index, DataType e);
Node* search(LinkList *L, DataType e);
Node* locate(LinkList *L, int index);
int length(LinkList *L);
void clear(LinkList *L);
void destroy(LinkList *L);
void display(LinkList *L);