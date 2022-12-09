
```scala
package spark

import org.apache.spark.{SparkConf, SparkContext}

object HelloWorld {

  def main(args: Array[String]): Unit = {
    val conf = new SparkConf().setMaster("local").setAppName("HelloWorld")

    val sc = new SparkContext(conf)

    val helloWorld = sc.parallelize(List("Hello,World!","Hello,Spark!","Hello,BigData!"))

    helloWorld.foreach(line => println(line))
  }

}
```


```java
import org.apache.spark.SparkConf;
import org.apache.spark.api.java.JavaRDD;
import org.apache.spark.api.java.JavaSparkContext;

import java.util.Arrays;

public class HelloWorldJava {

    public static void main(String[] args){

        SparkConf conf = new SparkConf().setMaster("local").setAppName("HelloWorldJava");

        JavaSparkContext sc = new JavaSparkContext(conf);

        JavaRDD<String> helloWorld = sc.parallelize(Arrays.asList("Hello,World","Hello,Spark","Hello,BigData"));

        System.out.println(helloWorld.collect());

    }

}

```