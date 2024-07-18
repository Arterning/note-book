## Executor接口

execute(Runnable run)



## ExecutorService接口

execute

submit 既可以提交Runnable，也可以提交Callable

Callable和Runnable区别：Callable定义的线程会返回值。



## Exectuors

是一个非常常用的工具类，比如Arrays，Collections,操作集合的工具类

线程池在等待的队列，会被扔到任务队列里面去，也就是BlockingQueue

创建线程需要从用户态到内核态，创建线程的开销非常大，所以线程创建了之后，就不要销毁他，直接搞一个池子，把任务丢进去，其实创建一个对象非常耗费资源的都是通过池思想来解决的，比如数据库连接池。io连接池。

我创建一个太耗时了，太慢了，还不如一开始就创建足够用的放在这里，需要的时候拿来用就可以了。类似于大保健和自己找女人的区别。大保健多好，不用花感情，花时间，花精力，直接使用就可以了。就是这个类似的意思。

```java
//通过Exectors创建线程池
//1.固定大小的线程池
ExecetorService service = Executors.newFixedThreadPool(5);
//2.来个任务就开启一个线程，最多支持你机器做多能支持多少线程。线程超过一分钟就销毁掉 alivetime可以指定。
ExecutorService service = Executors.newCachedThreadPool();
//3.固定大小的线程池，保证任务顺序执行。。
ExecutorService service = Executors.newSingleThreadExecutor();
//4.执行定时任务
ExecutorService service = Executors.newScheduledThreadPool(4);
public static void main(String[] args) {
		ScheduledExecutorService service = Executors.newScheduledThreadPool(4);
      //scheduleAtFixedRate() 以固定频率执行某个任务。有四个参数。每隔500毫秒。0表示最开始延时多久。
		service.scheduleAtFixedRate(()->{
			try {
				TimeUnit.MILLISECONDS.sleep(new Random().nextInt(1000));
			} catch (InterruptedException e) {
				e.printStackTrace();
			}
			System.out.println(Thread.currentThread().getName());
		}, 0, 500, TimeUnit.MILLISECONDS);

	}

//5.每个线程都会维护自己的队列，如果一个线程自己队列里面没有线程了，他就会从其他线程的队列里面偷任务来执行
//5.启动的是精灵deamon线程，底层是forkjoin实现的。
ExecutorService service = Executors.newWorkStealingPool();

//6.Fork/Join 线程池。
ForkJoinPool fjp = new ForkJoinPool();

service.shutdown()； //关闭线程

```



## ThreadPoolExecetor

前面几个线程池，都返回的是ThreadPoolExecetor

```java
 return new ThreadPoolExecutor(0, Integer.MAX_VALUE,
                                      60L, TimeUnit.SECONDS,
                                      new SynchronousQueue<Runnable>());
```

如果前面的线程池你用的都不爽，你可以自己new一个，然后指定那些乱七八糟的参数就完事了。

ThreadPoolExecutor

看下构造函数

> >corePoolSize 核心线程池数目大小 这个一般是你cpu有几个核，就定义成几个就完事了。
IO密集型任务
CPU密集型任务,coreSize不超过CPU核心数.
> >
> >maximumPoolSize 最大线程池数目大小
> >
> >keepAliveTime 线程如果没有用最大保持时间
> >
> >unit 前面参数keepAliveTime 的时间单位
> >
> >workQueue 任务队列
> >
> >

```java
public ThreadPoolExecutor(int corePoolSize,
                              int maximumPoolSize,
                              long keepAliveTime,
                              TimeUnit unit,
                              BlockingQueue<Runnable> workQueue) {
        this(corePoolSize, maximumPoolSize, keepAliveTime, unit, workQueue,
             Executors.defaultThreadFactory(), defaultHandler);
    }
```



## 并行结算

