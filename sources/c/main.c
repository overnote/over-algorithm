#include <stdio.h>
#include "./list/SqList.c"

void test_SqList();             // 测试顺序表

int main() {

    test_SqList();

    return 0;
}

void test_SqList() {
    printf("********** 顺序表测试开始 **********\n");
    struct SqList L;
    SqListInit(&L);
    printf("长度 %d\n", SqListLength(L));
    SqListInsert(&L, 1, 17);
    SqListInsert(&L, 2, 18);
    SqListInsert(&L, 3, 19);
    printf("长度 %d\n", SqListLength(L));
    SqListUpdate(&L, 2, 22);
    printf("修改 %d\n", L.elem[2]);
    printf("删除 %d\n", SqListDelete(&L, 2));
    printf("长度 %d\n", SqListLength(L));
    printf("********** 顺序表测试结束 **********\n");
}