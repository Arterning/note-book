十进制整数在转化成二进制数时不会有精度问题，
那么把十进制小数扩大N倍让它在整数的维度上进行计算，并保留相应的精度信息。

BigDecimal 由任意精度的整数非标度值 和32 位的整数标度 (scale) 组成。如果为零或正数，则标度是小数点后的位数。如果为负数，则将该数的非标度值乘以 10 的负scale 次幂。因此，BigDecimal表示的数值是(unscaledValue × 10-scale)

scale就是表示小数点后有多少位。

构造函数（主要测试参数类型为double和String的两个常用构造函数）
       BigDecimal aDouble =new BigDecimal(1.22);

        System.out.println("construct with a double value: " + aDouble);

        BigDecimal aString = new BigDecimal("1.22");

         System.out.println("construct with a String value: " + aString);

        你认为输出结果会是什么呢？如果你没有认为第一个会输出1.22，那么恭喜你答对了，输出结果如下：

         construct with a doublevalue:1.2199999999999999733546474089962430298328399658203125

         construct with a String value: 1.22

        JDK的描述：1、参数类型为double的构造方法的结果有一定的不可预知性。有人可能认为在Java中写入newBigDecimal(0.1)所创建的BigDecimal正好等于 0.1（非标度值 1，其标度为 1），但是它实际上等于0.1000000000000000055511151231257827021181583404541015625。这是因为0.1无法准确地表示为 double（或者说对于该情况，不能表示为任何有限长度的二进制小数）。这样，传入到构造方法的值不会正好等于 0.1（虽然表面上等于该值）。

        2、另一方面，String 构造方法是完全可预知的：写入 newBigDecimal("0.1") 将创建一个 BigDecimal，它正好等于预期的 0.1。因此，比较而言，通常建议优先使用String构造方法。

        3、当double必须用作BigDecimal的源时，请注意，此构造方法提供了一个准确转换；它不提供与以下操作相同的结果：先使用Double.toString(double)方法，然后使用BigDecimal(String)构造方法，将double转换为String。要获取该结果，请使用static valueOf(double)方法。


public class BigDecimal {
    //值的绝对long型表示
    private final transient long intCompact;
    //值的小数点后的位数
    private final int scale;

    private final BigInteger intVal;
    //值的有效位数，不包含正负符号
    private transient int precision;
    private transient String stringCache;

    //加、减、乘、除、绝对值
    public BigDecimal add(BigDecimal augend) {}
    public BigDecimal subtract(BigDecimal subtrahend) {}
    public BigDecimal multiply(BigDecimal multiplicand) {}
    public BigDecimal divide(BigDecimal divisor) {}
    public BigDecimal abs() {}
}

比方说
如果创建对象
BigDecimal BigDecimal = new BigDecimal("123124.234234")

那么intCompact就是123124234234
scale就是小数点后的位数， 也就是6
precision就是整数位和小数位一起有多少位， 也就是12


加减乘除的实现

加法：long类型 +

减法：转成加法，加负数

乘法:  long类型 *, 多些进位超界判断

除法:  long类型 /, 多些小数位数保留判断

private static long add(long xs, long ys){
　　long sum = xs + ys;
　　if ( (((sum ^ xs) & (sum ^ ys))) >= 0L) { // not overflowed
　　　　return sum;
　　}
　　return INFLATED;
}



/**
 * If the absolute value of the significand of this BigDecimal is
 * less than or equal to {@code Long.MAX_VALUE}, the value can be
 * compactly stored in this field and used in computations.
 */
private final transient long intCompact;

分析上面的注释， 可以知道， 如果Bigdecimal传入的绝对值(包括小数位)
如果小鱼Long.MAX_VALUE， 那么unscaledValue就可以使用long来表示
否则的化， 如果很大， 那么unscaledValue就只能使用BigInteger了

/**
 * The unscaled value of this BigDecimal, as returned by {@link
 * #unscaledValue}.
 *
 * @serial
 * @see #unscaledValue
 */
private final BigInteger intVal;



BigInteger存储大数的方式就是将数字存储在一个整型的数组中
只用一个整型数组的话，如何表示一个整数的正负呢？
那么就需要有一个单独的成员变量来标明该数的正负。


BigDecimal和BigInteger都是不可变对象， 所以也是线程安全对象
类似String，在操作之后会产生一个新对象
