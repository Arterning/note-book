public class jedis操作Redis集群 {
    

    base() {
         // 设置redis集群的节点信息
         Set<HostAndPort> nodes = new HashSet<>();
         nodes.add(new HostAndPort("192.168.1.3", 6379));
         nodes.add(new HostAndPort("192.168.1.4", 6379));
         nodes.add(new HostAndPort("192.168.1.5", 6379));
  
         // 创建jediscluster，可以理解为jedis对象
         JedisCluster cluster = new JedisCluster(nodes);
  
         // 和jedis的使用方式几乎一样
         cluster.set("name", "nameDemo");
  
         // 使用完毕后，不需要释放连接
         // cluster.close();
    }

    /**
     * JedisCluster增加连接池，只需要配置一下连接池即可
     */
    poll() {
        // 设置redis集群的节点信息
        Set<HostAndPort> nodes = new HashSet<>();
        nodes.add(new HostAndPort("192.168.1.3", 6379));
        nodes.add(new HostAndPort("192.168.1.4", 6379));
        nodes.add(new HostAndPort("192.168.1.5", 6379));
 
        // 配置连接池
        JedisPoolConfig jedisPoolConfig = new JedisPoolConfig();
        jedisPoolConfig.setMaxTotal(5);
        jedisPoolConfig.setMaxIdle(3);
        jedisPoolConfig.setMinIdle(2);
 
        // 创建jediscluster，传入节点列表和连接池配置
        JedisCluster cluster = new JedisCluster(nodes, jedisPoolConfig);
 
        // 和jedis的使用方式几乎一样
        cluster.set("name", "nameDemo2121");
 
        // 使用完毕后，不需要释放连接
        // cluster.close();
    }
}
