base = () => {
  var now = new Date();
  now; // Wed Jun 24 2015 19:49:22 GMT+0800 (CST)
  now.getFullYear(); // 2015, 年份
  now.getMonth(); // 5, 月份，注意月份范围是0~11，5表示六月
  now.getDate(); // 24, 表示24号
  now.getDay(); // 3, 表示星期三
  now.getHours(); // 19, 24小时制
  now.getMinutes(); // 49, 分钟
  now.getSeconds(); // 22, 秒
  now.getMilliseconds(); // 875, 毫秒数
  now.getTime(); // 1435146562875, 以number形式表示的时间戳
};

/**
 * 如果要创建一个指定日期和时间的Date对象，可以用：
 * 你可能观察到了一个非常非常坑爹的地方，就是JavaScript的月份范围用整数表示是0~11，0表示一月，1表示二月……，所以要表示6月，我们传入的是5！这绝对是JavaScript的设计者当时脑抽了一下，但是现在要修复已经不可能了。
 */
function base2() {
  var d = new Date(2015, 5, 19, 20, 15, 30, 123);
  d; // Fri Jun 19 2015 20:15:30 GMT+0800 (CST)
}

/**
 * 第二种创建一个指定日期和时间的方法是解析一个符合ISO 8601格式的字符串：
 */

function base3() {
  var d = Date.parse("2015-06-24T19:49:22.875+08:00");
  d; // 1435146562875
  /**
   * 但它返回的不是Date对象，而是一个时间戳。不过有时间戳就可以很容易地把它转换为一个Date：
   */
  var dd = new Date(1435146562875);
  dd; // Wed Jun 24 2015 19:49:22 GMT+0800 (CST)
  d.getMonth(); // 5
}
