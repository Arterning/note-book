
它是如何工作的呢？从使用方的角度来看，首先执行器要向服务端注册。那么这里你可能就有疑问了，执行器向服务端注册？怎么注册的？多久注册一次？采用什么通信协议？

注册完了之后，服务端才能知道有哪些执行器，并触发任务调度。那么服务端是如何记录每个任务的触发时机，并完成精准调度的呢？XXL-JOB采用的是Quartz调度框架，本文我打算用时间轮方案来替换。

最后，执行器接收到调度请求，是怎么执行任务的呢？

带着这些问题，我们开启XXL-JOB的探索之旅。我先来说说XXL-JOB项目模块，项目模块很简单，有2个：

- xxl-job-core：这个模块是给执行器依赖的；
- xxl-job-admin：对应架构图中的调度中心；

  

`# 任务调度实现

我们接着来看第二个核心技术点，任务调度。

XXL-JOB采用的是Quartz调度框架，这里我打算向你介绍一下时间轮的实现方案，核心源码如下：


```java
 @Component
 public class JobScheduleHandler {
 ​
     private Thread scheduler;
     private Thread ringConsumer;
     private final Map<Integer, List<Integer>> ring;
     
     @PostConstruct
     public void start() {
         scheduler = new Thread(new JobScheduler(), "job-scheduler");
         scheduler.setDaemon(true);
         scheduler.start();
 ​
         ringConsumer = new Thread(new RingConsumer(), "job-ring-handler");
         ringConsumer.setDaemon(true);
         ringConsumer.start();
     }
     
     class JobScheduler implements Runnable {
         @Override
         public void run() {
             sleep(5000 - System.currentTimeMillis() % 1000);
             while (!schedulerStop) {
                 try {
                     lock.lock();
                     // pre read to ring
                     // read db data to ring
                 } catch (Exception e) {
                     log.error("JobScheduler error", e);
                 } finally {
                     lock.unlock();
                 }
                 sleep(1000);
             }
         }
     }
     
     class RingConsumer implements Runnable {
         @Override
         public void run() {
             sleep(1000 - System.currentTimeMillis() % 1000);
             while (!ringConsumerStop) {
                 try {
                     int nowSecond = Calendar.getInstance().get(Calendar.SECOND);
                     List<Integer> jobIds = ring.remove(nowSecond % 60);
                     // 触发任务调度
                 } catch (Exception e) {
                     log.error("ring consumer error", e);
                 }
                 sleep(1000 - System.currentTimeMillis() % 1000);
             }
         }
     }
 }

```

上述通过两个线程池来实现，`job-scheduler`为预读线程，`job-ring-handler`为时间轮线程。那么时间轮是怎么实现任务的精准调度的呢？