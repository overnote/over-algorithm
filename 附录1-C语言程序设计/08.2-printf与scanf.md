## 一 printf 函数

| 打印格式 | 对应数据类型 | 含义 |
| ------ | ------ | ------ |
| %d | int | 接受整数值并将它表示为有符号的十进制整数 |
| %hd | short | 短整数 |
| %hu | unsigned short | 无符号短整数 |
| %o | unsigned int | 无符号8进制整数 |
| %u | unsigned int | 无符号10进制整数 |
| %x,%X | unsigned int | 无符号16进制整数，x对应的是abcdef，X对应的是ABCDEF |
| %f | float | 单精度浮点数 |
| %lf | double | 双精度浮点数 |
| %e,%E | double | 科学计数法表示的数，此处"e"的大小写代表在输出时用的"e"的大小写 |
| %c | char | 字符型。可以把输入的数字按照ASCII码相应转换为对应的字符 |
| %s | char * | 字符串。输出字符串中的字符直至字 类5通过  符串中的空字符（字符串以'\0‘结尾，这个'\0'即空字符） |
| %p | void * | 以16进制形式输出指针 |
| %% | % | 输出一个百分号 |

printf附加格式：
```
l：附加在d,u,x,o前面，表示长整数
-：左对齐
m：代表一个整数，数据最小宽度
0：数字0，将输出的前面补上0直到占满指定列宽为止不可以搭配使用-
m.n：代表一个整数，m指域宽，即对应的输出项在输出设备上所占的字符数。n指精度，用于说明输出的实型数的小数位数。对数值型的来说，未指定n时，隐含的精度为n=6位。
```

## 二 scanf

scanf函数与getchar函数：
- getchar是从标准输入设备读取一个char。
- scanf通过%转义的方式可以得到用户通过标准输入设备输入的数据。

```c
#include <stdio.h>

int main(){
	char ch1;
	char ch2;
	char ch3;
	int a;
	int b;

	printf("请输入ch1的字符：");
	ch1 = getchar();
	printf("ch1 = %c\n", ch1);

	getchar(); //测试此处getchar()的作用

	printf("请输入ch2的字符：");
	ch2 = getchar();
	printf("\'ch2 = %ctest\'\n", ch2);

	getchar(); //测试此处getchar()的作用
	printf("请输入ch3的字符：");
	scanf("%c", &ch3);//这里第二个参数一定是变量的地址，而不是变量名
	printf("ch3 = %c\n", ch3);

	printf("请输入a的值：");
	scanf("%d", &a);
	printf("a = %d\n", a);

	printf("请输入b的值：");
	scanf("%d", &b);
	printf("b = %d\n", b);

	return 0;
}
```