# 连接查询



## 继承pojo方式

### 第一步先编写sql

```sql
SELECT orders.*, user.`username` FROM orders,`user` WHERE orders.`user_id`=`user`.id;
```

比如这个sql写好了，我需要查询出user.username字段，当然你可以在order里面增加user属性，这是可以的，但是如果我们只需要查询username就可以了，我们就可以不用修改原来的order pojo,而是扩展一个pojo继承order。

**这一步就是确定我们需要select哪些字段，然后确定pojo，也就是确定resultType**

之前在测试一对多查询的时候，发现连接表查询之查出了teacherId，后来发现原来是自己的sql只select出了student，没有select teacher表的字段，都是这种低级错误。。

### 第二步 编写pojo

```java
//通过此类映射订单和用户查询的结果，让此类继承包括 字段较多的pojo类
public class OrdersCustom extends Orders {

    private String username;
    private String sex;
    private String address;

    public String getUsername() {
        return username;
    }

    public void setUsername(String username) {
        this.username = username;
    }

    public String getSex() {
        return sex;
    }

    public void setSex(String sex) {
        this.sex = sex;
    }

    public String getAddress() {
        return address;
    }

    public void setAddress(String address) {
        this.address = address;
    }

}
```





### 编写mapper.xml和mapper接口

```xml
    <!-- 查询订单关联查询用户信息 -->
    <!--输入就是parameterType 这里没有输入-->
    <!--输出就是resultType-->
    <select id="findOrdersUser" resultType="com.iot.mybatis.po.OrdersCustom">
        SELECT
        orders.*,
        user.username,
        user.sex,
        user.address
        FROM
        orders,
        user
        WHERE orders.user_id = user.id
    </select>
```



## resultMap灵活方式

上面这个问题如果我们使用resultMap的话，也可以这样。

```xml
    <resultMap type="com.iot.mybatis.po.Orders" id="OrdersUserResultMap">
        <!-- 配置映射的订单信息 -->
        <!-- id：指定查询列中的唯一标识，订单信息的中的唯 一标识，如果有多个列组成唯一标识，配置多个id
            column：订单信息的唯一标识列
            property：订单信息的唯一标识列所映射到Orders中哪个属性
          -->
        <id column="id" property="id"/>
        <result column="user_id" property="userId"/>
        <result column="number" property="number"/>
        <result column="createtime" property="createtime"/>
        <result column="note" property="note"/>

        <!-- 配置映射的关联的用户信息 -->
        <!-- association：用于映射关联查询单个对象的信息
        property：要将关联查询的用户信息映射到Orders中哪个属性
         -->
        <association property="user" javaType="com.iot.mybatis.po.User">
            <!-- id：关联查询用户的唯 一标识
            column：指定唯 一标识用户信息的列
            javaType：映射到user的哪个属性
             -->
            <id column="user_id" property="id"/>
            <result column="username" property="username"/>
            <result column="sex" property="sex"/>
            <result column="address" property="address"/>
        </association>
    </resultMap>
```



Order里面添加User对象即可

```java
 //用户信息
 private User user;
```



```xml
    <!-- 查询订单关联查询用户信息 -->
    <select id="findOrdersUserResultMap" resultMap="OrdersUserResultMap">
        SELECT
        orders.*,
        user.username,
        user.sex,
        user.address
        FROM
        orders,
        user
        WHERE orders.user_id = user.id
    </select>
```

