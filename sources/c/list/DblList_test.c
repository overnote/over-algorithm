#include <stdio.h>
#include "DblList.c"

int main(){
    DblList *L = newDblList();
    insert(L, 11, 1);
    insert(L, 12, 2);
    insert(L, 13, 3);
    insert(L, 14, 4);
    insert(L, 15, 5);
    insert(L, 33, 3);
    display(L);

    DataType e1;
    delete(L, 5, &e1);
    printf("删除元素：%d\n", e1);
    display(L);
    DataType e2;
    delete(L, 3, &e2);
    printf("删除元素：%d\n", e2);
    display(L);

    update(L, 1, 21);
    display(L);
    printf("查询元素:21，结果：%d\n", search(L, 21)->data);
    printf("查询前驱:21，结果：%d\n", prevElem(L, 21)->data);
    printf("查询后继:21，结果：%d\n", nextElem(L, 21)->data);

    update(L, 4, 44);
    display(L);
    printf("查询元素:44，结果：%d\n", search(L, 44)->data);
    printf("查询前驱:44，结果：%d\n", prevElem(L, 44)->data);
    printf("查询后继:44，结果：%d\n", nextElem(L, 44)->data);

    clear(L);
    display(L);

    return 0;
}