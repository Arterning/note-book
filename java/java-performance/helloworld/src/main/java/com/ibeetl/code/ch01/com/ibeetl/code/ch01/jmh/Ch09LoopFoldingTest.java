package com.ibeetl.code.ch01.com.ibeetl.code.ch01.jmh;

import org.openjdk.jmh.annotations.*;
import org.openjdk.jmh.runner.Runner;
import org.openjdk.jmh.runner.RunnerException;
import org.openjdk.jmh.runner.options.Options;
import org.openjdk.jmh.runner.options.OptionsBuilder;

import java.util.List;
import java.util.concurrent.TimeUnit;

/**
 * c.i.c.c.c.i.c.c.j.Ch08SubStringTest.common       avgt        5  682.595       28.747  ns/op
 * c.i.c.c.c.i.c.c.j.Ch08SubStringTest.subString    avgt        5  669.746      147.150  ns/op
 */
@BenchmarkMode(Mode.AverageTime)
@Warmup(iterations = 5)
@Measurement(iterations = 5, time = 1, timeUnit = TimeUnit.SECONDS)
@Threads(1)
@Fork(1)
@OutputTimeUnit(TimeUnit.NANOSECONDS)
@State(Scope.Benchmark)
public class Ch09LoopFoldingTest {

	List<Org> orgs = AreaBuilder.initSize(20);
	public static void main(String[] args) throws RunnerException {
		Options opt = new OptionsBuilder().include(Ch09LoopFoldingTest.class.getSimpleName()).forks(1).build();
		new Runner(opt).run();
	}

	@Benchmark
	public void folding() {
		AreaBuilder.buildProvinceWithLoop(orgs);
	}

	@Benchmark
	public void common() {
		AreaBuilder.buildProvince(orgs);
	}

}
