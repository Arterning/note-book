package com.ibeetl.code.ch05;


import org.openjdk.jmh.annotations.*;
import org.openjdk.jmh.infra.Blackhole;
import org.openjdk.jmh.runner.Runner;
import org.openjdk.jmh.runner.RunnerException;
import org.openjdk.jmh.runner.options.Options;
import org.openjdk.jmh.runner.options.OptionsBuilder;

import java.util.concurrent.TimeUnit;

/**
 * 测试StringBuffer和StringBuilder
 */
@BenchmarkMode(Mode.AverageTime)
@Warmup(iterations = 5)
@Measurement(iterations = 5, time = 1, timeUnit = TimeUnit.SECONDS)
@Threads(1)
@Fork(1)
@OutputTimeUnit(TimeUnit.MILLISECONDS)
@State(Scope.Benchmark)
public class ForDeadCodeTest {

  @Benchmark
  public long test(Blackhole hole) {
    long startTime = System.nanoTime();
    int i = 0, j = 0;
    for (i = 0; i < 100_000_00; i++) {
      for (j = 0; j < 10; j++) {
        hole.consume(j);
      }
      hole.consume(i);
    }

    Long endTime = System.nanoTime();
    return endTime - startTime + i + j;
  }


  @Benchmark
  public long tes2(Blackhole hole) {
    long startTime = System.nanoTime();
    int i = 0, j = 0;
    for (i = 0; i < 10; i++) {
      for (j = 0; j < 100_000_00; j++) {
        hole.consume(j);
      }
      hole.consume(i);
    }
    Long endTime = System.nanoTime();
    return endTime - startTime + i + j;
  }

  /**
   * 上面的测试，如果不用hole.consume，那么性能跟base一样，JIT优化了
   *
   * @param hole
   * @return
   */
  @Benchmark
  public long base(Blackhole hole) {
    long startTime = System.nanoTime();

    Long endTime = System.nanoTime();
    return endTime - startTime;
  }


  public static void main(String[] args) throws RunnerException {
    Options opt = new OptionsBuilder()
            .include(ForDeadCodeTest.class.getSimpleName())
            .forks(1)
            .build();
    new Runner(opt).run();
  }
}
