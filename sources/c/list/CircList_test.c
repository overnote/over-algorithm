#include <stdio.h>
#include "CircList.c"

int main(){
    CircList L = newCircList();
    insert(&L, 11, 1);
    insert(&L, 12, 2);
    insert(&L, 13, 3);
    insert(&L, 14, 4);
    insert(&L, 15, 5);
    // display(&L);

    printf("p0=%d\n", L.data);
    printf("p1=%d\n", L.next->data);
    printf("p2=%d\n", L.next->next->data);
    printf("p3=%d\n", L.next->next->next->data);
    printf("p4=%d\n", L.next->next->next->next->data);
    printf("p5=%d\n", L.next->next->next->next->next->data);

    printf("p6=%d\n", L.next->next->next->next->next->next->data);
    printf("p7=%d\n", L.next->next->next->next->next->next->next->data);
    printf("p8=%d\n", L.next->next->next->next->next->next->next->next->data);
    // DataType e1;
    // delete(&L, 5, &e1);
    // printf("删除元素：%d\n", e1);
    // display(&L);
    // DataType e2;
    // delete(&L, 3, &e2);
    // printf("删除元素：%d\n", e2);
    // display(&L);

    // update(&L, 3, 33);
    // display(&L);

    // clear(&L);
    // display(&L);
    return 0;
}