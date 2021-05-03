#include <stdio.h>
#include "SqList.c"

int main(){
    // 初始化一个顺序表
    SqList L = newSqList();
    display(&L);

    insert(&L, 11, 1);
    insert(&L, 12, 2);    
    insert(&L, 13, 3);
    insert(&L, 14, 4);
    insert(&L, 15, 5);
    insert(&L, 17, 7);
    insert(&L, 15, 5);
    display(&L);

    int e1;
    delete(&L, 3, &e1);
    printf("删除元素：%d\n", e1);
    display(&L);

    update(&L, 1, 100);
    display(&L);

    int e2;
    search(&L, &e2, 2);
    printf("查询元素：%d\n", e2);

    printf("查询元素：%d\n", locate(&L, 100));

    return 0;
}