#include <stdio.h>
#include <string.h>

int main()
{
    char *s = "hello3ning";
    char *p;
    p = strchr(s,'3');//找出 s 字符串中第一次出现的字符 3 的地址，然后将该地址返回。
    printf("%s\n",p);
    return 0 ;
}
