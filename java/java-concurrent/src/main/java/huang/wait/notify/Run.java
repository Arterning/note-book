package huang.wait.notify;


import org.springframework.scheduling.concurrent.ThreadPoolTaskExecutor;

import java.util.concurrent.ThreadPoolExecutor;

/**
 * @Auther: 宁哥
 * @Date: 2019/2/22 20:56
 * @Description:
 */
public class Run
{
    private static  ThreadPoolTaskExecutor  threadPool = null;
    static {
        threadPool = new ThreadPoolTaskExecutor();
        threadPool.setCorePoolSize(10);
        threadPool.setMaxPoolSize(15);
        threadPool.setKeepAliveSeconds(1);
        threadPool.setQueueCapacity(5);
        threadPool.setRejectedExecutionHandler(new ThreadPoolExecutor.CallerRunsPolicy());
        threadPool.initialize();
    }
    public static void main(String[] args)
    {
        MyList list = new MyList();
        Object lock = new Object();
        RunnableA runnableA = new RunnableA(lock,list);
        RunnbleB runnbleB = new RunnbleB(lock,list);


        //new Thread(runnableA).start();
        //new Thread(runnbleB).start();
        threadPool.execute(runnableA);
        threadPool.execute(runnbleB);
        



    }
}
