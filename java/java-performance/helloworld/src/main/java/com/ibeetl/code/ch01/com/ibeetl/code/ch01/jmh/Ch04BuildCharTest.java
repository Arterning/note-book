package com.ibeetl.code.ch01.com.ibeetl.code.ch01.jmh;

import org.openjdk.jmh.annotations.*;
import org.openjdk.jmh.runner.Runner;
import org.openjdk.jmh.runner.RunnerException;
import org.openjdk.jmh.runner.options.Options;
import org.openjdk.jmh.runner.options.OptionsBuilder;

import java.util.List;
import java.util.concurrent.TimeUnit;

/**
 * Benchmark                                        Mode  Samples    Score  Score error  Units
 * c.i.c.c.c.i.c.c.j.Ch04BuildCharTest.charBuild    avgt        5  531.824       32.604  ns/op
 * c.i.c.c.c.i.c.c.j.Ch04BuildCharTest.common       avgt        5  685.824       33.542  ns/op
 */
@BenchmarkMode(Mode.AverageTime)
@Warmup(iterations = 5)
@Measurement(iterations = 5, time = 1, timeUnit = TimeUnit.SECONDS)
@Threads(1)
@Fork(1)
@OutputTimeUnit(TimeUnit.NANOSECONDS)
@State(Scope.Benchmark)
public class Ch04BuildCharTest {

	List<Org> orgs = AreaBuilder.initSize(20);
	public static void main(String[] args) throws RunnerException {
		Options opt = new OptionsBuilder().include(Ch04BuildCharTest.class.getSimpleName()).forks(1).build();
		new Runner(opt).run();

	}
	@Benchmark
	public void charBuild() {
		AreaBuilder.buildProvinceWithChar(orgs);
	}

	@Benchmark
	public void common() {
		AreaBuilder.buildProvince(orgs);
	}

}
