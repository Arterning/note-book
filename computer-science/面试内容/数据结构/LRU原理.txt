用LinkedHashMap实现LRU非常简单 只需要重写一个方法就行

public class LRULinkedHashMap<K, V> extends LinkedHashMap<K, V> {

    private int threshold;

    public LRULinkedHashMap(int threshold) {
        super(16, 0.75f, true);
        this.threshold = threshold;
    }

    @Override
    protected boolean removeEldestEntry(Map.Entry eldest) {
        return size() > threshold;
    }
}



get操作，会将节点移动到队尾。
put操作，如果key不存在，则节点放置队尾，如果节点存在，会将节点移动到队尾。


TreeMap 可以排序 是按照元素大小排序

LinkedHashMap 按照放入顺序排序 实现了双向链表
access-order:访问顺序  当设置为访问顺序时，在get/put的时候调用afterNodeAccess方法，将节点移动到队尾。
insertion-order 插入顺序

LinkedHashMap 可以使元素按插入顺序（insertion-order， 缺省值，重复插入相同元素会重新排序）或 按访问顺序(access-order) 排序。

仅当 removeEldestEntry 为 true 时，afterNodeInsertion 会将链表头的 Node (最早插入的) 删除。在j dk1.8 中， removeEldestEntry 默认返回false，所以，默认情况下 afterNodeInsertion 不生效。




