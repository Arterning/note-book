# 关于手动设置主键

<p>
如果你要手动设置主键
就不要设置GenerationType

我也是服了，网上的结果和chatGPT的答案都是错误的
GenerationType.identity是主键自增
GenerationType.auto是hibernate生成主键
</p>

# 主键字段映射
<p>
如果主键类型是bigInt
那么对应的entity id类型最好定义为Long
虽然可以定义为String 但是最好不要这么做
我也是服了，由于之前的代码生成器已经生成好了
类型都是String
今天吭哧吭哧的改了大半天 真他妈费劲
</p>
