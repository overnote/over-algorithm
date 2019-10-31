#include "../list/LinkList.c"

void test_LinkList() {

    printf("********** 单链表测试开始 **********\n");

    struct LinkList L;
    LinkListInit(&L);

    printf("追加元素-------------\n");
    LinkListAppend(&L, 17);
    LinkListAppend(&L, 19);
    LinkListDisplay(&L);

    printf("插入元素-------------\n");
    LinkListInsert(&L, 2, 18);
    LinkListDisplay(&L);


    printf("********** 顺序表测试结束 **********\n");
    
}