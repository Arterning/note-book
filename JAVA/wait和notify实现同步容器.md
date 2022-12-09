使用wait/notify
这里要注意的是wait通常和while一起使用.

最好的表达就是
wait()完毕之后,开始执行add()的时候,不能保证你是最先获取锁的!有可能在你wait到add()之间,线程B获取到了锁,add()了一个元素进去了
那么你如果用if的话,直接add,那么就有问题了.所以你获取到锁之后,还得判断一下,不能直接add()

```java
/**
 * 面试题：写一个固定容量同步容器，拥有put和get方法，以及getCount方法，
 * 能够支持2个生产者线程以及10个消费者线程的阻塞调用
 *
 **/

public class MyContainer1<T>
{
    //这个List就是共享资源
    //所以锁要使用在和共享资源一个地方，互斥方才都可以访问。synchronized就是this.lock();
    final private LinkedList<T> lists = new LinkedList<>();
    final private int MAX = 10; //最多10个元素
    private int count = 0;


    public synchronized void put(T t)
    {
        while (lists.size() == MAX)
        { //想想为什么用while而不是用if？
            try
            {
                this.wait(); //effective java
            }
            catch (InterruptedException e)
            {
                e.printStackTrace();
            }
        }
        lists.add(t);
        ++count;
        this.notifyAll(); //通知消费者线程进行消费
    }

    public synchronized T get()
    {
        T t = null;
        while (lists.size() == 0)
        {
            try
            {
                this.wait();
            }
            catch (InterruptedException e)
            {
                e.printStackTrace();
            }
        }
        t = lists.removeFirst();
        count--;
        this.notifyAll(); //通知生产者进行生产
        return t;
    }

    public static void main(String[] args)
    {
        MyContainer1<String> c = new MyContainer1<>();

        //启动10个消费者线程
        for (int i = 0; i < 10; i++)
        {
            new Thread(() ->
            {
                for (int j = 0; j < 5; j++) System.out.println(c.get());
            }, "c" + i).start();



        try
        {
            TimeUnit.SECONDS.sleep(2);
        }
        catch (InterruptedException e)
        {
            e.printStackTrace();

        }

        //启动2个生产者线程
        for (int i = 0; i < 2; i++)
        {
            new Thread(() ->
            {
                for (int j = 0; j < 25; j++) c.put(Thread.currentThread().getName() + " " + j);
            }, "p" + i).start();
        }


    }
}

```
