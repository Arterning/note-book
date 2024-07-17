
> ORM能够让事情变得简单，也会让有些事情变得复杂。有人说，这不就是一个SQL语句的事嘛，干嘛在ORM里面就这么复杂。

上篇文章我们讲了什么是ORM（对象关系映射)，不了解的可以看看上一篇文章。

这篇我们来解释什么是N+1的问题，在所有的ORM中，这都会是一个问题，新手很容易踩到坑。进而导致系统变慢，然后拖垮整个系统。

还是拿代码来说事，上篇我们定义了一个User的模型，这次还继续沿用，然后增加一个Post（文章）的模型。User和Post是一对多的关系，也就是User是Post的外键。代码如下:


```python
from django.db import models
class User(models.Model):
    name = models.CharField(max_length=255)
class Post(models.Model):
    owner = models.ForeignKey(User)   # by the5fire
    title = models.CharField(max_length=255)
    content = models.TextField()
```


假设我们有这样的代码，现在系统里面有十个用户，每个用户写了一篇文章，也就是十篇文章。

接下来我们有一个需求，展示一个文章列表页，列表页上展示的信息包括：文章标题，文章作者名称。就这两个字段，也不需要分页。

我们要查询出这样的数据要怎么做呢。在ORM的世界中，我们直观的做法是这样:

```python
posts = Post.objects.all()  #  获取所有的文章数据，注意此时不会执行sql语句  by the5fire
result = []
for post in posts:   # 此时会执行select * from post的查询
    result.append({
        'title': post.title,
        'owner': post.owner.name,  # 此时会执行  select * from user where user_id = <post.user_id>
    })
```


写到这就明白了吧。每次循环都要查一下user表，也就是说，如果我第一次查询是10条记录，那么最终我需要执行的查询语句就是10 + 1 = 11条语句。如果我第一次查询出来的是N条记录，那么最终需要执行的sql语句就是N+1次。

这就是N+1的问题。

但是如果懂SQL的话，就知道，其实这就是一个简单的JOIN语句。一条语句就能查出所有的数据，搞什么N+1.

```sql
SELECT t1.title, t2.name from post as t1 INNER JOIN user as t2 ON t2.id = t1.owner_id;
```

想到这里，是不是觉得ORM有时候挺碍事的。

其实现在的ORM框架基本都提供了解决的方案，比如Django中，对这类问题就是通过select_related来解决。上面的代码直接改造为:

```text
posts_with_user = Post.objects.all().select_related('user')
```

这样产生的语句就是上面的那个JOIN语句。当然ORM还提供了其他类似的方法，比如prefetch_related，又是用来处理其他的问题。

总的来说，ORM给我们提供了便利，但某种程度上也对我们造成了限制，或者说是约束好了。


总结:

其实比较简单,回想起来经常要在列表中显示创建人姓名,这其实就是N+1问题.

我们一般的做法都是

- 手动join,该方法其实也有局限性，如果 user 和 department 不在同一个服务器是不可以连表的。
- 合并查询

合并查询其实就是这样的所谓1+1查询

1. 该方法先查询1次Post列表
    
2. 取出列表中的用户ID组成数组
    
3. 查询步骤2中的用户
    
4. 合并最终数据


该方法对两个表没有限制，在目前微服务盛行的情况下是比较好的一种做法。