package huang.threadpool;

import org.springframework.scheduling.concurrent.ThreadPoolTaskExecutor;

import java.util.concurrent.ThreadPoolExecutor;

/**
 * @Auther: 宁哥
 * @Date: 2019/2/22 21:31
 * @Description:
 */
public class MyThreadPool
{
    private static ThreadPoolTaskExecutor threadPool = null;

    static
    {
        threadPool = new ThreadPoolTaskExecutor();
        threadPool.setCorePoolSize(10);
        threadPool.setMaxPoolSize(15);
        threadPool.setKeepAliveSeconds(1);
        threadPool.setQueueCapacity(5);
        threadPool.setRejectedExecutionHandler(new ThreadPoolExecutor.CallerRunsPolicy());
        threadPool.initialize();
    }

    public static ThreadPoolTaskExecutor getThreadPool()
    {
        return threadPool;
    }
}
