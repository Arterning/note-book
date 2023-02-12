/**
 * <dependency>
        <groupId>redis.clients</groupId>
        <artifactId>jedis</artifactId>
        <version>3.2.0</version>
    </dependency>
 */
public class 基本使用 {

    base() {
        // 连接redis
        Jedis jedis = new Jedis("localhost", 6379);
        // Jedis jedis = new Jedis("localhost"); // 默认6379端口
 
        // string类型
        jedis.set("name", "demo");
        jedis.expire("name", 2332);
        String name = jedis.get("name");
 
        // list类型
        jedis.lpush("myList", "hello");
        jedis.rpush("myList", "world");
        String lpopVal = jedis.lpop("myList");
        String rpopVal = jedis.rpop("myList");
 
        // set类型
        jedis.sadd("mySet", "123");
        jedis.sadd("mySet", "456");
        jedis.sadd("mySet", "789");
        jedis.srem("mySet", "789");
        jedis.scard("mySet");
 
        // zset类型
        jedis.zadd("myZset", 99, "X");
        jedis.zadd("myZset", 90, "Y");
        jedis.zadd("myZset", 97, "Z");
        Double zscore = jedis.zscore("myZset", "Z");
 
        // 其他
        jedis.incr("intKey");
        jedis.incrBy("intKey", 5);
        jedis.del("intKey");
 
        // 触发持久化
        // jedis.save();
        // jedis.bgsave()
 
        // 关闭连接
        jedis.close();
    }
    

    pool() {
        // 配置连接池
        JedisPoolConfig config = new JedisPoolConfig();
        config.setMaxTotal(20);
        config.setMaxIdle(10);
        config.setMinIdle(5);
 
        // 创建连接池
        JedisPool jedisPool = new JedisPool(config, "localhost", 6379);
 
        Jedis jedis = jedisPool.getResource();
 
        // 使用jedis进行操作
        jedis.set("name", "otherNameVal");
 
        // 用完之后，一定要手动关闭连接（归还给连接池）
        jedis.close();
    }


}


public class JedisUtils {
 
    private static JedisPool jedisPool;
 
    static {
        // 配置连接池
        JedisPoolConfig config = new JedisPoolConfig();
        config.setMaxTotal(5);
        config.setMaxIdle(3);
        config.setMinIdle(2);
 
        // 创建连接池
        jedisPool = new JedisPool(config, "localhost", 6379);
    }
 
    /**
     * 获取redis连接
     */
    public static Jedis getJedis() {
        return jedisPool.getResource();
    }
}