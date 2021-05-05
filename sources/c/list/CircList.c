/**
 * 循环链表
 */

#include <stdio.h>
#include <stdlib.h>
#include "CircList.h"

// 创建结点
Node newNode(DataType data) {
    Node node = malloc(sizeof(Node));
    if(node == NULL){
        printf("申请结点内存失败\n");
        exit(1);
    }
    node->data = data;
    node->next = NULL;
    return node;
}

// 创建循环链表
CircList newCircList(){
    // 申请链表内存：其本身其实就是创建了一个头结点
    CircList *L = newNode(0);
    // 循环链表的重点：尾元素指向手原结点
    L->next = L;
    return *L;
}

// 增：插入结点
// 约定：插入时，只能在头结点之后插入，也不允许插入超过最大元素个数的位置
int insert(CircList *L, DataType e, int index){
    if(index < 1 || index > L->data + 1){
        printf("插入位置不合法\n");
        return -1;
    }

    // 获取插入位置前一个结点
    Node p = locate(L, index - 1);
    if(p == NULL){
        return -1;
    }
    printf("p->data=:%d\n", p->data);
    printf("p->next->data=:%d\n", p->next->data);

    // 执行插入
    Node temp = newNode(e);
    temp->next = p->next;
    p->next = temp;
    printf("temp.next->data=:%d\n", temp->next->data);

    L->data++;
    return 1;
}

// 删：根据位置删除，返回被删除的元素
int delete(CircList *L, int index, DataType *e){
    // 找到其前一个元素位置
    Node p = locate(L, index - 1);
    if(p == NULL || p->next == L){
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
void update(CircList *L, int index, DataType e){
    Node p = locate(L, index);
    if(p == NULL){
        return;
    }
    p->data = e;
}

// 查：根据值查询结点地址
Node search(CircList *L, DataType e){
    Node p = L->next;
    while(p != L && p->data != e){
        p = p->next;
    }
    return p;
}

// 定位
Node locate(CircList *L, int index){
    if(index < 0 || index > L->data + 1){
        printf("获取位置不合法\n");
        return NULL;
    }
    int k = 0;
    Node p = L;
    while(p->next != L && k < index){
        p = p->next;
        k++;
    }
    return p;
}

// 获取表长度
int length(CircList *L){
    return L->data; // 如果没有头结点一般使用循环获取长度
}

// 清空表:仅保留头结点
void clear(CircList *L){
    while(L->next != L){
        Node temp = L->next;
        printf("当前销毁：%d\n", temp->data);
        L->next = L->next->next;
        free(temp);
    }
    L->data = 0;
}

// 销毁表
void destroy(CircList *L){
    while(L->next != L){
        Node temp = L->next;
        L->next = temp->next;
        free(temp);
    }
    free(L);
}

// 显示循环链表
void display(CircList *L){
    if(L->data == 0){
        printf("空链表\n");
        return;
    }

    Node p = L;
    int pos = 0;
    while(p->next != L){
        // if(pos == L->data + 1){
        //     printf("%d->...\n", L->data);
        //     break;
        // }
        if(pos == 10){
            break;
        }
        printf("%d->", p->data);
        p = p->next;
        pos++;
    }
}