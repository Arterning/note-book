class TestRedis {

       @Test
    public void testRedis() throws JsonProcessingException {
        JedisPoolConfig config = new JedisPoolConfig();
        config.setMaxTotal(20);
        config.setMaxIdle(10);
        config.setMinIdle(5);

        // 创建连接池
        JedisPool jedisPool = new JedisPool(config, "localhost", 6379, 1000, "123456");

        Jedis resource = jedisPool.getResource();
        resource.select(0);


        String token = resource.get("1607978424434638850:1526525638388416598:token");
        System.out.println("test is " + token);


        String userInfo = resource.get("1607978424434638850:1526525638388416598:userInfo");
        System.out.println("user is" + userInfo);


        ObjectMapper objectMapper = new ObjectMapper();
        AppUserInfo userRoleInfo = null;
        if (userInfo != null) {
            userRoleInfo = objectMapper.readValue(userInfo, AppUserInfo.class);
        }

        System.out.println("userRoleInfo is" + userRoleInfo);
        resource.clos
        config.setMaxTotal(20);
        config.setMaxIdle(10);
        config.setMinIdle(5);

        // 创建连接池
        JedisPool jedisPool = new JedisPool(config, "localhost", 6379, 1000, "123456");

        Jedis resource = jedisPool.getResource();
        resource.select(0);


        String token = resource.get("1607978424434638850:1526525638388416598:token");
        System.out.println("test is " + token);


        String userInfo = resource.get("1607978424434638850:1526525638388416598:userInfo");
        System.out.println("user is" + userInfo);


        ObjectMapper objectMapper = new ObjectMapper();
        AppUserInfo userRoleInfo = null;
        if (userInfo != null) {
            userRoleInfo = objectMapper.readValue(userInfo, AppUserInfo.class);
        }

        System.out.println("userRoleInfo is" + userRoleInfo);
        resource.close();
        jedisPool.close();

    }
}