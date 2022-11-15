package huang.wait.notify;

/**
 * @Auther: 宁哥
 * @Date: 2019/2/22 20:49
 * @Description:
 */
public class RunnbleB implements  Runnable
{
    private  Object lock;
    private MyList list;

    public RunnbleB(Object lock, MyList list)
    {
        this.lock = lock;
        this.list = list;
    }

    @Override
    public void run()
    {
        synchronized (lock){
            for (int i = 0; i < 10; i++)
            {
                list.add();
                if(list.size()==5){
                    lock.notify();
                    System.out.println("已经发送通知");
                }
                System.out.println("已经发送了" + (i + 1) + "个元素");
                try
                {
                    Thread.sleep(1000);
                }
                catch (InterruptedException e)
                {
                    e.printStackTrace();
                }

            }
        }
    }
}
