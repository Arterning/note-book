# Java 系统性能优化实战



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



答案：在某些情况下，上面3个都可能会导致运行错误





