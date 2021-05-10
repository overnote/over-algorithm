#include <stdio.h>
#include <stdlib.h>
#include "../list/CircleList.c"

// L 循环链表， cur 当前检测指针，m 最大出列序号， n 参与约瑟夫环的人数
void Josephus(CircleList *L, int m, int n){
    // 获取当前检测指针：一般这里直接认为都是从第一个有值节点开始
    Node *p = L->head->next;
    if(p == NULL){
        printf("链表数据有误\n");
        exit(1);
    }
    printf("从 %d 开始报数\n", p->data);
    
    Node *temp; // 循环到要删除结点的前一个结点
    for(int i = 0; i < n; i++){
        for(int j = 1; j < m; j++){
            temp = p;
            p = p->next;
            // 若为头结点，则跳过
            if(p == L->head){
                p = p->next;
            }
        }
        printf("%d\n", p->data);

        // 删除该结点
        temp->next = p->next;
        free(p);

        // 下一次报数从下一个结点开始
        p = temp->next;
        if(p == L->head){
            temp = p;
            p = p->next;
        }

    }
}

int main(){
    // 构建一个链表： n = 8， m = 3，依次出列顺序为：3 6 1 5 2 8 4
    CircleList *L = newCircleList();
    // ==TODO==可以使用尾插法、头插法等创建，这里暂时未制作
    insert(L, 1, 1);
    insert(L, 2, 2);
    insert(L, 3, 3);
    insert(L, 4, 4);
    insert(L, 5, 5);
    insert(L, 6, 6);
    insert(L, 7, 7);
    insert(L, 8, 8);
    display(L);

    Josephus(L, 3, 8);
    return 0;
}