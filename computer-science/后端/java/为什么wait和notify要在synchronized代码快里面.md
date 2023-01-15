代码一般是这样

```java

synchronized (object){

  wait();

}


```

因为wait()会释放锁
释放锁,意味着什么,意味着执行wait的时候,当前线程肯定是要拿到锁的.
所以一定要在sync代码块里面!!
