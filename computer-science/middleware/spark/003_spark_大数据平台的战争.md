从上面可以看到，HDFS名列前茅。因为存储大数据，最稳定最可靠的就是HDFS。而且Hadoop相关的很多技术，
像MapReduce，HDFS，YARN等等，支持它们的底层文件系统就是HDFS。
同时我们可以看到MapReduce，YARN，Kafka，HBase，Cassandra，Hive，Pig，ZooKeeper，MongoDB, Oozie等这些技术也很热门。

其实我们上面提到的各种技术框架，可以划分为两个领域的战争，一个是偏向底层存储的战争，一个是偏向计算的战争。

偏向存储的战争有关系型数据库和非关系型数据库(Relational vs NoSQL)的战争，它们两者都有各自的应用特点。
关系型数据库最大的特点是事务的一致性，读写操作都是事务的，具有ACID的特点，它在银行这样对一致性有要求的系统中应用广泛。

而非关系型数据库一般对一致性要求不高，但支持高性能并发读写，海量数据访问，在微博、Facebook这类SNS应用中广泛使用。
另外，非关系型数据库内部也有战争，比如说HBase和Cassandra，前者注重一致性(Consistency)和可用性(Availability)，后者提供可用性(Availability)和分区容错性(Partition tolerance)。

Redis和Memcached，它们都是内存内的Key/Value存储，但Redis还支持哈希表，有序集和链表等多种数据结构。MongoDB，CouchDB和Couchbase这三个文档型数据库，MongoDB更适用于需要动态查询的场景，CouchDB偏向于预定义查询，Couchbase比CouchDB有更强的一致性，而且还可以作为Key/Value存储。搜索引擎Solr和Elasticsearch的，它们都是基于Lucene，性能上相近，但是前者在Java/C#开发者中大受欢迎，而后者深受Python/PHP开发者喜爱。

偏向计算的战争有MapReduce和Spark之间的战争，它们之间的特点在下文有更详细介绍。此外还有Spark Streaming和Storm之间的战争等等。

这些战争的赢家是谁呢？它们是Redis,MongoDB，Cassandra，Neo4j和Solr