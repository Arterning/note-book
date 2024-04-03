
我最近在看很多相关的资料，我试下能不能用比较简单的语言大家说清楚

## 第一， [ChatGPT](https://www.zhihu.com/search?q=ChatGPT&search_source=Entity&hybrid_search_source=Entity&hybrid_search_extra=%7B%22sourceType%22%3A%22answer%22%2C%22sourceId%22%3A3019957650%7D)核心就是个预测下一个字出现概率的函数

首先，我们得知道ChatGPT其实就是个语言模型，有点不一样的是，它是个大语言模型，也就是所谓的[LLM](https://www.zhihu.com/search?q=LLM&search_source=Entity&hybrid_search_source=Entity&hybrid_search_extra=%7B%22sourceType%22%3A%22answer%22%2C%22sourceId%22%3A3019957650%7D)。大在哪里呢？训练数据巨大，参数巨大。

参数是什么？要说清楚这个，首先这里就要说下模型这个概念。都说大语言模型是个[神经网络模型](https://www.zhihu.com/search?q=%E7%A5%9E%E7%BB%8F%E7%BD%91%E7%BB%9C%E6%A8%A1%E5%9E%8B&search_source=Entity&hybrid_search_source=Entity&hybrid_search_extra=%7B%22sourceType%22%3A%22answer%22%2C%22sourceId%22%3A3019957650%7D)，那这里的模型是什么意思呢？

其实模型就是个函数

**语言模型就是个接受你的输入，然后输出下一个字出现的概率的那么一个函数**

  
“谁” = ChatGPT(“你是”）


大语言模型就是你这个语言模型函数内部的参数非常巨大，百亿千亿级别的参数

一般我们见到的函数就是y = ax + b

试想下人家的函数是 z = a1x1 + a1x2 + ... a千亿x千亿 + b1 + ... b千亿，然后为了得到y，还需要考虑有好多个z，然后每个z要引入非线性变换[激活函数](https://www.zhihu.com/search?q=%E6%BF%80%E6%B4%BB%E5%87%BD%E6%95%B0&search_source=Entity&hybrid_search_source=Entity&hybrid_search_extra=%7B%22sourceType%22%3A%22answer%22%2C%22sourceId%22%3A3019957650%7D)（这句不好理解的话可以忽略）

这里的a1,b1，a千亿，b千亿的就是上面说的参数，巨大的参数！

知道了大语言模型是怎么回事之后，我们就稍微说下背后的原理。


## Transformer 模型

我们都知道现在的大语言模型基本上都是基于[transformer](https://www.zhihu.com/search?q=transformer&search_source=Entity&hybrid_search_source=Entity&hybrid_search_extra=%7B%22sourceType%22%3A%22answer%22%2C%22sourceId%22%3A3019957650%7D)实现的。

transformer分成两个部分，一个是encoder编码器做编码用，一个[decoder解码器](https://www.zhihu.com/search?q=decoder%E8%A7%A3%E7%A0%81%E5%99%A8&search_source=Entity&hybrid_search_source=Entity&hybrid_search_extra=%7B%22sourceType%22%3A%22answer%22%2C%22sourceId%22%3A3019957650%7D)做解码用。

encoder的输入是我们的文本，比如我们的问题，输出的是通过某些技术编码后的向量，你可以简单理解成加密。

至于是怎么加密的，下面会稍微说下，如果这一部分不感兴趣的可以忽略，不影响整体高层的理解。

首先输入文本会转化成word embeddings，所谓的[word embeddinds](https://www.zhihu.com/search?q=word%20embeddinds&search_source=Entity&hybrid_search_source=Entity&hybrid_search_extra=%7B%22sourceType%22%3A%22answer%22%2C%22sourceId%22%3A3019957650%7D)这里就是一个单词的向量化表示，该word embeddings会经过[多头注意力机制](https://www.zhihu.com/search?q=%E5%A4%9A%E5%A4%B4%E6%B3%A8%E6%84%8F%E5%8A%9B%E6%9C%BA%E5%88%B6&search_source=Entity&hybrid_search_source=Entity&hybrid_search_extra=%7B%22sourceType%22%3A%22answer%22%2C%22sourceId%22%3A3019957650%7D)来根据每个单词在句子中的重要程度来重新调整每个单词的word embeddings的表示，也就是说对句子更重要的单词就会获得更多的权重。

比如”我趁着月色去偷渡“，这里的”我“和”偷渡“对整个句子更重要，所以就会获得更多的权重。最后，经过注意力加权之后的word embeddings就会变成我们的注意力向量，这就是encoder的输出。




  
decoder的输入有两个部分，一个就是我们encoder的输出，也就是我们输入的问题转化成的注意力向量。另外一个输入就是动态生成的输出，比如我输入”我趁着月色去偷渡“，然后ChatGPT给出了第一个新生成的输出”日“， 那它就会和前面已经生成的静态的[自注意力向量](https://www.zhihu.com/search?q=%E8%87%AA%E6%B3%A8%E6%84%8F%E5%8A%9B%E5%90%91%E9%87%8F&search_source=Entity&hybrid_search_source=Entity&hybrid_search_extra=%7B%22sourceType%22%3A%22answer%22%2C%22sourceId%22%3A3019957650%7D)一起工作，去预测下一个单词，可能是“本”。然后又会将”日本“和自注意力向量一起输入decoder来预测下一个输出，如此往复。

decoder内部也用到了[自注意力机制](https://www.zhihu.com/search?q=%E8%87%AA%E6%B3%A8%E6%84%8F%E5%8A%9B%E6%9C%BA%E5%88%B6&search_source=Entity&hybrid_search_source=Entity&hybrid_search_extra=%7B%22sourceType%22%3A%22answer%22%2C%22sourceId%22%3A3019957650%7D)，主要是我们的输出，比如上面的"日本xxxx"，这些输出每个单词在已生成的句子中的重要程度也不一样，也应该让transformer生成下一个单词时关注句子不同的部分，以便得到更好的生成效果。

以上这些如果有听不懂的，没有关系，你就将decoder理解成解码器，另外一个黑箱，你将编码后的输入丢进去，出来的就是新生成回答就好了。

  
  
## 第三，预训练，instruction、[prompt](https://www.zhihu.com/search?q=prompt&search_source=Entity&hybrid_search_source=Entity&hybrid_search_extra=%7B%22sourceType%22%3A%22answer%22%2C%22sourceId%22%3A3019957650%7D) fine-tuning以及[RLHF](https://www.zhihu.com/search?q=RLHF&search_source=Entity&hybrid_search_source=Entity&hybrid_search_extra=%7B%22sourceType%22%3A%22answer%22%2C%22sourceId%22%3A3019957650%7D)

确定了transform作为语言模型核心之后，跟着要做的就是喂入大量的语料来进行预训练。训练时会用到一个叫做Masked Language Modeling的技术，aka MLM。

MLM其实说白了就是将从互联网比如wikipedia从拿下来的一个句子，掩盖住一些部分，然后让ChatGPT不停的去猜缺失部分是什么，然后根据结果的好坏去调整整个语言模型函数的参数([梯度下降算法](https://www.zhihu.com/search?q=%E6%A2%AF%E5%BA%A6%E4%B8%8B%E9%99%8D%E7%AE%97%E6%B3%95&search_source=Entity&hybrid_search_source=Entity&hybrid_search_extra=%7B%22sourceType%22%3A%22answer%22%2C%22sourceId%22%3A3019957650%7D))，直到结果逼近答案为止。

比如上面的”我趁着月色去偷渡日本“，MLM随机掩盖，比如”我趁着月色去偷渡__“，然后让ChatGPT去猜。其实这个就类似于监督学习，只是缺失部分的标准答案”日本“并不是我们手动去标的，而是通过MLM去自动做的。

完成了预训练后，ChatGPT还只能做输出预测，即你输入一部分，然后ChatGPT往下预测并生成下一部分是什么。

所以我输入

  
> 怎么去日本？

此时的模型很有可能回答的是

> 怎么去美国  
> 怎么去岛国  
> 怎么去越南

很明显，这不是我们想要的。

如果要ChatGPT能够更好的和人类进行交互，比如根据用户的instruction指令或者prompt提问提示词给出答案，而不是直接在后面追加生成的内容，那这个时候就还需要进行微调，

而这就是instruction和prompt fine-tuning要做的事情了。

完了后我们的模型就基本具备ChatGPT的功能了，但是这时的模型也只是知道作问答，答案的好坏它是不保证的

我们知道语言模型就是猜概率问题，最终会根据概率取样，看下一个单词应该生成什么。

因为这种随机性，所以生成的答案多种多样，很有可能不是我们人类想要的

  
怎么训练呢？

首先找一批人和ChatGPT提问，这里一个prompt提多次，就会得到多个答案，这时人类会给这些答案的好坏进行打分

有了样本，我们就可以根据这些样本去训练出一个奖励模型RM。其最终的输入就是一个prompt和一个答案，输出就是这个回答的奖励，即一个数字评分。

有了这个奖励模型，我们就能用机器自动判定ChatGPT每个输出的好坏，然后根据结果反过来通过梯度下降算法来调整ChatGPT的参数来让输出的答案的奖励得分越高

其实整一个过程就是我们常用的误差反向传导方法，当然，底层实现肯定更复杂，用的是强化学习，但意思就是这么个意思

最终这一套东西就被[openai](https://www.zhihu.com/search?q=openai&search_source=Entity&hybrid_search_source=Entity&hybrid_search_extra=%7B%22sourceType%22%3A%22answer%22%2C%22sourceId%22%3A3019957650%7D)命名为RLHF，即带有人类反馈的强化学习，人类反馈就是我们上面人类给打分得到个奖励模型，强化学习就是通过引入奖励模型来实现的这个能自我调优的强化我们生成结果的学习方法

完了后就得到我们现在大家都在用的ChatGPT了！


  
  
作者：天地会珠海分舵  
链接：https://www.zhihu.com/question/598243591/answer/3019957650  
来源：知乎  
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。