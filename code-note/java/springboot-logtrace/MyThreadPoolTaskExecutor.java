import org.slf4j.MDC;
import org.springframework.scheduling.concurrent.ThreadPoolTaskExecutor;
 
import java.util.concurrent.Callable;
import java.util.concurrent.Future;
 
/**
 * @Author: JCccc
 * @Date: 2022-5-30 11:13
 * @Description:
 */
public final class MyThreadPoolTaskExecutor  extends ThreadPoolTaskExecutor  {
    public MyThreadPoolTaskExecutor() {
        super();
    }
    
    @Override
    public void execute(Runnable task) {
        super.execute(ThreadMdcUtil.wrap(task, MDC.getCopyOfContextMap()));
    }
 
 
    @Override
    public <T> Future<T> submit(Callable<T> task) {
        return super.submit(ThreadMdcUtil.wrap(task, MDC.getCopyOfContextMap()));
    }
 
    @Override
    public Future<?> submit(Runnable task) {
        return super.submit(ThreadMdcUtil.wrap(task, MDC.getCopyOfContextMap()));
    }
}