package com.cyh.fakeaio;

import java.util.concurrent.ArrayBlockingQueue;
import java.util.concurrent.ExecutorService;
import java.util.concurrent.ThreadPoolExecutor;
import java.util.concurrent.TimeUnit;

public class TimeServerHandlerExecutePool {

    private ExecutorService executorService;

    public TimeServerHandlerExecutePool(int maxPoolSize, int queueSize) {
        int coresize = Runtime.getRuntime().availableProcessors();
        ArrayBlockingQueue<Runnable> queue = new ArrayBlockingQueue<>(queueSize);
        this.executorService = new ThreadPoolExecutor(coresize, maxPoolSize, 120L, TimeUnit.SECONDS, queue);
    }

    public void execute (java.lang.Runnable task) {
        executorService.execute(task);
    }
}
