## 一 基础示例

### 时间复杂度求解步骤

- 步骤一：找出基本操作，一般是最深层循环的循环体
- 步骤二：求出执行次数的最大数量级

### 示例一
```c
void fn(int n) {
    int i = 1;
    while(i < n){
        i = i * 2;
    }
}
```

- 第一步：本题基础操作是： `i = i * 2;`
- 第二步：循环次数的最大数量级
  - 每次循环后 `i = i * 2;`
  - 假设循环 t 次后，会出现 i 最终趋近于 n，跳出循环。i = 2<sup>t</sup>
  - 即：2<sup>t</sup> < n，也即 t<log<sup>2</sup>n
  - 时间复杂度函数f(n)=log<sup>2</sup>n，即T(n)=O(log<sup>2</sup>n) 

### 示例二
```c
void fn(int n) {
    int i = 0, sum = 0;
    while(sum < n) {
        ++i;
        sum = sum + i;
    }
}
```

- 第一步：本题基础操作是： `++i;sum = sum + i;`
- 第二步：循环次数的最大数量级
  - 每次循环后 `sum = sum + i;`
  - 假设循环 t 次后，会出现 i 最终趋近于 n，跳出循环，则：

```
sum1 = 1
sum2 = 1 + 2
sum3 = 1 + 2 + 3
...
sumt = 1 + 2 + 3 + ... + t = t(t + 1)/2
```

为了更精确的计算，不再像示例一种直接预估，而是引入一个常数K弥补差距：m(m+1)/2 + K = n。  

得到时间复杂度函数f(n) = m=  (-1+$\sqrt{8n+1-8K}$)/2   

那么时间复杂度为O($\sqrt{n}$)  

### 示例三

```c
void mergesort(int i, int j) {    // 本次示例传入参数为 1, n
    int m;
    if (i != j) {
        mergesort(i, (i + j) / 2);
        mergesort((i + j) / 2 + 1, j);
        merge(i, j);             // 本函数时间复杂度为O(n)
    }
}
```

本题规模为n，基本操作是函数 merge()，由于该函数时间复杂度为O(n)，那么其内部的基本操作次数设定为 cn。mergesort函数的时间复杂度f(n)为：  
- f(n)=2f(n/2)+cn=`2*(2*f(n/4)+cn*1/2)+cn`
- f(n)=2<sup>2</sup>f(n/4)+2cn
- f(n)=2<sup>3</sup>f(n/8)+3cn
- ...
- f(n)=2<sup>k</sup>f(n/2<sup>k</sup>)+kcn
- f(1)=c*1=c

由最后2步得知n=2<sup>k</sup>时，即k=log<sub>2</sub>k时，  
- f(n)=c*n+cnlog<sub>2</sub>n
- 最终时间复杂度为T(n)=O(nlog<sub>2</sub>n)