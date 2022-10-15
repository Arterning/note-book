# java-performance

# 介绍
书配套例子,**Java系统性能优化实战** ，程序员的优化宝典。购买地址可以从[京东购买](https://item.jd.com/12742086.html)。书中代码包含了关键注释和结论，也可以直接浏览代码获得知识

> 如果对SpringBoot技术有兴趣，可以购买[<SpringBoot2.0精髓>](https://item.jd.com/12214143.html)或者电子版[看云广场购买](https://www.kancloud.cn/xiandafu/springboot2-in-practice/)


# 例子

## 性能优化

如下代码，拼接字符串，能否找到5处性能优化点

```java
  public String buildProvince(List<Org> orgs){
      StringBuilder sb = new StringBuilder();
      for(Org org:orgs){
        if(sb.length()!=0){
          sb.append(",")
        }
        sb.append(org.getProvinceId());
      }
      return sb.toString();
    }
```

工程提供了9种性能优化的办法可以尝试看看能否找到所有优化办法



## 内卷面试题

下列代码调整是否会导编译或者运行错误？ 多选

1) 删除代码里的 System.out.println("start");
2) 删除代码里的 log.info("start");
3) 修改接口 update(String id,Object o),更改为 update(Object id,Object o); 


