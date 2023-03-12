/**
 * 从Java 8开始，java.time包提供了新的日期和时间API，主要涉及的类型有：

    本地日期和时间：LocalDateTime，LocalDate，LocalTime；
    带时区的日期和时间：ZonedDateTime；
    时刻：Instant；
    时区：ZoneId，ZoneOffset；
    时间间隔：Duration。
    以及一套新的用于取代SimpleDateFormat的格式化类型DateTimeFormatter。


    此外，新API修正了旧API不合理的常量设计：
    Month的范围用1~12表示1月到12月；
    Week的范围用1~7表示周一到周日。

    最后，新API的类型几乎全部是不变类型（和String类似），可以放心使用不必担心被修改。
 */
public class LocalDateTime用法 {

    base() {
        //now()获取到的总是以当前默认时区返回的，
        LocalDate d = LocalDate.now(); // 当前日期
        LocalTime t = LocalTime.now(); // 当前时间
        LocalDateTime dt = LocalDateTime.now(); // 当前日期和时间
        System.out.println(d); // 严格按照ISO 8601格式打印
        System.out.println(t); // 严格按照ISO 8601格式打印
        System.out.println(dt); // 严格按照ISO 8601格式打印
    }

    /**
     * 上述代码其实有一个小问题，在获取3个类型的时候，由于执行一行代码总会消耗一点时间，因此，3个类型的日期和时间很可能对不上（时间的毫秒数基本上不同）。为了保证获取到同一时刻的日期和时间，可以改写如下：
     */
    base2() {
        LocalDateTime dt = LocalDateTime.now(); // 当前日期和时间
        LocalDate d = dt.toLocalDate(); // 转换到当前日期
        LocalTime t = dt.toLocalTime(); // 转换到当前时间
    }

    /**
     * 通过指定的日期和时间创建LocalDateTime可以通过of()方法
     */
    base3() {
        LocalDate d2 = LocalDate.of(2019, 11, 30); // 2019-11-30, 注意11=11月
        LocalTime t2 = LocalTime.of(15, 16, 17); // 15:16:17
        LocalDateTime dt2 = LocalDateTime.of(2019, 11, 30, 15, 16, 17);
        LocalDateTime dt3 = LocalDateTime.of(d2, t2);        
    }

    

    base4() {
         // 自定义格式化:
         DateTimeFormatter dtf = DateTimeFormatter.ofPattern("yyyy/MM/dd HH:mm:ss");
         System.out.println(dtf.format(LocalDateTime.now()));
 
         // 用自定义格式解析:
         LocalDateTime dt2 = LocalDateTime.parse("2019/11/30 15:16:17", dtf);
         System.out.println(dt2);
    }


    /**
     * 
     * 对日期和时间进行加减的非常简单的链式调用：
     */
    base5() {
        LocalDateTime dt = LocalDateTime.of(2019, 10, 26, 20, 30, 59);
        System.out.println(dt);
        // 加5天减3小时:
        LocalDateTime dt2 = dt.plusDays(5).minusHours(3);
        System.out.println(dt2); // 2019-10-31T17:30:59
        // 减1月:
        LocalDateTime dt3 = dt2.minusMonths(1);
        System.out.println(dt3); // 2019-09-30T17:30:59
    }

    /**
     * 对日期和时间进行调整则使用withXxx()方法
     * 注意到调整月份时，会相应地调整日期，即把2019-10-31的月份调整为9时，日期也自动变为30。
     */
    base6() {
        LocalDateTime dt = LocalDateTime.of(2019, 10, 26, 20, 30, 59);
        System.out.println(dt);
        // 日期变为31日:
        LocalDateTime dt2 = dt.withDayOfMonth(31);
        System.out.println(dt2); // 2019-10-31T20:30:59
        // 月份变为9:
        LocalDateTime dt3 = dt2.withMonth(9);
        System.out.println(dt3); // 2019-09-30T20:30:59



        // 本月第一天0:00时刻:
        LocalDateTime firstDay = LocalDate.now().withDayOfMonth(1).atStartOfDay();
        System.out.println(firstDay);

        // 本月最后1天:
        LocalDate lastDay = LocalDate.now().with(TemporalAdjusters.lastDayOfMonth());
        System.out.println(lastDay);

        // 下月第1天:
        LocalDate nextMonthFirstDay = LocalDate.now().with(TemporalAdjusters.firstDayOfNextMonth());
        System.out.println(nextMonthFirstDay);

        // 本月第1个周一:
        LocalDate firstWeekday = LocalDate.now().with(TemporalAdjusters.firstInMonth(DayOfWeek.MONDAY));
        System.out.println(firstWeekday);
    }


    /**
     * 判断两个LocalDateTime的先后，可以使用isBefore()、isAfter()方法，对于LocalDate和LocalTime类似
     */
    base7() {
        LocalDateTime now = LocalDateTime.now();
        LocalDateTime target = LocalDateTime.of(2019, 11, 19, 8, 15, 0);
        System.out.println(now.isBefore(target));
        System.out.println(LocalDate.now().isBefore(LocalDate.of(2019, 11, 19)));
        System.out.println(LocalTime.now().isAfter(LocalTime.parse("08:15:00")));
    }


    
}
