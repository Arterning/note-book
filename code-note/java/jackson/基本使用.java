/**

<dependency>
  <groupId>com.fasterxml.jackson.core</groupId>
  <artifactId>jackson-databind</artifactId>
  <version>2.9.6</version>
</dependency>

 **/

/**
 *  jackson-databind 依赖 jackson-core 和 jackson-annotations，所以可以只显示地添加jackson-databind依赖，jackson-core 和 jackson-annotations 也随之添加到 Java 项目工程中

 */

 /**
  * Jackson 最常用的 API 就是基于"对象绑定" 的 ObjectMapper：

ObjectMapper可以从字符串，流或文件中解析JSON，并创建表示已解析的JSON的Java对象。 将JSON解析为Java对象也称为从JSON反序列化Java对象。
ObjectMapper也可以从Java对象创建JSON。 从Java对象生成JSON也称为将Java对象序列化为JSON。

  */


  /**
   * 
   * Jackson通过将JSON字段的名称与Java对象中的getter和setter方法进行匹配，将JSON对象的字段映射到Java对象中的属性。 Jackson删除了getter和setter方法名称的“ get”和“ set”部分，并将其余名称的第一个字符转换为小写
   */
public class 基本使用 {

    jsonToObject() {
        try {
            Car car = objectMapper.readValue(carJson, Car.class);      
        } catch (IOException e) {
            e.printStackTrace();
        }
    }

    jsonToObject_v2() {
        try {
            ObjectMapper objectMapper = new ObjectMapper();

            String carJson =
                    "{ \"brand\" : \"Mercedes\", \"doors\" : 4 }";
            Reader reader = new StringReader(carJson);
            
            Car car = objectMapper.readValue(reader, Car.class);
            
        } catch (IOException e) {
            e.printStackTrace();
        }
    }

    jsonToObject_v3() {
        String jsonArray = "[{\"brand\":\"ford\"}, {\"brand\":\"Fiat\"}]";
        ObjectMapper objectMapper = new ObjectMapper();
        Car[] cars2 = objectMapper.readValue(jsonArray, Car[].class);
    }

    jsonToList() {
        String jsonArray = "[{\"brand\":\"ford\"}, {\"brand\":\"Fiat\"}]";
        ObjectMapper objectMapper = new ObjectMapper();
        List<Car> cars1 = objectMapper.readValue(jsonArray, new TypeReference<List<Car>>(){});
    }


    /**
     * Java对象-->JSON
        Jackson ObjectMapper也可以用于从对象生成JSON。 可以使用以下方法之一进行操作：
        writeValue()
        writeValueAsString()
        writeValueAsBytes()
     */

    
}
