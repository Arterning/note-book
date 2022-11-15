/**
 * 认识Executor
 */
package c_026;

import java.util.concurrent.Executor;
/**
 * 我们可以自己实现Executor接口
 */
public class T01_MyExecutor implements Executor{

	public static void main(String[] args) {
		new T01_MyExecutor().execute(()->System.out.println("hello executor"));
	}

	@Override
	public void execute(Runnable command) {
		//new Thread(command).run();
		command.run();
		
	}

}

