## 一 引用基本用法

引用是c++对c的重要扩充。在c/c++中指针的作用基本都是一样的，但是c++增加了另外一种给函数传递地址的途径，这就是按引用传递(pass-by-reference)，它也存在于其他一些编程语言中，并不是c++的发明。

- 变量名实质上是一段连续内存空间的别名，是一个标号(门牌号)
- 程序中通过变量来申请并命名内存空间
- 通过变量的名字可以使用存储空间

```
对一段连续的内存空间只能取一个别名吗？
c++中新增了引用的概念，引用可以作为一个已定义变量的别名。
```

引用语法：
```c+=
Type& ref = val;
```

贴士：
- &在此不是求地址运算，而是起标识作用。
- 类型标识符是指目标变量的类型
- 必须在声明引用变量时进行初始化。
- 引用初始化之后不能改变。
- 不能有NULL引用。必须确保引用是和一块合法的存储单元关联。
- 可以建立对数组的引用。

```c++
//1. 认识引用
void test01(){

	int a = 10;
	//给变量a取一个别名b
	int& b = a;
	cout << "a:" << a << endl;
	cout << "b:" << b << endl;
	cout << "------------" << endl;
	//操作b就相当于操作a本身
	b = 100;
	cout << "a:" << a << endl;
	cout << "b:" << b << endl;
	cout << "------------" << endl;
	//一个变量可以有n个别名
	int& c = a;
	c = 200;
	cout << "a:" << a << endl;
	cout << "b:" << b << endl;
	cout << "c:" << c << endl;
	cout << "------------" << endl;
	//a,b,c的地址都是相同的
	cout << "a:" << &a << endl;
	cout << "b:" << &b << endl;
	cout << "c:" << &c << endl;
}
```

使用引用注意事项
```c++
void test02(){
	//1) 引用必须初始化
	//int& ref; //报错:必须初始化引用
	//2) 引用一旦初始化，不能改变引用
	int a = 10;
	int b = 20;
	int& ref = a;
	ref = b; //不能改变引用
	//3) 不能对数组建立引用
	int arr[10];
	//int& ref3[10] = arr;
}
```

建立数组引用:
```c++
//1. 建立数组引用方法一
	typedef int ArrRef[10];
	int arr[10];
	ArrRef& aRef = arr;
	for (int i = 0; i < 10;i ++){
		aRef[i] = i+1;
	}
	for (int i = 0; i < 10;i++){
		cout << arr[i] << " ";
	}
	cout << endl;
	//2. 建立数组引用方法二
	int(&f)[10] = arr;
	for (int i = 0; i < 10; i++){
		f[i] = i+10;
	}
	for (int i = 0; i < 10; i++){
		cout << arr[i] << " ";
	}
	cout << endl;
```

## 二 函数中的引用

最常见看见引用的地方是在函数参数和返回值中。当引用被用作函数参数的时，在函数内对任何引用的修改，将对还函数外的参数产生改变。当然，可以通过传递一个指针来做相同的事情，但引用具有更清晰的语法。  

如果从函数中返回一个引用，必须像从函数中返回一个指针一样对待。当函数返回值时，引用关联的内存一定要存在。

```c++
//值传递
void ValueSwap(int m,int n){
	int temp = m;
	m = n;
	n = temp;
}
//地址传递
void PointerSwap(int* m,int* n){
	int temp = *m;
	*m = *n;
	*n = temp;
}
//引用传递
void ReferenceSwap(int& m,int& n){
	int temp = m;
	m = n;
	n = temp;
}
void test(){
	int a = 10;
	int b = 20;
	//值传递
	ValueSwap(a, b);
	cout << "a:" << a << " b:" << b << endl;
	//地址传递
	PointerSwap(&a, &b);
	cout << "a:" << a << " b:" << b << endl;
	//引用传递
	ReferenceSwap(a, b);
	cout << "a:" << a << " b:" << b << endl;
}
```

