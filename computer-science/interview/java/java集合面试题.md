Java集合的快速失败机制 “fail-fast”？


怎么确保一个集合不能被修改？
可以使用 Collections. unmodififiableCollection(Collection c) 方法来创建一个只读集合，这样改变集合的任何操作都会抛出 Java. lang. UnsupportedOperationException 异常。 

如何边遍历边移除 Collection 中的元素？
使用Iterator.remove()


 如何实现数组和 List 之间的转换？
数组转 List：使用 Arrays. asList(array) 进行转换。 
List 转数组：使用 List 自带的 toArray() 方法。 


comparable 和 comparator的区别？
comparable接口实际上是出自java.lang包，它有一个 compareTo(Object obj)方法用来排序
comparator接口实际上是出自 java.util 包，它有一个compare(Object obj1, Object obj2)方法用来排序 


ConcurrentHashMap 底层具体实现知道吗？实现原理是什么？
JDK7是分段锁
JDK8是CAS配合Synchronized
如果没有hash冲突就直接CAS插入
只要hash不冲突 就不会产生并发
如果存在hash冲突，就加synchronized锁来保证线程安全，这里有两种情况，一种是链表形式就直接遍历到尾端插入，一种是红黑树就按照红黑树结构插入，
JDK1.8 以后的 HashMap 在解决哈希冲突时有了较大的变化，当链表长度大于阈值（默认为8）时，将链表转化为红黑树，以减少搜索时间。Hashtable 没有这样的机制。 


为什么HashMap中String、Integer这样的包装类适合作为K？
因为他们都重写了equals hashcode方法


HashMap是怎么解决哈希冲突的？
数组结合链表的方式


HashMap的扩容操作是怎么实现的？


HashSet如何检查重复？HashSet是如何保证数据不可重复的？
向HashSet 中add ()元素时，判断元素是否存在的依据，不仅要比较hash值，同时还要结合equles 方法比较。 


说一下 HashSet 的实现原理？
HashSet 是基于 HashMap 实现的，HashSet的值存放于HashMap的key上，HashMap的value统一为present，因此 HashSet 的实现比较简单，相关 HashSet 的操作，基本上都是直接调用底层HashMap 的相关方法来完成，HashSet 不允许重复的值。



