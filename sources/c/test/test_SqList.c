#include "../list/SqList.c"

void test_SqList() {

    printf("********** 顺序表测试开始 **********\n");

    struct SqList L;
    SqListInit(&L);
    SqListDisplay(&L);

    printf("增加元素-------------\n");
    SqListInsert(&L, 1, 17);
    SqListInsert(&L, 2, 18);
    SqListInsert(&L, 3, 19);
    printf("长度 %d\n", SqListLength(L));
    SqListDisplay(&L);

    printf("修改元素-------------\n");
    SqListUpdate(&L, 2, 22);
    SqListDisplay(&L);

    int pe;
    int ne;
    SqListPriorElem(&L, 22, &pe);
    SqListNextElem(&L, 22, &ne);
    printf("查找第二个元素前驱：%d\n", pe);
    printf("查找第二个元素后继：%d\n", ne);
    
    printf("删除元素-------------\n");
    SqListDelete(&L, 2);
    SqListDisplay(&L);

    printf("********** 顺序表测试结束 **********\n");
    
}