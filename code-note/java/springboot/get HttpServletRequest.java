class Test {

    /**
    线程安全性分析：
HttpServletRequest 对象作为方法参数，相当于局部变量，是 线程安全 的。

缺点：

代码冗余，每个 Controller 方法中都要有 HttpServletRequest request 参数；
HttpServletRequest 对象的获取必须从 Controller 开始，如果使用 HttpServletRequest 对象的地方在函数调用层级比较深的地方，那么整个函数调用链上的所有方法都要有 HttpServletRequest request 参数。实际上在请求处理的全过程中，HttpServletRequest 对象始终存在，相当于线程内部的一个全局变量。
     */
    @GetMapping("/method1")
    public String method1(HttpServletRequest request) {
        System.out.println("Request URI: " + request.getRequestURI());
        return "Invoke HttpServletRequest by method param.";
    }



    /**
    线程安全性分析：
使用 @Autowired 注解时，Spring 实际注入的并非 HttpServletRequest 对象，而是一个代理 proxy（代理实现参考 AutowireUtils.ObjectFactoryDelegatingInvocationHandler），当方法中需要使用 HttpServletRequest 对象时通过此代理获取，所以虽然 Controller 是个单例类，但通过此方法使用 HttpServletRequest 对象是 线程安全 的。
    */
    @Autowired
    private HttpServletRequest autowiredRequest;



    /**
    手动调用
     */
    HttpServletRequest request =
    ((ServletRequestAttributes) (RequestContextHolder.currentRequestAttributes())).getRequest();
}