通过引用参数产生的效果同按地址传递是一样的。引用的语法更清楚简单：	
- 1) 函数调用时传递的实参不必加“&”符 
- 2) 在被调函数中不必在参数前加“*”符
引用作为其它变量的别名而存在，因此在一些场合可以代替指针。C++主张用引用传递取代地址传递的方式，因为引用语法容易且不易出错。  

```c++
//返回局部变量引用
int& TestFun01(){
	int a = 10; //局部变量
	return a;
}
//返回静态变量引用
int& TestFunc02(){	
	static int a = 20;
	cout << "static int a : " << a << endl;
	return a;
}
int main(){
	//不能返回局部变量的引用
	int& ret01 = TestFun01();
	//如果函数做左值，那么必须返回引用
	TestFunc02();
	TestFunc02() = 100;
	TestFunc02();

	return EXIT_SUCCESS;
}
```

注意：
- 不能返回局部变量的引用。
- 函数当左值，必须返回引用。

## 三 引用的本质

引用的本质在c++内部实现是一个指针常量：
```c++
Type& ref = val; // Type* const ref = &val;
```
c++编译器在编译过程中使用常指针作为引用的内部实现，因此引用所占用的空间大小与指针相同，只是这个过程是编译器内部实现，用户不可见。
```c++
//发现是引用，转换为 int* const ref = &a;
void testFunc(int& ref){
	ref = 100; // ref是引用，转换为*ref = 100
}
int main(){
	int a = 10;
	int& aRef = a; //自动转换为 int* const aRef = &a;这也能说明引用为什么必须初始化
	aRef = 20; //内部发现aRef是引用，自动帮我们转换为: *aRef = 20;
	cout << "a:" << a << endl;
	cout << "aRef:" << aRef << endl;
	testFunc(a);
	return EXIT_SUCCESS;
}
```

## 四 指针引用

在c语言中如果想改变一个指针的指向而不是它所指向的内容，函数声明可能这样:`void fun(int**);`  

给指针变量取一个别名：

```c++
Type* pointer = NULL;  
Type*& = pointer;
```

```c++
struct Teacher{
	int mAge;
};
//指针间接修改teacher的年龄
void AllocateAndInitByPointer(Teacher** teacher){
	*teacher = (Teacher*)malloc(sizeof(Teacher));
	(*teacher)->mAge = 200;  
}
//引用修改teacher年龄
void AllocateAndInitByReference(Teacher*& teacher){
	teacher->mAge = 300;
}
void test(){
	//创建Teacher
	Teacher* teacher = NULL;
	//指针间接赋值
	AllocateAndInitByPointer(&teacher);
	cout << "AllocateAndInitByPointer:" << teacher->mAge << endl;
	//引用赋值,将teacher本身传到ChangeAgeByReference函数中
	AllocateAndInitByReference(teacher);
	cout << "AllocateAndInitByReference:" << teacher->mAge << endl;
	free(teacher);
}
```
## 五 常量引用

常量引用的定义格式:
```c++
const Type& ref = val;
```

常量引用注意：
- 字面量不能赋给引用，但是可以赋给const引用
- const修饰的引用，不能修改。
```c++
void test01(){
	int a = 100;
	const int& aRef = a; //此时aRef就是a
	//aRef = 200; 不能通过aRef的值
	a = 100; //OK
	cout << "a:" << a << endl;
	cout << "aRef:" << aRef << endl;
}
void test02(){
	//不能把一个字面量赋给引用
	//int& ref = 100;
	//但是可以把一个字面量赋给常引用
	const int& ref = 100; //int temp = 200; const int& ret = temp;
}
```

const引用使用场景
```
    常量引用主要用在函数的形参，尤其是类的拷贝/复制构造函数。
将函数的形参定义为常量引用的好处:
- 引用不产生新的变量，减少形参与实参传递时的开销。
- 由于引用可能导致实参随形参改变而改变，将其定义为常量引用可以消除这种副作用。
    如果希望实参随着形参的改变而改变，那么使用一般的引用，如果不希望实参随着形参改变，那么使用常引用。

//const int& param防止函数中意外修改数据
void ShowVal(const int& param){
	cout << "param:" << param << endl;
}
```