/**
 * 静态链表
 */
#include <stdio.h>
#include <stdlib.h>
#include "StaticList.h"

//将结构体数组中所有分量链接到备用链表中
void reserveArr(component *array);
//初始化静态链表
int initArr(component *array);
//输出函数
void displayArr(component * array, int body);
//从备用链表上摘下空闲节点的函数
int mallocArr(component * array);
int main() {
    component array[MAX_SIZE];
    int body = initArr(array);
    printf("静态链表为：\n");
    displayArr(array, body);
    return 0;
}
//创建备用链表
void reserveArr(component *array) {
    int i = 0;
    for (i = 0; i < MAX_SIZE; i++) {
        array[i].cur = i + 1;//将每个数组分量链接到一起
        array[i].data = 0;
    }
    array[MAX_SIZE - 1].cur = 0;//链表最后一个结点的游标值为0
}
//提取分配空间
int mallocArr(component * array) {
    //若备用链表非空，则返回分配的结点下标，否则返回 0（当分配最后一个结点时，该结点的游标值为 0）
    int i = array[0].cur;
    if (array[0].cur) {
        array[0].cur = array[i].cur;
    }
    return i;
}
//初始化静态链表
int initArr(component *array) {
    int tempBody = 0, body = 0;
    int i = 0;
    reserveArr(array);
    body = mallocArr(array);
    //建立首元结点
    array[body].data = 1;
    array[body].cur = 0;
    //声明一个变量，把它当指针使，指向链表的最后的一个结点，当前和首元结点重合
    tempBody = body;
    for (i = 2; i < 4; i++) {
        int j = mallocArr(array); //从备用链表中拿出空闲的分量
        array[j].data = i;      //初始化新得到的空间结点
        array[tempBody].cur = j; //将新得到的结点链接到数据链表的尾部
        tempBody = j;             //将指向链表最后一个结点的指针后移
    }
    array[tempBody].cur = 0;//新的链表最后一个结点的指针设置为0
    return body;
}
void displayArr(component * array, int body) {
    int tempBody = body;//tempBody准备做遍历使用
    while (array[tempBody].cur) {
        printf("%d,%d\n", array[tempBody].data, array[tempBody].cur);
        tempBody = array[tempBody].cur;
    }
    printf("%d,%d\n", array[tempBody].data, array[tempBody].cur);
}