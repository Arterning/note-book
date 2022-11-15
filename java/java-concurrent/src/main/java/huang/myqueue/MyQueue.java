package huang.myqueue;

import java.util.LinkedList;
import java.util.concurrent.atomic.AtomicInteger;

/**
 * @Auther: 宁哥
 * @Date: 2019/2/22 21:08
 * @Description:
 */
public class MyQueue
{
    private final LinkedList<Object> list = new LinkedList<>();

    private final AtomicInteger count = new AtomicInteger(0);

    private final int maxsize = 5;

    private final int minsize = 0;

    private final Object lock = new Object();

    public void put(Object obj)
    {
        synchronized (lock)
        {
            while (count.get() == maxsize)
            {
                try
                {
                    lock.wait();
                }
                catch (InterruptedException e)
                {
                    e.printStackTrace();
                }
            }
            list.add(obj);
            count.getAndIncrement();
            lock.notifyAll();
        }

    }

    public Object get()
    {
        Object temp = null;
        synchronized (lock){
            while(count.get() == minsize){
                try
                {
                    lock.wait();
                }
                catch (InterruptedException e)
                {
                    e.printStackTrace();
                }
                temp = list.removeFirst();
                lock.notifyAll();
            }
        }
        return temp;
    }

    private int size(){
        return count.get();
    }



}
