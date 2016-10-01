#ifndef MD5_H
#define MD5_H

typedef struct
{
	unsigned int a[5];
	unsigned int b[5];
	unsigned char c[5];
} testStruct;

//函数在C语言中需要先声明后使用 否则会出现Implicit Declaration of Function 警告
int printBoyNextDoor();
#endif
