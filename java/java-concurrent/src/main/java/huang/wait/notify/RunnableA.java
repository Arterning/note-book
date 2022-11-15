package huang.wait.notify;

/**
 * @Auther: 宁哥
 * @Date: 2019/2/22 20:43
 * @Description:
 */
public class RunnableA implements  Runnable
{

    private  Object lock;
    private MyList list;

    public RunnableA(Object lock, MyList list)
    {
        this.lock = lock;
        this.list = list;
    }

    @Override
    public void run()
    {
        synchronized (lock){

            if(list.size() != 5){
                System.out.println("wait begin" + System.currentTimeMillis());
                try
                {
                    lock.wait();
                }
                catch (InterruptedException e)
                {
                    e.printStackTrace();
                }
                System.out.println("wait end" + System.currentTimeMillis());

            }
        }
    }
}
