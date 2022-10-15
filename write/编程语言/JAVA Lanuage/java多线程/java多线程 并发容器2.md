## 大纲

> >总结：
> >
> >**1：对于map/set的选择使用**
> >HashMap
> >TreeMap
> >LinkedHashMap
> >
> >Hashtable
> >Collections.sychronizedXXX
> >
> >ConcurrentHashMap   支持高并发下。
> >ConcurrentSkipListMap 支持高并发下。排序跳跃表。
> >
> >**2：队列**
> >ArrayList
> >LinkedList
> >Collections.synchronizedXXX
> >CopyOnWriteList 写时复制列表，写的效率低，写入的时候会复制出一个新的，读的效率很高。因为读取的时候不需要加锁，适合的场景：写的很少，读取的很多的业务场景。
> >
> >**Queue 队列非常重要。**
> >CocurrentLinkedQueue //concurrentArrayQueue 这是并发的加锁的。
> >
> >ConcurrentLinkedQueue是Queue的一个安全实现．Queue中元素按FIFO原则进行排序．采用CAS操作，来保证元素的一致性。
> >
> >
> >
> >BlockingQueue 这是阻塞的容器。
> > LinkedBQ
> > ArrayBQ
> > TransferQueue
> >SynchronusQueue
> >DelayQueue执行定时任务
> >
> >  
> >
> >在Java多线程应用中，队列的使用率很高，多数生产消费模型的首选数据结构就是队列(先进先出)。Java提供的线程安全的Queue可以分为阻塞队列和非阻塞队列，其中阻塞队列的典型例子是BlockingQueue，非阻塞队列的典型例子是ConcurrentLinkedQueue，在实际应用中要根据实际需要选用阻塞队列或者非阻塞队列。
> >
> >不知道你有没有注意到，java.util.concurrent包提供的容器从命名上可以大概分为
> >
> >**Concurrent、CopyOnWrite和Blocking三类**，同样是线程安全容器，可以简单认为：
> >
> >Concurrent类型容器没有CopyOnWrite之类容器相对较重的修改开销。但是，凡事都是有代价的，Concurrent往往提供了较低的遍历一致性。或称之为“弱一致性”，例如，当利用迭代器遍历时，如果容器发生修改，迭代器仍然可以继续进行遍历。
> >
> >与弱一致性对应的，就是其它同步容器常见的行为“fast-fail”，也就是检测到容器在遍历过程中发生了修改，则抛出ConcurrentModificationException，不再继续遍历。
> >
> >弱一致性的另外一个体现是，size等操作准确性是有限的，未必是百分百准确。与此同时，读取的性能具有一定的不确定性。
> >
> >
> >
> >

在多线程环境下，什么样的容器效率比较高？

```java
/**
* HashTable是对整个map对象加锁。Collections.synchronized可以把容器变成同步的。
* ConcurrentHashMap因为锁的粒度比较小，所以他的效率比HashMap要高。
* 如果要在高并发下，而且要排序，TreeMap就GG了，
* 我们使用ConcurrentSkipListMap
*/
public class T01_ConcurrentMap {
	public static void main(String[] args) {
		//Map<String, String> map = new ConcurrentHashMap<>();
		Map<String, String> map = new ConcurrentSkipListMap<>(); //高并发并且排序
		
		//Map<String, String> map = new Hashtable<>();
		//Map<String, String> map = new HashMap<>(); //Collections.synchronizedXXX
		//TreeMap
		Random r = new Random();
		Thread[] ths = new Thread[100];
		CountDownLatch latch = new CountDownLatch(ths.length);
		long start = System.currentTimeMillis();
		for(int i=0; i<ths.length; i++) {
			ths[i] = new Thread(()->{
				for(int j=0; j<10000; j++) map.put("a" + r.nextInt(100000), "a" + r.nextInt(100000));
				latch.countDown();
			});
		}
		
		Arrays.asList(ths).forEach(t->t.start());
		try {
			latch.await();
		} catch (InterruptedException e) {
			e.printStackTrace();
		}
		
		long end = System.currentTimeMillis();
		System.out.println(end - start);
	}
}

```



## Queue

### ConcurrentLinkedQueue

```java

public class T04_ConcurrentQueue {
	public static void main(String[] args) {
		Queue<String> strs = new ConcurrentLinkedQueue<>();
		
		for(int i=0; i<10; i++) {
			strs.offer("a" + i);  //offer是添加
		}
		
		System.out.println(strs);
		
		System.out.println(strs.size());
		
		System.out.println(strs.poll());//poll是拿出来。删掉
		System.out.println(strs.size());
		
		System.out.println(strs.peek());//peek是拿出来不删掉。
		System.out.println(strs.size());
		
		//双端队列Deque
	}
}

```



### LinkedBlockingQueue

LinkedBlockingQueue是无界队列。

