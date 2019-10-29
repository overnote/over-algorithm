## 一 gdb初识

gdb是GNU 发布的一个强大的程序调试工具，使用方式：

```
# -g 编译
gcc -g main.c -o main 

# 启动 gdb 调试
$ gdb main

# 进入gdb命令行，输入help查看帮助
help
help break          # 查看 break 相关子命令

# 启动后设置运行参数
set args 可指定运行时参数。（如：set args 10 20 30 40 50 ）
show args 命令可以查看设置好的运行参数

# 启动程序
run：程序开始执行, 如果有断点, 停在第一个断点处
start：程序向下执行一行。(在第一条语句处停止)

# 退出
q
```

## 二 利用gdb显示源码

GDB 可以打印出所调试程序的源代码,当程序停下来以后, GDB会报告程序停在了那个文件的第几行上。你可以用list命令来打印程序的源代码, 默认打印10行, list命令的用法如下所示:

- list linenum：打印第linenum行的上下文内容.
- list function：显示函数名为function的函数的源程序。
- list： 显示当前行后面的源程序。
- list -：显示当前文件开始处的源程序。
- list file:linenum: 显示file文件下第n行
- list file:function: 显示file文件的函数名为function的函数的源程序,一般是打印当前行的上5行和下5行, 如果显示函数是是上2行下8行, 默认是10行
- set listsize count：设置一次显示源代码的行数。         
- show listsize：   查看当前listsize的设置。

## 三 设置断点

断点：break 设置断点, 可以简写为b
```
# 当前文件
b 10 设置断点, 在源程序第10行
b func 设置断点, 在func函数入口处

# 多文件设置断点---其他文件
b filename:linenum --在源文件filename的linenum行处停住
b filename:function --在源文件filename的function函数的入口处停住

# 查询断点
info b == info break == i break == i b
```

条件断点：一般来说, 为断点设置一个条件, 我们使用if关键词, 后面跟其断点条件
```
b test.c:8 if intValue == 5
```

删除断点：delete [range...] 删除指定的断点, 其简写命令为d
```
如果不指定断点号, 则表示删除所有的断点。range表示断点号的范围（如：3-7）。
    删除某个断点: delete num
    删除多个断点: delete num1 num2  ...
    删除连续的多个断点: delete m-n
    删除所有断点: delete
```


使断点无效：disable [range...] 使指定断点无效, 简写命令是dis
```
如果什么都不指定, 表示disable所有的停止点。
    使一个断点无效/有效: disable num
    使多个断点无效有效: disable num1 num2 ...
    使多个连续的断点无效有效: disable m-n
    使所有断点无效有效: disable
```

使断点生效：enable [range...] 使无效断点生效, 简写命令是ena
```
如果什么都不指定, 表示enable所有的停止点。
    使一个断点无效/有效: enable num
    使多个断点无效有效: enable num1 num2 ...
    使多个连续的断点无效有效: enable m-n
    使所有断点无效有效: disable/enable
```

## 四 调试代码

- run 运行程序, 可简写为r
- next 单步跟踪, 函数调用当作一条简单语句执行, 可简写为n
- step 单步跟踪, 函数调进入被调用函数体内, 可简写为s
- finish 退出进入的函数, 如果出不去, 看一下函数体中的循环中是否有断点，如果有删掉，或者设置无效
- until 在一个循环体内单步跟踪时, 这个命令可以运行程序直到退出循环体,可简写为u,如果出不去, 看一下函数体中的循环中是否有断点，如果有删掉，或者设置无效
- continue 继续运行程序, 可简写为c(若有断点则跳到下一个断点处)

## 五 查看变量的值

查看运行时变量的值：print 打印变量、字符串、表达式等的值, 可简写为p
```
p count -----打印count的值
```

自动显示变量的值：你可以设置一些自动显示的变量, 当程序停住时, 或是在你单步跟踪时, 这些变量会自动显示。相关的GDB命令是display
```
display 变量名

info display -- 查看display设置的自动显示的信息。

undisplay num（info display时显示的编号）

delete display dnums… -- 删除自动显示, dnums意为所设置好了的自动显式的编号。如果要同时删除几个, 编号可以用空格分隔, 如果要删除一个范围内的编号, 可以用减号表示（如：2-5）
    删除某个自动显示: undisplay num 或者delete display num
    删除多个: delete display num1 num2
    删除一个范围: delete display m-n

disable display dnums…
    使一个自动显示无效: disable display num
    使多个自动显示无效: delete display num1 num2
    使一个范围的自动显示无效: delete display m-n

enable display dnums…
    使一个自动显示有效: enable display num
    使多个自动显示有效: enable display num1 num2
    使一个范围的自动显示有效: enable display m-n

disable和enalbe不删除自动显示的设置, 而只是让其失效和恢复。
```

查看修改变量的值：
```
ptype width --查看变量width的类型
type = double

p width --打印变量width 的值
$4 = 13
```

你可以使用`set var`命令来告诉GDB, width不是你GDB的参数, 而是程序的变量名, 如：
```
set var width=47  // 将变量var值设置为47
```
在你改变程序变量取值时, 最好都使用`set var`格式的GDB命令。