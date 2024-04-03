
  最近遇到一个需求，要求监听mysql数据库的数据变化，并通过websocket实时将数据传回给前端。其难点主要在于监听数据变化，本人经查阅资料总结了四种监听数据库的方法，并选择其中一种作为我解决需求的办法详细给大家介绍。

##   一、监听数据库的方式

###   1.轮询方式

   这种方式就是定期去查询数据库，通过查询结果与上一次查询结果进行对比，来判断要监听的字段是否发生变化，在SpringBoot中可以通过写定时任务来实现。这种方式实现起来十分简单，但是缺点也很明显，就是难以确定一个合适的时间去定期查看，时间设置的太短会浪费资源，时间设置的太长又会有较高的延迟。

###  2.触发器方式

  数据库的触发器可以在数据库进行数据插入或更新时，将事件信息传递给后端，实时性非常好，但是触发器的开销很大，很容易导致性能问题。

###  3.监听binlog日志

  binlog日志是mysql数据库自带的记录数据库所有修改操作的日志文件，通过解析binlog日志可以分析出数据库的数据变化，实时性也非常好，并且不需要额外添加触发器。但是binlog日志是数据库级别的，想要精确到表需要对binlog日志进行解析，实现起来较为复杂。另外，binlog日志有丢失和损坏的风险，可能导致无法监听数据库的数据变化。


###  4.使用开源工具

  网上有一些开源工具可以监听数据库的数据变化，如阿里巴巴的canal工具可以通过监听binlog日志来获取数据库的数据变化情况。不过canal提供的功能很多，对一些项目来说可能较为复杂。




## 二、监听binlog日志的详细实现

 综合比较下来，本人选择了监听binlog日志的方式来实现监听数据库的数据变化。下面介绍在SpringBoot框架下如何实现监听binlog日志。

 首先引入maven依赖

```java
<!-- binlog监听 -->
<dependency>
    <groupId>com.github.shyiko</groupId>
	<artifactId>mysql-binlog-connector-java</artifactId>
	<version>0.21.0</version>
</dependency>
```

接下来在yml配置

```java
binlog:
  config:
    host: 数据库主机号
    port: 数据库端口号
    username: 用户名
    password: 密码
```

config配置类

```java
@Configuration
@ConfigurationProperties(prefix = "binlog.config")//读取上面配置的yml
@Data
public class BinlogConfig {
    String host;
    int port;
    String username;
    String password;
}
```


主体代码

```java
@Slf4j
@Component
public class BinLogRunner implements CommandLineRunner {
    @Autowired
    private BinlogConfig config;
    @Resource
    private WebSocket webSocket;

    private HashMap<String,Long> listenerMap  =new HashMap<>();

    //需要监控的数据表 schema.tableName
    private List<String> listenerTableList = Arrays.asList("schema.tableName");

    private String[] column={"你想监听的表的列名的集合"};
    
    //给监听的数据表一个tableId对应
    private void putListenerMap(String database,String table , Long tableId){
        String key = database+"."+table;
        if(listenerTableList.contains(key)){
            listenerMap.put(key,tableId);
        }
    }

    //通过tableId判断是否是要监听的表
    private String getTableName(Long tableId){
        for (String key : listenerMap.keySet()) {
            Long aLong = listenerMap.get(key);
            if(aLong.equals(tableId)){
                return key;
            }
        }
        return null;
    }

    //监听数据库
    @Override
    public void run(String... args) throws Exception {
        //获取数据库信息
        String host = config.getHost();
        int port = config.getPort();
        String username = config.getUsername();
        String password = config.getPassword();
        if(CollectionUtils.isEmpty(listenerTableList)) {
            return;
        }
        BinaryLogClient client = new BinaryLogClient(host, port, username, password);
        //进行监听
        client.registerEventListener(event -> {
            EventType eventType = event.getHeader().getEventType();
            //binlog在记录操作前首先会记录表的结构等信息
            if (eventType == EventType.TABLE_MAP) {
                TableMapEventData tableData = event.getData();
                String db = tableData.getDatabase();
                String table = tableData.getTable();
                long tableId = tableData.getTableId();
                // 第一遍  会显示数据库  第二遍有tableId
                putListenerMap(db,table,tableId);
            }

             // 如果日志是更新记录
            if(EventType.isUpdate(eventType)){
                long tableId = 0L;
                String tableName = null;
                UpdateRowsEventData update = event.getData();
                tableId = update.getTableId();
                tableName = getTableName(tableId);
                if(tableName == null){
                    return;
                }
                //存储更新前后的数据，但是不会给列名，只有列的值
                List<Map.Entry<Serializable[], Serializable[]>> rows=update.getRows();
                build(rows.get(0).getValue());
            }else{
                return;
            }
        });
        client.connect();
    }

    //对更新后的数据进行解析，并给出列名的对应
    private void build(Serializable[] after){
        Map<String,String> map=new HashMap<>();
        for(int i=0;i<13;i++){
            if(after[i]==null){
                map.put(column[i],null);
                continue;
            }
            map.put(column[i],after[i].toString());
        }
        webSocket.sendAllMessage(JSONObject.toJSONString(map));//websocket的广播消息，本文没给出websocket的配置
    }
}
```

最后的实现效果就是每当监听的表发生变化时就会向前端返回更新后的数据

```json
{"id":"108","uniqueId":"yoyi20230714073915643829","process":"19.0","city":"长春市","tifPath":"长春市","model":"tree_default","inferType":"tree","status":"success"}
```