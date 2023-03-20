

#include <stdio.h>
#include <string.h>

//char *strcat(char *dest, const char *src);
//strcat() 会将参数 src 字符串复制到参数 dest 所指的字符串尾部
int main()
{
    char str[80];
    strcpy(str,"these ");
    strcat(str, "add aaaaa");
    strcat(str,"add bbbbb");
    strcat(str,"add ccccccc");
    puts(str);
    return 0 ;
}
