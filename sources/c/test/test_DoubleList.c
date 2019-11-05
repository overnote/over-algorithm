#include "../list/DoubleList.c"

void test_DoubleList() {

    printf("********** 测试开始 **********\n");

    struct DoubleList L;
    DoubleListInit(&L);
    DoubleListDisplay(&L);

    printf("追加元素-------------\n");
    DoubleListAppend(&L, 17);
    DoubleListAppend(&L, 20);
    DoubleListDisplay(&L);

    printf("插入元素-------------\n");
    DoubleListInsert(&L, 2, 18);
    DoubleListInsert(&L, 3, 19);
    DoubleListDisplay(&L);

    printf("删除元素-------------\n");
    DoubleListDelete(&L, 4);
    DoubleListDisplay(&L);

    printf("修改元素-------------\n");
    DoubleListUpdate(&L, 2, 22);
    DoubleListDisplay(&L);

    printf("索引查找值-------------\n");
    int ge;
    DoubleListGetElem(&L, 3, &ge);
    printf("索引值：%d\n", ge);

    printf("值查找索引-------------\n");
    int le;
    DoubleListLocateElem(&L, 22, &le);
    printf("索引值：%d\n", le);

    printf("查找前驱---------------\n");
    int pe;
    DoubleListPriorElem(&L, 17, &pe);
    printf("元素前驱：%d\n", pe);

    printf("查找后继---------------\n");
    int ne;
    int r = DoubleListNextElem(&L, 22, &ne);
    printf("元素后继：%d\n", ne);
    printf("查找结果：%d,\n", r);

    printf("清空表---------------\n");
    DoubleListClear(&L);
    DoubleListDisplay(&L);

    printf("删除表---------------\n");
    DoubleListDestroy(&L);

    printf("********** 测试结束 **********\n");
    
}