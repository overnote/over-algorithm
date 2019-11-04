#include "../list/LinkList.c"

void test_LinkList() {

    printf("********** 单链表测试开始 **********\n");

    struct LinkList L;
    LinkListInit(&L);

    printf("追加元素-------------\n");
    LinkListAppend(&L, 17);
    LinkListAppend(&L, 20);
    LinkListDisplay(&L);

    printf("插入元素-------------\n");
    LinkListInsert(&L, 2, 18);
    LinkListInsert(&L, 3, 19);
    LinkListDisplay(&L);

    printf("修改元素-------------\n");
    LinkListUpdate(&L, 2, 22);
    LinkListDisplay(&L);

    printf("查找元素-------------\n");
    int te;
    int pe;
    int ne;
    LinkListLocateElem(&L, 19, &te);
    LinkListPriorElem(&L, 22, &pe);
    LinkListNextElem(&L, 22, &ne);
    printf("查找第三个元素位置：%d\n", te);
    printf("查找第二个元素前驱：%d\n", pe);
    printf("查找第二个元素后继：%d\n", ne);
    
    printf("删除元素-------------\n");
    LinkListDelete(&L, 2);
    LinkListDisplay(&L);

    printf("清空表---------------\n");
    LinkListClear(&L);
    LinkListDisplay(&L);

    printf("删除表---------------\n");
    LinkListDestroy(&L);

    printf("********** 顺序表测试结束 **********\n");
    
}