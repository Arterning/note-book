#include <stdio.h>
#include <string.h>

//从源source中拷贝n个字节到目标destin中 
int main()
{
    char src[] = "*********";
    char dest[] = "klmnopqrst";
    char *ptr;

    printf("Destination before memcpy is: %s\n",dest);
    ptr = memcpy(dest,src,strlen(src));

    if(!ptr){
	    printf("memcpy failed \n");
    }

    return 0 ;

}

