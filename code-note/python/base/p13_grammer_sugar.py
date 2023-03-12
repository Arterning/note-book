# Python 体贴程序员的语法糖

# 交换变量的值
a = 12
b = 45
a, b = b, a
print (a,b)


# 判断范围
score = 95
if 90 <= score <= 100:
    print("good score!!")


# 快速构造字符串
print("-"*60)

# 拼接数组
array_a = [3,3,4,5,5,5]
array_b = [9,9,9,9]
print(array_a + array_b)


# array slice
array_c = [1,2,3,4,5,6,7,8,9]
print(array_c[3:-2])
print(array_c[3:])
print(array_c[:-2])



# 解构赋值
a = (1,2,3)
# x, y, z, d = a
# print(x,y,z, d)
x, y, z = a
print(x,y,z)


