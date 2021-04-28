/**
 * 单链表
 */

#include <stdio.h>
#include <stdlib.h>

typedef int DataType;

typedef int DataType;
typedef struct Node{
    DataType        data;       // 结点存储的书
    struct Node     *next;      // 结点指针
}*Node, LinkList;

// 创建节点
Node newNode(DataType data){
    Node node = malloc(sizeof(Node));
    if(node == NULL){
        printf("申请结点内存失败");
        exit(1);
    }
    node->data = data;
    node->next = NULL;
    return node;
}

// 创建单链表
LinkList newLinkList(){
    // 申请链表内存：其本身其实就是创建了一个头节点
    LinkList *L = newNode(0);
    return *L;
}

// 增：插入节点。笔者这里约定：带头节点的链表，插入时，只能在头节点之后插入，也允许插入超过最大元素个数的位置
int insert(LinkList *L, DataType e, int index){
    if(index < 1 || index > L->data + 1){
        printf("插入位置不合法\n");
        return -1;
    }
    
    // 创建要插入的节点
    Node insertNode = newNode(e);
    // 当前循环到的节点：要先找到插入位置前一个节点(index - 1)
    Node currentNode = L;
    // 没有使用 while(currentNode->next != NULL) 遍历原因：当首次插入元素时，while不执行，需要额外判断一次
    for(int i = 0; i <= index - 1; i++){
        if(i == index - 1){
            insertNode->next = currentNode->next;
            currentNode->next = insertNode;
        } else {
            currentNode = currentNode->next;
        }
    }

    L->data++;  // 不要忘记存储的长度+1
    return 1;
}

// 查：根据值查询元素位置
int locate(){
    return 0;
}

// 查：根据位置查询元素值
DataType search(){

}

// 删：根据位置删除

// 获取表长度
int length(LinkList *L){
    return L->data; // 如果没有头节点一般使用循环获取长度
}

// 清空表:仅保留头节点
void clear(LinkList *L){
    while(L->next != NULL){
        Node temp = L->next;
        L->next = temp->next;
        free(temp);
    }
    // L->data = 0;
}

// 销毁表

// 显示单链表
void display(LinkList *L){
    if(L->data == 0){
        printf("空链表\n");
        return;
    }

    Node currentNode = L;
    int pos = 0;
    while(currentNode != NULL){
        if(pos == L->data){    // 最后一位
            printf("%d\n", currentNode->data);
            break;
        } else {
            printf("%d->", currentNode->data);
            currentNode = currentNode->next;
            pos++;
        }
    }
}

int main(){

    LinkList L = newLinkList();
    insert(&L, 11, 1);
    insert(&L, 12, 2);
    insert(&L, 13, 3);
    display(&L);

    clear(&L);
    display(&L);
    return 0;
}