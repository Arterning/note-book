package com.ibeetl.code.ch01.com.ibeetl.code.ch01.jmh;

import org.openjdk.jmh.annotations.*;
import org.openjdk.jmh.runner.Runner;
import org.openjdk.jmh.runner.RunnerException;
import org.openjdk.jmh.runner.options.Options;
import org.openjdk.jmh.runner.options.OptionsBuilder;

import java.util.ArrayList;
import java.util.List;
import java.util.concurrent.TimeUnit;

/**
 * c.i.c.c.c.i.c.c.j.Ch01EmptyCheckTest.common        avgt        5  3.996        0.127  ns/op
 * c.i.c.c.c.i.c.c.j.Ch01EmptyCheckTest.emptyCheck    avgt        5  1.026        0.102  ns/op
 */
@BenchmarkMode(Mode.AverageTime)
@Warmup(iterations = 5)
@Measurement(iterations = 5, time = 1, timeUnit = TimeUnit.SECONDS)
@Threads(1)
@Fork(1)
@OutputTimeUnit(TimeUnit.NANOSECONDS)
@State(Scope.Benchmark)
public class Ch01EmptyCheckTest {

	List<Org> orgs = AreaBuilder.initSize(0);
	public static void main(String[] args) throws RunnerException {
		Options opt = new OptionsBuilder().include(Ch01EmptyCheckTest.class.getSimpleName()).forks(1).build();
		new Runner(opt).run();

	}

	@Benchmark
	public void emptyCheck() {
		AreaBuilder.buildProvinceWithEmptyCheck(orgs);
	}

	@Benchmark
	public void common() {
		AreaBuilder.buildProvince(orgs);
	}
	
	static List<Org> empty(){
		return new ArrayList<>();
	}

}
