/**
 * 单链表
 */

typedef int DataType;

typedef int DataType;
typedef struct Node{
    DataType        data;       // 结点存储的书
    struct Node     *next;      // 结点指针
}*Node, LinkList;

Node newNode(DataType data);
LinkList newLinkList();
int insert(LinkList *L, DataType e, int index);
int delete(LinkList *L, int index, DataType *e);
void update(LinkList *L, int index, DataType e);
Node search(LinkList *L, DataType e);
Node locate(LinkList *L, int index);
int length(LinkList *L);
void clear(LinkList *L);
void destroy(LinkList *L);
void display(LinkList *L);