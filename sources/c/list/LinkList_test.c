#include <stdio.h>
#include "LinkList.c"

int main(){
    LinkList L = newLinkList();
    insert(&L, 11, 1);
    insert(&L, 12, 2);
    insert(&L, 13, 3);
    insert(&L, 14, 4);
    insert(&L, 15, 5);
    display(&L);

    DataType e1;
    delete(&L, 5, &e1);
    printf("删除元素：%d\n", e1);
    display(&L);
    DataType e2;
    delete(&L, 3, &e2);
    printf("删除元素：%d\n", e2);
    display(&L);

    update(&L, 3, 33);
    display(&L);

    clear(&L);
    display(&L);

    return 0;
}