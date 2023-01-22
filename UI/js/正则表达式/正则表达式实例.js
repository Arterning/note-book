/**
 * 
 * 参考资料
 * https://www.bilibili.com/video/BV1QK4y1K72U/?spm_id_from=333.1007.top_right_bar_window_custom_collection.content.click&vd_source=5b90c6e8f3c6d4f4052d5cbb23dc556a
 */

/**
 * 电话号码
 */
var phone = /^1[34578]\d{9}$/g;


const qqNumber = /^[1-9][0-9]{4-9}$/g

/**
 * 16进制数字
 * |表示或者
 * ()使得两种情况合为一组 使得#应用于括号后面的内容
 * ? 使得前面的字符可有可无
 */
const SixTeenNumber = /#?([0-9a-zA-F]{6}|[0-9a-zA-F]{3})/g



/**
 * Email
 */
const eMail = /([a-zA-Z0-9_\-\.]+)@([a-zA-Z0-9_\-\.]+)\.([A-Za-z]{2,6})$/g



/**
 * URL
 */
const URL = /^((https?|ftp|file):\/\/)?([\da-z\.\-]+)\.([a-z\.]{2,6})([\/\w\.\-]*)*(\/?$)/g


/**
 * 匹配html标签
 * <div class="test">BoyNextDoor♂</div>
 * <img />
 * 
 * 怕你忘记了 *代表0到无穷大 ?表示可有可无 +代表一到无穷广大
 */
const tag = /^<([a-z]+)([^>]+)*(?:>.*<\/\1>|\s+\/>)$/gm

/**
 * ip地址 4组0-255
 * 255.255.255.255
 */
const ipAddr = /[]/g


/**
 * 身份证号码
 * 
 * 顶不住了 这正则表达式太难受了 他妈的和火星文有什么差别 
 * 直接编写这种东西很容易出错 希望以后有那种可以自动拿来用的正则表达式。
 * 我直接引用就可以了
 */
const idCard = [];