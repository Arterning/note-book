package com.ibeetl.code.ch01.com.ibeetl.code.ch01.jmh;

import org.openjdk.jmh.annotations.*;
import org.openjdk.jmh.runner.Runner;
import org.openjdk.jmh.runner.RunnerException;
import org.openjdk.jmh.runner.options.Options;
import org.openjdk.jmh.runner.options.OptionsBuilder;

import java.util.List;
import java.util.concurrent.TimeUnit;

/**
 * c.i.c.c.c.i.c.c.j.Ch06StringBufferTest.common        avgt        5  683.379       50.402  ns/op
 * c.i.c.c.c.i.c.c.j.Ch06StringBufferTest.shareChars    avgt        5  650.265       19.100  ns/op
 */
@BenchmarkMode(Mode.AverageTime)
@Warmup(iterations = 5)
@Measurement(iterations = 5, time = 1, timeUnit = TimeUnit.SECONDS)
@Threads(1)
@Fork(1)
@OutputTimeUnit(TimeUnit.NANOSECONDS)
@State(Scope.Benchmark)
public class Ch06StringBufferTest {

	List<Org> orgs = AreaBuilder.initSize(20);
	public static void main(String[] args) throws RunnerException {
		Options opt = new OptionsBuilder().include(Ch06StringBufferTest.class.getSimpleName()).forks(1).build();
		new Runner(opt).run();
	}

	@Benchmark
	public void shareChars() {
		AreaBuilder.buildProvinceWithStringBuffer(orgs);
	}

	@Benchmark
	public void common() {
		AreaBuilder.buildProvince(orgs);
	}

}
