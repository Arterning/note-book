### **核心组件：**



**1.DispatcherServlet：**前端控制器，是整个流程控制的核心，控制其他组件的执行，统一调度，降低组件之间的耦合性，相当于总指挥。



**2.Handler：**处理器，完成具体业务逻辑，相当于Servlet或Action。



**3.HandlerMapping：**DispatcherServlet接收到请求之后，通过HandlerMapping将不同的请求分发到不同的Handler。



**4.HandlerInterceptor：**处理器拦截器，是一个接口，如果我们需要做一些拦截处理，可以来实现这个接口。



**5.HandlerExecutionChain：**处理器执行链，包括两部分内容：Handler和HandlerInterceptor（系统会有一个默认的HandlerInterceptor，如果需要额外拦截处理，可以添加拦截器设置）。



**6.HandlerAdapter：**处理器适配器，Handler执行业务方法之前，需要进行一系列的操作包括表单数据的验证，数据类型的转换，将表单数据封装到JavaBean等等，这一系列的操作，都是由HandlerAdapter来完成，DispatcherServlet通过HandlerAdapter执行不同的Handler。



**7.ModelAndView：**装载了模型数据和视图信息，作为Handler的处理结果，返回给DispatcherServlet。



**8.ViewResolver：**视图解析器，DispatcherServlet通过它将逻辑视图解析成物理视图，最终将渲染结果响应给客户端。



![](D:\repo\笔记\images\搜狗截图20190302222959.png)





### 注解

#### 





### 数据绑定

![](../../images/搜狗截图20190303123516.png)

#### **HttpMessageConverter**

Http请求响应报文其实都是字符串，当请求报文到java程序会被封装为一个ServletInputStream流，开发人员再读取报文，响应报文则通过ServletOutputStream流，来输出响应报文。

从流中只能读取到原始的字符串报文，同样输出流也是。那么在报文到达SpringMVC / SpringBoot和从SpringMVC / SpringBoot出去，都存在一个字符串到java对象的转化问题。这一过程，在SpringMVC / SpringBoot中，是通过HttpMessageConverter来解决的。HttpMessageConverter接口源码：

```java
public interface HttpMessageConverter<T> {

  boolean canRead(Class<?> clazz, MediaType mediaType);

  boolean canWrite(Class<?> clazz, MediaType mediaType);

  List<MediaType> getSupportedMediaTypes();

  T read(Class<? extends T> clazz, HttpInputMessage inputMessage)
      throws IOException, HttpMessageNotReadableException;

  void write(T t, MediaType contentType, HttpOutputMessage outputMessage)
      throws IOException, HttpMessageNotWritableException;
}
```

下面以一例子来说明，

```
@RequestMapping("/test")
@ResponseBody
public String test(@RequestBody String param) {
  return "param '" + param + "'";
}
```

 在请求进入test方法前，会根据@RequestBody注解选择对应的HttpMessageConverter实现类来将请求参数解析到param变量中**，因为这里的参数是String类型的，所以这里是使用了StringHttpMessageConverter类**，它的canRead()方法返回true，然后read()方法会从请求中读出请求参数，绑定到test()方法的param变量中。

**同理当执行test方法后，由于返回值标识了@ResponseBody，SpringMVC / SpringBoot将使用StringHttpMessageConverter的write()方法，将结果作为String值写入响应报文**，当然，此时canWrite()方法返回true。

处理请求时，由合适的消息转换器将请求报文绑定为方法中的形参对象，在这里同一个对象就有可能出现多种不同的消息形式，如json、xml。同样响应请求也是同样道理。

在Spring中，针对不同的消息形式，有不同的HttpMessageConverter实现类来处理各种消息形式，至于各种消息解析实现的不同，则在不同的HttpMessageConverter实现类中。



```java
	    <!-- 消息转换器 -->
	    <mvc:message-converters register-defaults="true">
	      <bean class="org.springframework.http.converter.StringHttpMessageConverter">
			<property name="supportedMediaTypes" value="text/html;charset=UTF-8"/>
		  </bean>
	      <!-- 阿里fastjson -->
	      <bean class="com.alibaba.fastjson.support.spring.FastJsonHttpMessageConverter4"/>
	    </mvc:message-converters>
```



### 模型数据解析





### 数据转换器





### RestFul架构





### 文件上传下载





### 表单标签库





### 国际化



