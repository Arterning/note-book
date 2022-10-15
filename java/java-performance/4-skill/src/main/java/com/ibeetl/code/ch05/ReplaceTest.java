package com.ibeetl.code.ch05;

import org.openjdk.jmh.annotations.*;
import org.openjdk.jmh.runner.Runner;
import org.openjdk.jmh.runner.RunnerException;
import org.openjdk.jmh.runner.options.Options;
import org.openjdk.jmh.runner.options.OptionsBuilder;

import java.util.concurrent.TimeUnit;
/**
 * 一个演示使用local变量带来的性能略微提升
 * @author 公众号 java系统优化
 */
@BenchmarkMode(Mode.Throughput)
@Warmup(iterations = 5)
@Measurement(iterations = 5, time = 1, timeUnit = TimeUnit.SECONDS)
@Threads(1)
@Fork(1)
@OutputTimeUnit(TimeUnit.MILLISECONDS)
@State(Scope.Benchmark)
public class ReplaceTest {

	String str = "hello,this is effective java code";

    @Benchmark
    public char[] replaceByGetField() {
		return new ReplaceByGetField(str.toCharArray()).replace();

    }


    @Benchmark
    public char[] replaceByLocal() {
		return new ReplaceByLocal(str.toCharArray()).replace();

    }

	/**
	 * ReplaceByLocal和ReplaceByGetField 代码几乎一样，不同地方在于ReplaceByLocal使用栈变量
	 */
	static class ReplaceByLocal{
		char[] value = new char[0];
		char oldChar = 'e';
		char newChar ='t';
		public ReplaceByLocal(char[] value){
			this.value = value;
		}

		public char[] replace(){
			int len = value.length;
			int i = -1;
			char[] val = value; /* avoid getfield opcode */

			while (++i < len) {
				if (val[i] == oldChar) {
					break;
				}
			}
			if (i < len) {
				char buf[] = new char[len];
				for (int j = 0; j < i; j++) {
					buf[j] = val[j];
				}
				while (i < len) {
					char c = val[i];
					buf[i] = (c == oldChar) ? newChar : c;
					i++;
				}
				return buf;
			}

			return new char[0];
		}

	}


    static class ReplaceByGetField{
		char[] value = new char[0];
		char oldChar = 'e';
		char newChar ='t';
		public ReplaceByGetField(char[] value){
			this.value = value;
		}


		public char[] replace(){
			int len = value.length;
			int i = -1;

			while (++i < len) {
				if (value[i] == oldChar) {
					break;
				}
			}
			if (i < len) {
				char buf[] = new char[len];
				for (int j = 0; j < i; j++) {
					buf[j] = value[j];
				}
				while (i < len) {
					char c = value[i];
					buf[i] = (c == oldChar) ? newChar : c;
					i++;
				}
				return buf;
			}

			return new char[0];
		}

	}



    public static void main(String[] args) throws RunnerException {
        Options opt = new OptionsBuilder()
                .include(ReplaceTest.class.getSimpleName())
                .forks(1)
                .build();
        new Runner(opt).run();
    }
}
