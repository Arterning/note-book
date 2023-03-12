/**
 * HttpServletRequest 没有提供修改heads的方法
 * 但是我们又不能修改HttpServletRequest的源码
 * 所以我们可以使用反射去修改head
 */
public class 反射设置requestHead {
    
    private void reflectSetHeader(HttpServletRequest request, String key, String value){
        Class<? extends HttpServletRequest> requestClass = request.getClass();
        log.info("request实现类={}", requestClass.getName());
        try {
            Field request1 = requestClass.getDeclaredField("request");
            request1.setAccessible(true);
            Object o = request1.get(request);
            Field coyoteRequest = o.getClass().getDeclaredField("coyoteRequest");
            coyoteRequest.setAccessible(true);
            Object o1 = coyoteRequest.get(o);
            Field headers = o1.getClass().getDeclaredField("headers");
            headers.setAccessible(true);
            MimeHeaders o2 = (MimeHeaders)headers.get(o1);
            o2.removeHeader(key);
            o2.addValue(key).setString(value);
        } catch (Exception e) {
            log.info("reflect set header error {}", e);
        }
    }

}
