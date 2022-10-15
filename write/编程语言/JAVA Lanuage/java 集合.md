![](..\images\搜狗截图20190307234346.png)

上述类图中，实线边框的是实现类，比如ArrayList，LinkedList，HashMap等，折线边框的是抽象类，比如AbstractCollection，AbstractList，AbstractMap等，而点线边框的是接口，比如Collection，Iterator，List等。

  发现一个特点，上述所有的集合类，都实现了Iterator接口，这是一个用于遍历集合中元素的接口，主要包含hasNext(), next(), remove()三种方法。它的一个子接口LinkedIterator在它的基础上又添加了三种方法，分别是add(),previous(),hasPrevious()。也就是说如果是先Iterator接口，那么在遍历集合中元素的时候，只能往后遍历，被遍历后的元素不会在遍历到，通常无序集合实现的都是这个接口，比如HashSet，HashMap；而那些元素有序的集合，实现的一般都是LinkedIterator接口，实现这个接口的集合可以双向遍历，既可以通过next()访问下一个元素，又可以通过previous()访问前一个元素，比如ArrayList。



### Collection



### Map

Map->AbstractMap->HashMap

Map->StoredMap->TreeMap(有序的)



### Set

Set->AbstractSet->HashSet

Set->StoredSet->TreeSet



### List

List->AbstractList->ArrayList

List->AbstractList->LinkedList



### 工具类

Collections

Arrays