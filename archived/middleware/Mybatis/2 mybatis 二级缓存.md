mybatis自己无法实现分布式缓存

所以需要和第三方的缓存框架进行整合，也就是ecache

因为mybatis也不是专门做缓存的，所以他不擅长干这事。





配置cache实现是ecache

```xml
    <cache type="org.mybatis.caches.ehcache.EhcacheCache"/>
```



ehcache配置文件

```xml
<ehcache xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
         xsi:noNamespaceSchemaLocation="../config/ehcache.xsd">
    <diskStore path="F:\develop\ehcache"/>
    <defaultCache
            maxElementsInMemory="1000"
            maxElementsOnDisk="10000000"
            eternal="false"
            overflowToDisk="false"
            timeToIdleSeconds="120"
            timeToLiveSeconds="120"
            diskExpiryThreadIntervalSeconds="120"
            memoryStoreEvictionPolicy="LRU">
    </defaultCache>
</ehcache>

```

