package com.ibeetl.code.ch05;

import org.openjdk.jmh.annotations.*;
import org.openjdk.jmh.runner.Runner;
import org.openjdk.jmh.runner.RunnerException;
import org.openjdk.jmh.runner.options.Options;
import org.openjdk.jmh.runner.options.OptionsBuilder;

import java.util.concurrent.TimeUnit;

/**
 * 位运算，不过最好兼顾代码可读性，比如
 * <pre>
 *  q = (i * 52429) >>> (16+3);
 * </pre>
 * 其实就是q=i*0.1000000003,这个代码来自JDK Integer.toString
 * @author 公众号 java系统优化
 */
@BenchmarkMode(Mode.AverageTime)
@Warmup(iterations = 5)
@Measurement(iterations = 5, time = 1, timeUnit = TimeUnit.SECONDS)
@Threads(1)
@Fork(1)
@OutputTimeUnit(TimeUnit.NANOSECONDS)
@State(Scope.Benchmark)
public class BitTest2 {
	int a = 22312;

	public static void main(String[] args) throws RunnerException {

		Options opt = new OptionsBuilder().include(BitTest2.class.getSimpleName()).forks(1).build();
		new Runner(opt).run();
	}

	@Benchmark
	public int general() {
		return a/10;
	}

	@Benchmark
	public int bit() {
		return  (a * 52429) >>> (16+3);
	}
}
