package com.ibeetl.code.ch05;

import org.openjdk.jmh.annotations.*;
import org.openjdk.jmh.infra.Blackhole;
import org.openjdk.jmh.runner.Runner;
import org.openjdk.jmh.runner.RunnerException;
import org.openjdk.jmh.runner.options.Options;
import org.openjdk.jmh.runner.options.OptionsBuilder;

import java.util.concurrent.TimeUnit;

/**
 * 关于一个过时优化：“嵌套循环外小内大性能好“，起始都一样，用hole.consume可以验证
 *
 * @author 公众号 java系统优化
 */
@BenchmarkMode(Mode.Throughput)
@Warmup(iterations = 5)
@Measurement(iterations = 5, time = 1, timeUnit = TimeUnit.SECONDS)
@Threads(1)
@Fork(1)
@OutputTimeUnit(TimeUnit.MILLISECONDS)
@State(Scope.Benchmark)
public class StringBufferTest {
  java.lang.String a = "abc";
  java.lang.String b = "eft";

  @Benchmark
  public String builder(Blackhole hole) {
    StringBuilder sb = new StringBuilder();
    for(int i=0;i<20;i++) {
      sb.append(a).append(b);
    }
    return sb.toString();
  }

  @Benchmark
  public String buffer(Blackhole hole) {
    StringBuffer sb = new StringBuffer();
    for(int i=0;i<20;i++){
      sb.append(a).append(b);
    }

    return sb.toString();
  }


  public static void main(String[] args) throws RunnerException {
    Options opt = new OptionsBuilder()
            .include(StringBufferTest.class.getSimpleName())
            .forks(1)
            .build();
    new Runner(opt).run();
  }
}
