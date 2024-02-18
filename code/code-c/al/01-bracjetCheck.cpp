/**
 * 括号匹配
 */
#include <stdio.h>
#include <stdlib.h>

#define MAX_SIZE 10

typedef struct {
    char data[MAX_SIZE];
    int top;
} SqStack;

bool bracketCheck(char str[], int length) {
    SqStack S;
    InitStack(S);
    for (int i = 0; i < length; i++) {
        if(str[i] == '(' || str[i] == '[' || str[i] == '{'){
            Push(S, str[i]);
        } else {
            if(IsEmpty(S)){
                return false;       // 括号匹配失败
            }

            char topElem;
            Pop(S, topElem);
            if(str[i] == ')' && topElem != '('){
                return false;
            }
            if(str[i] == ']' && topElem != '['){
                return false;
            }
            if(str[i] == '}' && topElem != '{'){
                return false;
            }
        }
    }
    return IsEmpty(S);
}
