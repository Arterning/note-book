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
 */


