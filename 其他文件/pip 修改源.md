### 清华源

pip install xxxx -i https://pypi.tuna.tsinghua.edu.cn/simple

### 阿里源

pip install xxxx -i https://mirrors.aliyun.com/pypi/simple/

### 腾讯源

pip install xxxx -i http://mirrors.cloud.tencent.com/pypi/simple

### 豆瓣源

pip install xxxx -i http://pypi.douban.com/simple/

**将xxxx换成需要安装的包的名字**

## 2.永久换源：

### 清华源

pip config set global.index-url https://pypi.tuna.tsinghua.edu.cn/simple

### 阿里源

pip config set global.index-url https://mirrors.aliyun.com/pypi/simple/

### 腾讯源

pip config set global.index-url http://mirrors.cloud.tencent.com/pypi/simple

### 豆瓣源

pip config set global.index-url http://pypi.douban.com/simple/

## 3.换回默认源

pip config unset global.index-url