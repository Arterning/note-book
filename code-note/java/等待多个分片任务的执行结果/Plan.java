/**
 * 
 * 开发中遇到的需求
 * 开发过程中有这样一个需求，主线程有一个步骤需要拆分为多个任务线程，然后主线程等待多个任务线程全部执行完后，收集任务线程的结果，然后继续主线程的逻辑。其中要求可以控制这段等待的总时间，即如果超过一定时间还没有完成所有的任务，就直接收集已经结束的线程运算的结果，然后继续主线程的逻辑。
 * 
 * 这算是一个比较常见的通过多线程来加快任务执行的场景（在具体业务里面分片任务涉及IO阻塞，所以同步执行下会比较慢）。
 * 
 * 初始的想法与实现
 * 在分割任务之前已知总的任务数有多少，为了满足限定总体时长，我选择了CountDownLatch作为锁等待所有线程执行完毕。为了得到线程的执行结果，我使用了Futrue获取。我构造的线程类如下：
 */
public class WorkThread implements Callable<Object> {
    /** 控制任务完成的闭锁 */
    private CountDownLatch finishLatch;

    public WorkThread(CountDownLatch finishLatch) {
        this.finishLatch = finishLatch;
    }

    /**
     * @see java.util.concurrent.Callable#call()
     */
    public Object call() throws Exception {
        try {
            // TODO Work
            return result;
        } finally {
            finishLatch.countDown();
        }
    }
}

public class ManiThread {
    final CountDownLatch latch = new CountDownLatch(n);
    final ThreadPoolExecutors executors = Executors.newCachedThreadPool();

    public static void main(String[] args) {
        List<Future<Object>> futureList = new ArrayList<Future<Object>>();
        // 添加执行异步任务
        for (int i = 0; i < n; i++) {
            futureList.add(executors.submit(new WorkThread(latch)));
        }

        try {
            boolean res = latch.await(MAX_EXCUTE_TIME, TimeUnit.MILLISECONDS);
            for (Future<Object> future : futureList) {
                if (future.isDone()) {
                    try {
                        // 有isDone()判断，肯定已完成
                        Object result = future.get();
                        // handle the result
                    } catch (Exception e) {
                    }
                } else {
                    future.cancel(true);
                }
            }
        } catch (InterruptedException e) {
        }
    }

}

/**
 * 
 * 问题
在实际测试时发现程序总是会漏掉一些子任务，漏掉的概率和被漏的任务是随机的，便觉得很奇怪。

在Debug时发现有部分任务在时间足够的情况下进入了future.cancel(true)分支 于是意识到了问题

线程类任务完成后Future的isDone()状态的更新操作是由线程池控制。
在主线程中，等到countdown为0时立即开始下面的逻辑，而future的完成状态由另一个线程控制
因此已完成的任务也有可能没有被及时的更改为已完成的状态而让主线程执行了cancel的分支。

简单来说就是future的完成状态由另一个线程控制的
就算任务已经完成了 但是future.isDone仍然可能为false

我觉得这怎么也得是future.isDone方法的设计有问题啊？？
 */


/**
 * 
 * 解决方案
1. 去掉future.isDone()判断，直接用future.get()的时限版获取结果，设定一个较短的时间。但是这种写法其实还是没有满足控制整体任务执行超时的限制，因为存在未完成的任务的可能，尽管设定了一个较短的时间，也有可能因为较长的任务队列使得总体运行时间延长。但是有一个好处，如果超时未获取到结果，就可以采取cancel的策略。

2. 利用一个线程安全的List存放结果，取代Future的结构。即将线程类修改为如下。

需要保证对resultList的修改操作是线程安全的。比如采用：
List<Object> resultList = Collections
            .synchronizedList(new ArrayList<Object>());
 */
public class WorkThread implements Runnable {
    /** 控制任务完成的闭锁 */
    private CountDownLatch finishLatch;
    /** 存放结果的同步List */
    private List<Object> resultList;
    public WorkThread(List<Object> resultList, CountDownLatch finishLatch) {
        this.resultList = resultList;
        this.finishLatch = finishLatch;
    }
    /**
     * @see java.util.concurrent.Callable#call()
     */
    public void run() {
        try {
            //TODO Work
            resultList.add(result);
        } finally {
            finishLatch.countDown();
        }
    }
}