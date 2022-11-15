/**
 * 认识future
 */
package c_026;

import java.util.concurrent.ExecutionException;
import java.util.concurrent.ExecutorService;
import java.util.concurrent.Executors;
import java.util.concurrent.Future;
import java.util.concurrent.FutureTask;
import java.util.concurrent.TimeUnit;

public class T06_Future {
	public static void main(String[] args) throws InterruptedException, ExecutionException {

		//任务返回的结果是Integer FurtureTask传入的是一个Callable对象。
        //FurtureTask就是Callback的一个封装，线程池里面只可以扔FurtureTask，不可以直接扔Callable
		FutureTask<Integer> task = new FutureTask<>(()->{
			TimeUnit.MILLISECONDS.sleep(500);
			return 1000;
		}); //new Callable () { Integer call();}
		
		new Thread(task).start();
		
		System.out.println(task.get()); //这里是阻塞的，main要等待task线程执行完。
		
		//*******************************
		ExecutorService service = Executors.newFixedThreadPool(5);
		Future<Integer> f = service.submit(()->{
			TimeUnit.MILLISECONDS.sleep(500);
			return 1;
		});
		System.out.println(f.get());//f.get()拿到的就是FurtureTask返回的100
		System.out.println(f.isDone());
		
	}
}
