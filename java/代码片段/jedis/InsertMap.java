class InsertMap {
    public void insertUser(User user){
        Map<String, String> userProperties = new HashMap<String, String>();
        userProperties.put("username", user.getUsername());
        userProperties.put("firstName", user.getFirstName());
        userProperties.put("lastName", user.getLastName());
        userProperties.put("email", user.getEmail());
        userProperties.put("password", user.getPassword());
        
        jedis.hmset("hashmapkey", userProperties);

        Map<String, String> cache = jedis.hgetAll("hashmapkey");

        //使用hset命令给hash类型设置一个字段
        jedis.hset("hashmapkey", "email", user.getEmail());
 }
}