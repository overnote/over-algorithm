## 一 需要的文件

test.h：
```c
#ifndef TEST
#define TEST
int test();
#endif
```

test.c：
```
#include <stdio.h>

int test() {
    printf("test@@@@....\n");
    return 0;
}
```

同样test.h和test.c内容一致，main.c如下：
```c
#include <stdio.h>
#include "test.h"
#include "test2.h"

int main() {

    printf("main...\n");
    test();
    test2();
    return 0;
}
```

编译：
```
gcc -o main main.c test.c
```

## 二 makefile

#### 2.1 makefile概述

多文件开发会带来复杂的编译问题，Makefile可以让编译进入自动化，免去手工编译大量文件的困难。makefile文件中定义了一系列的规则来指定哪些文件需要先编译/后编译/重新编译等，其操作类似shell脚本，使用make命令即可实现一个大型项目的自动编译。  

通常在一个项目里，其规则是：
- 如果这个工程没有编译过，那么所有C文件都要编译并被链接。
- 如果这个工程的某几个C文件被修改，那么只编译被修改的C文件，并链接目标程序。
- 如果这个工程的头文件被改变了，那么需要编译引用了这几个头文件的C文件，并链接目标程序。

make命令会自动智能地根据当前的文件修改的情况来确定哪些文件需要重编译，从而自己编译所需要的文件和链接目标程序。  

#### 2.1 makefile基础示例

在1.3示例创建文件makefile：
```
main: main.c test.c
	gcc -o main main.c test.c
```

在该项目中只有一个main.c文件，makefile的意思：
- 第一行：main代表要生成的文件名，以及生成该文件依赖的文件
- 第二行：tab空格后，代表要生成第一行文件时使用的命令

运行：直接执行 make 命令就可以生成一个名为main的二进制文件，再次执行main则会提示已经是最新

#### 2.2 makefile语法规范

makefile由一组规则组成,规则如下:
```
目标: 依赖
（tab）命令
```

makefile基本规则三要素：
- 目标: 要生成的目标文件
- 依赖: 目标文件由哪些文件生成
- 命令: 通过执行该命令由依赖文件生成目标

#### 2.3 makefile原理

生成目标的过程：
- 检查规则中的所有的依赖文件是否都存在（会持续依次向下搜索查看），不存在则报错
- 如果所有依赖都存在，则检查是否需要更新，任何一个依赖被更新，则目标必须更新


#### 2.4 示例增强1-避免书写冗余

第一版示例：该示例中，每次修改都要全部编译一遍
```
main: main.c test.c test2.c
	gcc -o main main.c test.c test2.c -I./
```

第二版：这时，改动一个文件，只会编译被改动的文件，而不会全部编译。缺点是：.c文件太多时，书写非常冗余
```
main: main.o test.o test2.o
	gcc -o main main.o test.o test2.o

main.o: main.c
	gcc -o main.o -c main.c -I./

test.o: test.c
	gcc -o test.o -c test.c

test2.o: test2.c
	gcc -o test2.o -c test2.c
```

#### 2.5 示例增强2-引入变量

makefile有三种变量：普通变量、自带变量、自动变量。  
```
# 普通变量
name = zs
myname = $(name)

# 自带变量：makefike提供的一些可以直接使用的变量
CC = gcc #arm-linux-gcc
CPPFLAGS : C预处理的选项 -I
CFLAGS:   C编译器的选项 -Wall -g -c
LDFLAGS :  链接器选项 -L  -l

# 自动变量：自动变量只能在规则的命令中使用
$@: 表示规则中的目标
$<: 表示规则中的第一个条件
$^: 表示规则中的所有条件, 组成一个列表, 以空格隔开, 如果这个列表中有重复的项则消除重复项。
```

第三版的makefile：
```
target=main
object=main.o test.o test2.o
CC=gcc
CPPFLAGS=-I./

$(target):$(object)
	$(CC) -o $@ $^
%.o:%.c 
	$(CC) -o $@ -c $< $(CPPFLAGS)
```

#### 2.6 示例增强3-引入函数  

makefile常用函数：
```
wildcard：查找指定目录下的指定类型的文件
src=$(wildcard *.c)  //找到当前目录下所有后缀为.c的文件,赋值给src

patsubst：匹配替换
obj=$(patsubst %.c,%.o, $(src)) //把src变量里所有后缀为.c的文件替换成.o
```

增强后的makefile：
```
src=$(wildcard ./*.c)
target=main
object=$(patsubst %.c, %.o, $(src))
CC=gcc
CPPFLAGS=-I./

$(target):$(object)
	$(CC) -o $@ $^
%.o:%.c 
	$(CC) -o $@ -c $< $(CPPFLAGS)
```


#### 2.7 示例增强3-清理编译后多余文件

`make clean`：如果当前目录下有同名clean文件,则不执行clean对应的命令, 解决方案：
```
伪目标声明:  .PHONY:clean
声明目标为伪目标之后, makefile将不会检查该目标是否存在或者该目标是否需要更新
```

clean命令中的特殊符号：
```
“-”此条命令出错,make也会继续执行后续的命令。如:“-rm main.o”
rm -f: 强制执行, 比如若要删除的文件不存在使用-f不会报错
“@”不显示命令本身, 只显示结果。如:“@echo clean done”
```

贴士：
```
make 默认执行第一个出现的目标, 可通过make dest指定要执行的目标
– make -f : -f执行一个makefile文件名称, 使用make执行指定的makefile: make -f mainmak
```

最终的makefile：
```
src=$(wildcard ./*.c)
target=main
object=$(patsubst %.c, %.o, $(src))
CC=gcc
CPPFLAGS=-I./

$(target):$(object)
	$(CC) -o $@ $^
%.o:%.c 
	$(CC) -o $@ -c $< $(CPPFLAGS)
.PHONY:clean
clean:
	-rm -f $(target) $(object)
```

执行：
```
make:进行编译
make clean：删除不必要的编译文件
```