#include <stdio.h>
#include "DynamicArray.c"

// 测试
int main(){

    // 初始化一个数组
    DynamicArray arr;
    int res = newDynamicArray(&arr);
    if(res == -1){
        return 0;
    }
    display(&arr);

    // 测试插入
    insert(&arr, 10, 0);
    insert(&arr, 11, 1);
    insert(&arr, 12, 2);
    insert(&arr, 13, 3);
    insert(&arr, 14, 4);
    display(&arr);

    // 测试扩容
    insert(&arr, 15, 5);
    insert(&arr, 16, 6);
    display(&arr);
    printf("容量:%d\n", capacity(&arr));

    // 测试删除
    delete(&arr, 0);
    display(&arr);
    delete(&arr, 2);
    display(&arr);

    // 测试缩容
    printf("容量:%d\n", capacity(&arr));
    delete(&arr, 4);
    delete(&arr, 0);
    delete(&arr, 0);
    delete(&arr, 0);
    display(&arr);
    printf("容量:%d\n", capacity(&arr));

    // 测试更新
    update(&arr, 100, 0);

    // 测试查找
    printf("位置：%d\n", locate(&arr, 100));

    // 测试取值
    int e = 100;
    get(&arr, 0, &e);
    printf("e:%d\n", e);

    return 0;
}