## 用内部类实现单例

1.线程安全

2.不用加锁也可以实现懒加载

```java
/**
 * 线程安全的单例模式：
 * 
 * 阅读文章：http://www.cnblogs.com/xudong-bupt/p/3433643.html
 * 
 * 更好的是采用下面的方式，既不用加锁，也能实现懒加载
 * 
 * @author 马士兵
 */
package c_023;

import java.util.Arrays;

public class Singleton {
	
	private Singleton() {
		System.out.println("single");
	}
	
	private static class Inner {
		private static Singleton s = new Singleton();
	}
	
	public static Singleton getSingle() {
		return Inner.s;
	}
	
}

```





## 从一个业务场景说起

有N张火车票，同时有十个窗口对外卖票。

```java

//这个代码肯定是有问题的，因为List不是并发容器。会多卖，还可能会卖重复了
//比如春运卖票出现这种问题肯定是要完蛋的。谢罪的。。。
public class TicketSeller1 {
	static List<String> tickets = new ArrayList<>();
	
	static {
		for(int i=0; i<10000; i++) tickets.add("票编号：" + i);
	}
	
	
	
	public static void main(String[] args) {
		for(int i=0; i<10; i++) {
			new Thread(()->{
				while(tickets.size() > 0) {
					System.out.println("销售了--" + tickets.remove(0));
				}
			}).start();
		}
	}
}

```



```java
/**
 * 换了一个容器Vector,Vector加了synchronized锁。
 * 但是还是有问题，因为判断和操作分离了
 * ticket.size()是原子性的，ticket.remove是原子性的
 * 但是两个原子性的中间是有可能被打断的。。
 * 就算操作A和B都是同步的，但A和B组成的复合操作也未必是同步的，仍然需要自己进行同步
 */
public class TicketSeller2 {
	static Vector<String> tickets = new Vector<>();
	
	
	static {
		for(int i=0; i<1000; i++) tickets.add("票 编号：" + i);
	}
	
	public static void main(String[] args) {
		
		for(int i=0; i<10; i++) {
			new Thread(()->{
				while(tickets.size() > 0) {
					
					try {
						TimeUnit.MILLISECONDS.sleep(10);
					} catch (InterruptedException e) {
						e.printStackTrace();
					}
					
					
					System.out.println("销售了--" + tickets.remove(0));
				}
			}).start();
		}
	}
}

```

```java
/**
 * 有N张火车票，每张票都有一个编号
 * 同时有10个窗口对外售票
 * 请写一个模拟程序
 * 
 * 分析下面的程序可能会产生哪些问题？
 * 重复销售？超量销售？
 * 
 * 使用Vector或者Collections.synchronizedXXX
 * 分析一下，这样能解决问题吗？
 * 
 * 就算操作A和B都是同步的，但A和B组成的复合操作也未必是同步的，仍然需要自己进行同步
 * 就像这个程序，判断size和进行remove必须是一整个的原子操作
 * 
 * @author 马士兵
 */
package c_024;

import java.util.LinkedList;
import java.util.List;
import java.util.concurrent.TimeUnit;

public class TicketSeller3 {
	static List<String> tickets = new LinkedList<>();
	
	
	static {
		for(int i=0; i<1000; i++) tickets.add("票 编号：" + i);
	}
	
	public static void main(String[] args) {
		
		for(int i=0; i<10; i++) {
			new Thread(()->{
				while(true) {
                    //这里上锁。。。。将判断size()和操作remove()加到同一个原子性操作里面去。
					synchronized(tickets) {
						if(tickets.size() <= 0) break;
						
						try {
							TimeUnit.MILLISECONDS.sleep(10);
						} catch (InterruptedException e) {
							e.printStackTrace();
						}
						
						System.out.println("销售了--" + tickets.remove(0));
					}
				}
			}).start();
		}
	}
}

```



```java
/**
* 这里的判断if s==null 和操作poll虽然没有放在一起加锁但是仍然是线程安全的，因为你在判断之后没有再对队列进行操作，就算判断和操作直接被打断，最多也是重新poll一个空值出来。
*/
public class TicketSeller4 {
    //并发的列表实现的队列。在尾巴上加，从头上拿走
	static Queue<String> tickets = new ConcurrentLinkedQueue<>();
	
	static {
		for(int i=0; i<1000; i++) tickets.add("票 编号：" + i);
	}
	
	public static void main(String[] args) {
		
        //启动十个线程，卖票
        //这里没有加锁，但是不会出问题。而且效率是最高的。poll底层是CAS实现的，不是加锁的实现。       
		for(int i=0; i<10; i++) {
			new Thread(()->{
				while(true) {
					String s = tickets.poll();//poll表示往外拿走一个数据。如果没拿着，s为空
					if(s == null) break;
					else System.out.println("销售了--" + s);
				}
			}).start();
		}
	}
}

```

