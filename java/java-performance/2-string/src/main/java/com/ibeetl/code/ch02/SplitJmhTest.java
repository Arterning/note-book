package com.ibeetl.code.ch02;

import cn.hutool.core.util.StrUtil;
import org.openjdk.jmh.annotations.Benchmark;
import org.openjdk.jmh.annotations.BenchmarkMode;
import org.openjdk.jmh.annotations.Fork;
import org.openjdk.jmh.annotations.Measurement;
import org.openjdk.jmh.annotations.Mode;
import org.openjdk.jmh.annotations.OutputTimeUnit;
import org.openjdk.jmh.annotations.Scope;
import org.openjdk.jmh.annotations.State;
import org.openjdk.jmh.annotations.Threads;
import org.openjdk.jmh.annotations.Warmup;
import org.openjdk.jmh.runner.Runner;
import org.openjdk.jmh.runner.RunnerException;
import org.openjdk.jmh.runner.options.Options;
import org.openjdk.jmh.runner.options.OptionsBuilder;

import java.util.ArrayList;
import java.util.List;
import java.util.StringTokenizer;
import java.util.concurrent.TimeUnit;

@BenchmarkMode(Mode.AverageTime)//每次执行平均花费时间
@Warmup(iterations = 5, time = 1) //预热5次调用
@Measurement(iterations = 5, time = 1, timeUnit = TimeUnit.SECONDS) // 执行5此，每次1秒
@Threads(1) //单线程
@Fork(1) //
@OutputTimeUnit(TimeUnit.NANOSECONDS) // 单位：纳秒
@State(Scope.Benchmark) // 共享域
public class SplitJmhTest {

  String str = StrUtil.repeat("hello,world", 1000);
  public static void main(String[] args) throws RunnerException {

    Options opt = new OptionsBuilder().include(SplitJmhTest.class.getSimpleName()).forks(1).build();
    new Runner(opt).run();


  }
  @Benchmark
  public String[] splitByRegex(){
    return str.split(",");

  }

  @Benchmark
  public List<String> splitByTokenizer(){
    StringTokenizer st = new StringTokenizer(str,",");
    List<String> list = new ArrayList<>();
    while (st.hasMoreTokens()) {
      list.add(st.nextToken());
    }
    return list;
  }

  @Benchmark
  public List<String> splitByIndexOfChar(){
    int pos = 0, end;
    List<String> list = new ArrayList<>();
    while ((end = str.indexOf(',', pos)) >= 0) {
      list.add(str.substring(pos, end));
      pos = end + 1;
    }
    if(pos!=str.length()-1){
      list.add(str.substring(pos));
    }
    return list;
  }

  @Benchmark
  public List<String> splitByIndexOfString(){
    int pos = 0, end;
    List<String> list = new ArrayList<>();
    while ((end = str.indexOf(",", pos)) >= 0) {
      list.add(str.substring(pos, end));
      pos = end + 1;
    }
    if(pos!=str.length()-1){
      list.add(str.substring(pos));
    }
    return list;
  }

  @Benchmark
  public List<String> hutoolSplitByChar(){
    return StrUtil.split(str, ',');
  }

  @Benchmark
  public List<String> hutoolSplitByString(){
    return StrUtil.split(str, ',', 0,false, false);
  }
}