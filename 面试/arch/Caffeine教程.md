
缓存(Cache)在代码世界中无处不在。从底层的CPU多级缓存，到客户端的页面缓存，处处都存在着缓存的身影。缓存从本质上来说，是一种空间换时间的手段，通过对数据进行一定的空间安排，使得下次进行数据访问时起到加速的效果。

就Java而言，其常用的缓存解决方案有很多，例如数据库缓存框架EhCache，分布式缓存Memcached等，这些缓存方案实际上都是为了提升吞吐效率，避免持久层压力过大。

对于常见缓存类型而言，可以分为本地缓存以及分布式缓存两种，Caffeine就是一种优秀的本地缓存，而Redis可以用来做分布式缓存



Caffeine官方：

> - [https://github.com/ben-manes/caffeine](https://github.com/ben-manes/caffeine)

Caffeine是基于Java 1.8的高性能本地缓存库，由Guava改进而来，而且在Spring5开始的默认缓存实现就将Caffeine代替原来的Google Guava，官方说明指出，其缓存命中率已经接近最优值。实际上Caffeine这样的本地缓存和ConcurrentMap很像，即支持并发，并且支持O(1)时间复杂度的数据存取。二者的主要区别在于：

- ConcurrentMap将存储所有存入的数据，直到你显式将其移除；
- Caffeine将通过给定的配置，自动移除“不常用”的数据，以保持内存的合理占用。

因此，一种更好的理解方式是：Cache是一种带有存储和移除策略的Map。


#### [1、缓存加载策略](https://mp.weixin.qq.com/s?__biz=MzUzMTA2NTU2Ng==&mid=2247487551&idx=1&sn=18f64ba49f3f0f9d8be9d1fdef8857d9&scene=21#wechat_redirect)

##### 1.1 Cache手动创建

最普通的一种缓存，无需指定加载方式，需要手动调用`put()`进行加载。需要注意的是`put()`方法对于已存在的key将进行覆盖，这点和Map的表现是一致的。在获取缓存值时，如果想要在缓存值不存在时，原子地将值写入缓存，则可以调用`get(key, k -> value)`方法，该方法将避免写入竞争。调用`invalidate()`方法，将手动移除缓存。

在多线程情况下，当使用`get(key, k -> value)`时，如果有另一个线程同时调用本方法进行竞争，则后一线程会被阻塞，直到前一线程更新缓存完成；而若另一线程调用`getIfPresent()`方法，则会立即返回null，不会被阻塞。

```java
Cache<Object, Object> cache = Caffeine.newBuilder()
          //初始数量
          .initialCapacity(10)
          //最大条数
          .maximumSize(10)
          //expireAfterWrite和expireAfterAccess同时存在时，以expireAfterWrite为准
          //最后一次写操作后经过指定时间过期
          .expireAfterWrite(1, TimeUnit.SECONDS)
          //最后一次读或写操作后经过指定时间过期
          .expireAfterAccess(1, TimeUnit.SECONDS)
          //监听缓存被移除
          .removalListener((key, val, removalCause) -> { })
          //记录命中
          .recordStats()
          .build();
  cache.put("1","张三");
  //张三
  System.out.println(cache.getIfPresent("1"));
  //存储的是默认值
  System.out.println(cache.get("2",o -> "默认值"));
```



##### 1.2 Loading Cache自动创建

LoadingCache是一种自动加载的缓存。其和普通缓存不同的地方在于，当缓存不存在/缓存已过期时，若调用`get()`方法，则会自动调用`CacheLoader.load()`方法加载最新值。调用`getAll()`方法将遍历所有的key调用get()，除非实现了`CacheLoader.loadAll()`方法。使用LoadingCache时，需要指定CacheLoader，并实现其中的`load()`方法供缓存缺失时自动加载。

在多线程情况下，当两个线程同时调用`get()`，则后一线程将被阻塞，直至前一线程更新缓存完成。


```java
LoadingCache<String, String> loadingCache = Caffeine.newBuilder()
        //创建缓存或者最近一次更新缓存后经过指定时间间隔，刷新缓存；refreshAfterWrite仅支持LoadingCache
        .refreshAfterWrite(10, TimeUnit.SECONDS)
        .expireAfterWrite(10, TimeUnit.SECONDS)
        .expireAfterAccess(10, TimeUnit.SECONDS)
        .maximumSize(10)
        //根据key查询数据库里面的值，这里是个lamba表达式
        .build(key -> new Date().toString());
```


##### 1.3 Async Cache异步获取

AsyncCache是Cache的一个变体，其响应结果均为`CompletableFuture`，通过这种方式，AsyncCache对异步编程模式进行了适配。默认情况下，缓存计算使用`ForkJoinPool.commonPool()`作为线程池，如果想要指定线程池，则可以覆盖并实现`Caffeine.executor(Executor)`方法。`synchronous()`提供了阻塞直到异步缓存生成完毕的能力，它将以Cache进行返回。

在多线程情况下，当两个线程同时调用`get(key, k -> value)`，则会返回同一个`CompletableFuture`对象。由于返回结果本身不进行阻塞，可以根据业务设计自行选择阻塞等待或者非阻塞。

```java
AsyncLoadingCache<String, String> asyncLoadingCache = Caffeine.newBuilder()
        //创建缓存或者最近一次更新缓存后经过指定时间间隔刷新缓存；仅支持LoadingCache
        .refreshAfterWrite(1, TimeUnit.SECONDS)
        .expireAfterWrite(1, TimeUnit.SECONDS)
        .expireAfterAccess(1, TimeUnit.SECONDS)
        .maximumSize(10)
        //根据key查询数据库里面的值
        .buildAsync(key -> {
            Thread.sleep(1000);
            return new Date().toString();
        });
//异步缓存返回的是CompletableFuture
CompletableFuture<String> future = asyncLoadingCache.get("1");
future.thenAccept(System.out::println);
```


