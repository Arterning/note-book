package com.ibeetl.code.ch01.com.ibeetl.code.ch01.jmh;

import org.openjdk.jmh.annotations.*;
import org.openjdk.jmh.runner.Runner;
import org.openjdk.jmh.runner.RunnerException;
import org.openjdk.jmh.runner.options.Options;
import org.openjdk.jmh.runner.options.OptionsBuilder;

import java.util.List;
import java.util.concurrent.TimeUnit;

/**
 * c.i.c.c.c.i.c.c.j.Ch02OneSizeCheckTest.common          avgt        5  34.571        1.006  ns/op
 * c.i.c.c.c.i.c.c.j.Ch02OneSizeCheckTest.oneSizeCheck    avgt        5  17.315        0.324  ns/op
 */
@BenchmarkMode(Mode.AverageTime)
@Warmup(iterations = 5)
@Measurement(iterations = 5, time = 1, timeUnit = TimeUnit.SECONDS)
@Threads(1)
@Fork(1)
@OutputTimeUnit(TimeUnit.NANOSECONDS)
@State(Scope.Benchmark)
public class Ch02OneSizeCheckTest {

	List<Org> orgs = AreaBuilder.initSize(1);
	public static void main(String[] args) throws RunnerException {
		Options opt = new OptionsBuilder().include(Ch02OneSizeCheckTest.class.getSimpleName()).forks(1).build();
		new Runner(opt).run();

	}

	@Benchmark
	public void oneSizeCheck() {
		AreaBuilder.buildProvinceWithSizeCheck(orgs);
	}

	@Benchmark
	public void common() {
		AreaBuilder.buildProvince(orgs);
	}
	

}
