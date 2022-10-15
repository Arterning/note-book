#!/usr/bin/python
# -*- coding: UTF-8 -*-
 
# 调用函数时，默认参数的值如果没有传入，则被认为是默认值。下例会打印默认的age，如果age没有被传入：
def printinfo( name, age = 35 ):
   "打印任何传入的字符串"
   print "Name: ", name
   print "Age ", age
   return
 
printinfo( age=50, name="miki" )
printinfo( name="miki" )


# 加了星号（*）的变量名会存放所有未命名的变量参数。不定长参数实例如下：
def printinfo_2( arg1, *vartuple ):
       "打印任何传入的参数"
   print "输出: "
   print arg1
   for var in vartuple:
      print var
   return
 
# 调用printinfo 函数
printinfo_2( 10 )
printinfo_2( 70, 60, 50 )


# 可写函数说明
# 函数内取值:  [10, 20, 30, [1, 2, 3, 4]]
# 函数外取值:  [10, 20, 30, [1, 2, 3, 4]]
def changeme( mylist ):
   mylist.append([1,2,3,4])
   print "函数内取值: ", mylist
   return
mylist = [10,20,30]
changeme( mylist )
print "函数外取值: ", mylist


def ChangeInt( a ):
    a = 10
b = 2
ChangeInt(b)
print b # 结果是 2

