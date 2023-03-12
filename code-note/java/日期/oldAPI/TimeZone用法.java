public class TimeZone用法 {
    
    public static void main(String[] args) {
        TimeZone tzDefault = TimeZone.getDefault(); // 当前时区
        TimeZone tzGMT9 = TimeZone.getTimeZone("GMT+09:00"); // GMT+9:00时区
        TimeZone tzNY = TimeZone.getTimeZone("America/New_York"); // 纽约时区
        System.out.println(tzDefault.getID()); // Asia/Shanghai
        System.out.println(tzGMT9.getID()); // GMT+09:00
        System.out.println(tzNY.getID()); // America/New_York
    }


    /*
     * 利用Calendar进行时区转换的步骤是：

        清除所有字段；
        设定指定时区；
        设定日期和时间；
        创建SimpleDateFormat并设定目标时区；
        格式化获取的Date对象（注意Date对象无时区信息，时区信息存储在SimpleDateFormat中）。
     */
    public void SetCalendarWithTimeZone() {
        // 当前时间:
        Calendar c = Calendar.getInstance();
        // 清除所有:
        c.clear();
        // 设置为北京时区:
        c.setTimeZone(TimeZone.getTimeZone("Asia/Shanghai"));
        // 设置年月日时分秒:
        c.set(2019, 10 /* 11月 */, 20, 8, 15, 0);
        
        // 显示时间:
        var sdf = new SimpleDateFormat("yyyy-MM-dd HH:mm:ss");
        //注意哦 时区是设置在SimpleDateFormat上的哦
        sdf.setTimeZone(TimeZone.getTimeZone("America/New_York"));
        System.out.println(sdf.format(c.getTime()));
        // 2019-11-19 19:15:00
    }


    /**
     * Calendar对时间进行加减
     */
    public void addOrSub() {
         // 当前时间:
         Calendar c = Calendar.getInstance();
         // 清除所有:
         c.clear();
         // 设置年月日时分秒:
         c.set(2019, 10 /* 11月 */, 20, 8, 15, 0);
         // 加5天并减去2小时:
         c.add(Calendar.DAY_OF_MONTH, 5);
         c.add(Calendar.HOUR_OF_DAY, -2);
         // 显示时间:
         var sdf = new SimpleDateFormat("yyyy-MM-dd HH:mm:ss");
         Date d = c.getTime();
         System.out.println(sdf.format(d));
         // 2019-11-25 6:15:00
    }

}
