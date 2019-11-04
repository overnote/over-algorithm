#include "../list/LinkList.c"

void test_LinkList() {

    printf("********** 单链表测试开始 **********\n");

    struct LinkList L;
    LinkListInit(&L);
    LinkListDisplay(&L);

    printf("追加元素-------------\n");
    LinkListAppend(&L, 17);
    LinkListAppend(&L, 20);
    LinkListDisplay(&L);

    printf("插入元素-------------\n");
    LinkListInsert(&L, 2, 18);
    LinkListInsert(&L, 3, 19);
    LinkListDisplay(&L);

    printf("删除元素-------------\n");
    LinkListDelete(&L, 1);
    LinkListDisplay(&L);

    printf("修改元素-------------\n");
    LinkListUpdate(&L, 2, 22);
    LinkListDisplay(&L);

    printf("索引查找值-------------\n");
    int ge;
    LinkListGetElem(&L, 3, &ge);
    printf("索引值：%d\n", ge);

    printf("值查找索引-------------\n");
    int le;
    LinkListLocateElem(&L, 20, &le);
    printf("索引值：%d\n", le);

    printf("查找前驱---------------\n");
    int pe;
    LinkListPriorElem(&L, 17, &pe);
    printf("元素前驱：%d\n", pe);

    printf("查找后继---------------\n");
    int ne;
    int r = LinkListNextElem(&L, 22, &ne);
    printf("元素后继：%d\n", ne);
    printf("查找结果：%d,\n", r);

    printf("清空表---------------\n");
    LinkListClear(&L);
    // LinkListDisplay(&L);

    // printf("删除表---------------\n");
    // LinkListDestroy(&L);

    printf("********** 顺序表测试结束 **********\n");
    
}