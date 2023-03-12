# 不需要close
with open('test.txt', 'r') as f:
    data = f.read()
    print(data)