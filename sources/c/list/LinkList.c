/**
 * 单链表
 */

#include <stdio.h>
#include <stdlib.h>
#include "LinkList.h"

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

// 增：插入节点
// 约定：带头节点的链表，插入时，只能在头节点之后插入，也不允许插入超过最大元素个数的位置
int insert(LinkList *L, DataType e, int index){

    // 找到其前一个元素位置
    Node p = locate(L, index - 1);
    if(p == NULL){
        return -1;
    }

    // 创建要插入的节点
    Node temp = newNode(e);
    temp->next = p->next;
    p->next = temp;

    L->data++;  // 不要忘记存储的长度+1
    return 1;
}

// 删：根据位置删除，返回被删除的元素
int delete(LinkList *L, int index, DataType *e){
    // 找到其前一个元素位置
    Node p = locate(L, index - 1);
    if(p == NULL || p->next == NULL){
        return -1;
    }

    // 执行删除
    Node temp = p->next;
    *e = temp->data;
    p->next = p->next->next;
    free(temp);

    L->data--;  // 不要忘记存储的长度-1
    return 0;
}

// 改
void update(LinkList *L, int index, DataType e){
    Node p = locate(L, index);
    if(p == NULL){
        return;
    }
    p->data = e;
}

// 查：根据值查询节点地址
Node search(LinkList *L, DataType e){
    Node p = L->next;
    while(p != NULL && p->data != e){
        p =p->next;
    }
    return p;
}

// 查：根据位置查询节点地址
Node locate(LinkList *L, int index){
    if(index < 0 || index > L->data + 1 ){
        printf("获取位置不合法\n");
        return NULL;
    }

    int k = 0;
    Node p = L;
    while(p != NULL && k < index){
        p = p->next;
        k++;
    }
    return p;
}

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
    L->data = 0;
}

// 销毁表
void destroy(LinkList *L){
    while(L->next != NULL){
        Node temp = L->next;
        L->next = temp->next;
        free(temp);
    }
    free(L);
}

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