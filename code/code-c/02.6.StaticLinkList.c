/**
 * 静态链表
 */

#include <stdio.h>
#include <stdlib.h>

#define MAX_SIZE 10 // 静态链表最大长度
typedef int ElemType;

// 该定义使用方式： SLinkList arr;
typedef struct {
    ElemType data;
    int next; // 下一个元素的数组下标
} SLinkList[MAX_SIZE];

// 上述定义等价于：(初始化可使用 struct Node arr[MAX_SIZE];)
// typedef struct SNode {
//     ElemType data;
//     int next; // 下一个元素的数组下标
// };
// typedef struct SNode SLinkList[MAX_SIZE];

int InitSLinkList() {
    // arr[0].next = -1
}