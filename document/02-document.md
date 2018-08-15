# 一些字符串的操作方法

## memset

void *memset(void *str, int c, size_t n)
copy c char to  str 的前 n 个字符。

```c
#include <stdio.h>
#include <string.h>
int main ()
{
   char str[50];
 
   strcpy(str,"This is string.h library function");
   puts(str);//print This is string.h library function
 
   memset(str,'$',7);//print $$$$$$$ string.h library function
   puts(str);
   
   return(0);
}
```

## strstr

char *strstr(const char *haystack, const char *needle)
该函数返回在 haystack 中第一次出现 needle 字符串的位置，如果未找到则返回 null。


# strncpy

char *strncpy(char *dest, const char *src, size_t n)
dest -- 指向用于存储复制内容的目标数组。
src -- 要复制的字符串。
n -- 要从源中复制的字符数。


# strchr

char *strchr(const char *str, int c)
str -- 要被检索的 C 字符串。
c -- 在 str 中要搜索的字符。

# bzero

bzero()函数,是一个C语言函数，但不是标准库函数，没有在ANSI中定义。 目前Linux的GCC支持。 函数功能：将指定内存块的前n个字节全部设置为零。



# 关于逻辑表达式的返回数据类型
在 C 语言中，没有boolean 类型，逻辑表达式返回的数据类型是 int 类型。当逻辑表达式为真时，返回整数值1；当逻辑表达式为假时，返回整数值0。这个整数值可以被用于做条件判断，例如在 if 语句、while 循环等中使用。在 C 语言中，0 代表假，非0代表真。

