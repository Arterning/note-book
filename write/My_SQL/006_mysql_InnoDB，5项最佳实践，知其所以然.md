一、关于count(*)
知识点：MyISAM会直接存储总行数，InnoDB则不会，需要按行扫描。

二、关于全文索引
知识点：MyISAM支持全文索引，InnoDB5.6之前不支持全文索引。

三、关于事务
知识点：MyISAM不支持事务，InnoDB支持事务。

四、关于外键
知识点：MyISAM不支持外键，InnoDB支持外键。
实践：不管哪种存储引擎，在数据量大并发量大的情况下，都不应该使用外键，而建议由应用程序保证完整性。

五、关于行锁与表锁
知识点：MyISAM只支持表锁，InnoDB可以支持行锁。


总结
在大数据量，高并发量的互联网业务场景下，对于MyISAM和InnoDB
有where条件，count(*)两个存储引擎性能差不多
不要使用全文索引，应当使用《索引外置》的设计方案
事务影响性能，强一致性要求才使用事务
不用外键，由应用程序来保证完整性
不命中索引，InnoDB也不能用行锁

https://mp.weixin.qq.com/s?__biz=MjM5ODYxMDA5OQ==&mid=2651961428&idx=1&sn=31a9eb967941d888fbd4bb2112e9602b&chksm=bd2d0d888a5a849e7ebaa7756a8bc1b3d4e2f493f3a76383fc80f7e9ce7657e4ed2f6c01777d&scene=21#wechat_redirect