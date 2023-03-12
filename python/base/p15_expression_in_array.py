# 列表推导式
a = [1,2,3,4]

b = [e+233 for e in a]
print(b)


# old method
b = []
for e in a:
    b.append(e +233)

print(b)