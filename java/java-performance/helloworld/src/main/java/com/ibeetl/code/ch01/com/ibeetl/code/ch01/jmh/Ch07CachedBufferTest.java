package com.ibeetl.code.ch01.com.ibeetl.code.ch01.jmh;

import org.openjdk.jmh.annotations.*;
import org.openjdk.jmh.runner.Runner;
import org.openjdk.jmh.runner.RunnerException;
import org.openjdk.jmh.runner.options.Options;
import org.openjdk.jmh.runner.options.OptionsBuilder;

import java.util.List;
import java.util.concurrent.TimeUnit;

/**
 * c.i.c.c.c.i.c.c.j.Ch07CachedBufferTest.cacheBufferBuild    avgt        5  562.344        4.455  ns/op
 * c.i.c.c.c.i.c.c.j.Ch07CachedBufferTest.common              avgt        5  686.118       15.698  ns/op
 */
@BenchmarkMode(Mode.AverageTime)
@Warmup(iterations = 5)
@Measurement(iterations = 5, time = 1, timeUnit = TimeUnit.SECONDS)
@Threads(1)
@Fork(1)
@OutputTimeUnit(TimeUnit.NANOSECONDS)
@State(Scope.Benchmark)
public class Ch07CachedBufferTest {

	List<Org> orgs = AreaBuilder.initSize(20);
	public static void main(String[] args) throws RunnerException {
		Options opt = new OptionsBuilder().include(Ch07CachedBufferTest.class.getSimpleName()).forks(1).build();
		new Runner(opt).run();
	}
	@Benchmark
	public void cacheBufferBuild() {
		AreaBuilder.buildProvinceWithCache(orgs);

	}

	@Benchmark
	public void common() {
		AreaBuilder.buildProvince(orgs);
	}
	

}
