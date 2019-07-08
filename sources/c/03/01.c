#include <stdio.h>

// 卡拉兹（Callatz）猜想
int callatz(int num) {

    if (num <= 0) {
        printf("参数不合法");
        return 0;
    }

    int step = 0;

    while (num != 1) {
        if (num % 2 == 0) {
            num = num / 2;
        } else {
            num = (3 * num + 1) / 2;
        }
        step ++;
    }
    
    return step;
}

// 奥巴马绘图
int obama(int col, char c) {

    if (col <= 0) {
        printf("参数不合法\n");
        return 0;
    }

    int row = 0;

    if (col % 2 == 1) {
        row = col / 2 + 1;
    } else {
        row = col / 2;
    }

    // 绘制第一行
    for (int i = 0; i < col; i++) {
        printf("%c", c);
    }
    printf("\n");

    // 绘制中间部分
    for (int i = 2; i < row; i++) {
        printf("%c", c);
        for (int j = 0; j < col - 2; j++){
            printf(" ");
        }
        printf("%c\n", c);
    }

    // 绘制最后一行
    for (int i = 0; i < col; i++){
        printf("%c", c);
    }
    printf("\n");
    return 0;

}

// 判断回文

int judge(char str[]) {
    int len = strlen(str);
    for (int i = 0; i < len / 2; i++) {
        if (str[i] != str[len - 1 - i]){
            return 0;
        }
    }
    return 1;
}

int main() {

    return 0;
}