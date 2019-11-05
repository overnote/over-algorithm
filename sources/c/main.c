#include "./test/test_SqList.c"
#include "./test/test_LinkList.c"
#include "./test/test_DoubleList.c"


int main() {

    // test_SqList();           // 测试顺序表
    // test_LinkList();          // 测试单链表
    test_DoubleList();          // 测试双向循环链表

    return 0;
}