```java

public class T05_LinkedBlockingQueue {

    //使用BlockingQueue，相当于自动实现了生产者，消费者。。
	static BlockingQueue<String> strs = new LinkedBlockingQueue<>();

	static Random r = new Random();

	public static void main(String[] args) {
        //启动十个线程put 生产者。
		new Thread(() -> {
			for (int i = 0; i < 100; i++) {
				try {
					strs.put("a" + i); //如果满了，就会阻塞
					TimeUnit.MILLISECONDS.sleep(r.nextInt(1000));
				} catch (InterruptedException e) {
					e.printStackTrace();
				}
			}
		}, "p1").start();

        //启动十个线程get，消费者。
		for (int i = 0; i < 5; i++) {
			new Thread(() -> {
				for (;;) {
					try {
						System.out.println(Thread.currentThread().getName() + " take -" + strs.take()); //如果空了，就会阻塞，等待
					} catch (InterruptedException e) {
						e.printStackTrace();
					}
				}
			}, "c" + i).start();

		}
	}
}

```



### ArrayBlockingQueue

这是一个有界队列，最多装十个。

```java

public class T06_ArrayBlockingQueue {

	static BlockingQueue<String> strs = new ArrayBlockingQueue<>(10);

	static Random r = new Random();

	public static void main(String[] args) throws InterruptedException {
		for (int i = 0; i < 10; i++) {
			strs.put("a" + i);
		}
		
		strs.put("aaa"); //满了就会等待，程序阻塞
		//strs.add("aaa");//add方法，如果满了，会抛出异常。
		//strs.offer("aaa");//如果队列满了，offer不会抛出异常
		//strs.offer("aaa", 1, TimeUnit.SECONDS);
		
		System.out.println(strs);
	}
}

```





### DelayQueue

可以用来做定时执行任务，ScheduledExecutorService就是他实现的。

```java

public class T07_DelayQueue {

	static BlockingQueue<MyTask> tasks = new DelayQueue<>();

	static Random r = new Random();
	
    //往DelayQueue里面添加的对象，需要实现Delayed接口。
	static class MyTask implements Delayed {
		long runningTime;
		
		MyTask(long rt) {
			this.runningTime = rt;
		}

		@Override
		public int compareTo(Delayed o) {
			if(this.getDelay(TimeUnit.MILLISECONDS) < o.getDelay(TimeUnit.MILLISECONDS))
				return -1;
			else if(this.getDelay(TimeUnit.MILLISECONDS) > o.getDelay(TimeUnit.MILLISECONDS)) 
				return 1;
			else 
				return 0;
		}

		@Override
		public long getDelay(TimeUnit unit) {
			
			return unit.convert(runningTime - System.currentTimeMillis(), TimeUnit.MILLISECONDS);
		}
		
		
		@Override
		public String toString() {
			return "" + runningTime;
		}
	}

	public static void main(String[] args) throws InterruptedException {
		long now = System.currentTimeMillis();
		MyTask t1 = new MyTask(now + 1000);//从现在开始1秒钟之后可以在队列中被拿走
		MyTask t2 = new MyTask(now + 2000);//从现在开始2秒钟之后可以在队列中被拿走
		MyTask t3 = new MyTask(now + 1500);//从现在开始1.5秒钟之后可以在队列中被拿走
		MyTask t4 = new MyTask(now + 2500);//从现在开始2.5秒钟之后可以在队列中被拿走
		MyTask t5 = new MyTask(now + 500);
		
		tasks.put(t1);
		tasks.put(t2);
		tasks.put(t3);
		tasks.put(t4);
		tasks.put(t5);
		
		System.out.println(tasks);
		
		for(int i=0; i<5; i++) {
			System.out.println(tasks.take());
		}
	}
}

```



### LinkedTransferQueue

生产者直接把消息发给消费者，适合于并发非常高的情况。

如果找不到消费者，就会阻塞。

```java

public class T08_TransferQueue {
	public static void main(String[] args) throws InterruptedException {
		LinkedTransferQueue<String> strs = new LinkedTransferQueue<>();
		
		/*new Thread(() -> {
			try {
				System.out.println(strs.take());
			} catch (InterruptedException e) {
				e.printStackTrace();
			}
		}).start();*/
		
		strs.transfer("aaa");//会阻塞
		
		strs.put("aaa");//不会阻塞
		

		new Thread(() -> {
			try {
				System.out.println(strs.take());
			} catch (InterruptedException e) {
				e.printStackTrace();
			}
		}).start();
	}
}

```



### SynchronousQueue

SynchronousQueue是容量为0的队列，表示来的任何东西，消费者必须马上消费掉，不然会出问题。因为他容量为0.

SynchronousQueue是一种特殊的LinkedTransferQueue，只是容量为0.

```java
public class T09_SynchronusQueue { //容量为0
   public static void main(String[] args) throws InterruptedException {
      BlockingQueue<String> strs = new SynchronousQueue<>();
      
      new Thread(()->{
         try {
            System.out.println(strs.take());
         } catch (InterruptedException e) {
            e.printStackTrace();
         }
      }).start();
      
      strs.put("aaa"); //阻塞等待消费者消费
      //strs.add("aaa");//这里会异常，因为没有消费者消费。
      System.out.println(strs.size());
   }
}
```



strs.add("aaa");//这里会异常，因为没有消费者消费。

```java
Exception in thread "main" java.lang.IllegalStateException: Queue full
	at java.util.AbstractQueue.add(AbstractQueue.java:98)
	at c_025.T09_SynchronusQueue.main(T09_SynchronusQueue.java:19)
```

