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
    SqListDisplay(&L);
    printf("%d\n", L.length);
    printf("********** 顺序表测试结束 **********\n");
}