```java
/**
 * 线程池的概念
 * nasa 比如nasa研究天文，需要大量计算力，要不要做点贡献？？？这就是多线程。。。
 * 计算从1到20万所有的质数 对于这种耗时的操作，一个线程肯定是很慢的，所以我们可以开启多个线程来计算。
 * 一般你的cpu有几个核，你线程池就不该少于你的核数。
 */
package c_026;

import java.util.ArrayList;
import java.util.List;
import java.util.concurrent.Callable;
import java.util.concurrent.ExecutionException;
import java.util.concurrent.ExecutorService;
import java.util.concurrent.Executors;
import java.util.concurrent.Future;

public class T07_ParallelComputing {
   public static void main(String[] args) throws InterruptedException, ExecutionException {
      long start = System.currentTimeMillis();
      getPrime(1, 200000);
      long end = System.currentTimeMillis();
      System.out.println(end - start);

      final int cpuCoreNum = 4;

      ExecutorService service = Executors.newFixedThreadPool(cpuCoreNum);

       //1到二十万数字分层四个任务，注意不是均分的，因为后面数字越大，计算时间越长，所以后面要分少一点。
      MyTask t1 = new MyTask(1, 80000); //1-5 5-10 10-15 15-20
      MyTask t2 = new MyTask(80001, 130000);
      MyTask t3 = new MyTask(130001, 170000);
      MyTask t4 = new MyTask(170001, 200000);

      Future<List<Integer>> f1 = service.submit(t1);
      Future<List<Integer>> f2 = service.submit(t2);
      Future<List<Integer>> f3 = service.submit(t3);
      Future<List<Integer>> f4 = service.submit(t4);

      start = System.currentTimeMillis();
      f1.get();
      f2.get();
      f3.get();
      f4.get();
      end = System.currentTimeMillis();
      System.out.println(end - start);
   }

   static class MyTask implements Callable<List<Integer>> {
      int startPos, endPos;

      MyTask(int s, int e) {
         this.startPos = s;
         this.endPos = e;
      }

      @Override
      public List<Integer> call() throws Exception {
         List<Integer> r = getPrime(startPos, endPos);
         return r;
      }

   }

   static boolean isPrime(int num) {
      for(int i=2; i<=num/2; i++) {
         if(num % i == 0) return false;
      }
      return true;
   }

   static List<Integer> getPrime(int start, int end) {
      List<Integer> results = new ArrayList<>();
      for(int i=start; i<=end; i++) {
         if(isPrime(i)) {results.add(i);}
      }

      return results;
   }
}
```





## Fork Join Pool

fork/join有点类似于map/reduce

![](D:\repo\笔记\images\搜狗截图20190306204614.png)

他的思想是这样的，就是如果一个任务非常大，非常耗时，可以把它分成多个小的task来执行，其实就是分而治之。然后小task执行完毕后，把结果做一个汇总。比如下面这个任务，对一百万个数进行求和操作。

```java
package c_026;

import java.io.IOException;
import java.util.Arrays;
import java.util.Random;
import java.util.concurrent.ForkJoinPool;
import java.util.concurrent.RecursiveTask;

public class T12_ForkJoinPool {
	static int[] nums = new int[1000000];
	static final int MAX_NUM = 50000;
	static Random r = new Random();

	static {
		for(int i=0; i<nums.length; i++) {
			nums[i] = r.nextInt(100);
		}

		System.out.println(Arrays.stream(nums).sum()); //stream api
	}

	/*
	static class AddTask extends RecursiveAction {

		int start, end;

		AddTask(int s, int e) {
			start = s;
			end = e;
		}

		@Override
		protected void compute() {

			if(end-start <= MAX_NUM) {
				long sum = 0L;
				for(int i=start; i<end; i++) sum += nums[i];
				System.out.println("from:" + start + " to:" + end + " = " + sum);
			} else {

				int middle = start + (end-start)/2;

				AddTask subTask1 = new AddTask(start, middle);
				AddTask subTask2 = new AddTask(middle, end);
				subTask1.fork();
				subTask2.fork();
			}


		}

	}
	*/

    /**
    * 必须继承递归任务RecurisveTask，为什么叫做递归，因为就是把任务不断切分，这个过程就是递归的。
    * RecurisveTask是继承的顶层父类ForkJoinTask 才可以扔到ForkJoinPool里面。
    */
	static class AddTask extends RecursiveTask<Long> {

		private static final long serialVersionUID = 1L;
		int start, end;

		AddTask(int s, int e) {
			start = s;
			end = e;
		}

		@Override
		protected Long compute() {

			if(end-start <= MAX_NUM) {//如果小于五万，就不再划分子任务了，直接返回结果
				long sum = 0L;
				for(int i=start; i<end; i++) sum += nums[i];
				return sum;
			}

            //否则再划分任务。
			int middle = start + (end-start)/2;

            //划分task
			AddTask subTask1 = new AddTask(start, middle);
			AddTask subTask2 = new AddTask(middle, end);

            //fork 然后 join
			subTask1.fork();
			subTask2.fork();

			return subTask1.join() + subTask2.join();
		}

	}

	public static void main(String[] args) throws IOException {
		ForkJoinPool fjp = new ForkJoinPool();
		AddTask task = new AddTask(0, nums.length);
		fjp.execute(task);
		long result = task.join();
		System.out.println(result);

		//System.in.read();

	}
}

```
