import base from './Base'
class Array {
  /**
   * 将数组转成map
   * @method
   * @param {Array.Object}  data  带转换的数据
   * @param {String}        key   转换成map所依据的key
   * @returns {Object.Array}  转换成的map对象
   */
  static listToMap(data, key) {
    var map = {}
    data.forEach(obj => {
      map[obj[key]] = map[obj[key]] || []
      map[obj[key]].push(obj)
    })
    return map
  }

  /**
   * 从数组中删除指定位置删除多个元素
   * @static
   * @method
   * @param   {Array}     arr             数组
   * @param   {Number}    idx             删除位置
   * @param   {Number}    number          个数
   * @returns {Object}                    this
   */
  static removeAt(arr, idx, number) {
    if (base.isArr(arr)) {
      arr.splice(idx, number)
    }
    return this
  }
}

export default Array
