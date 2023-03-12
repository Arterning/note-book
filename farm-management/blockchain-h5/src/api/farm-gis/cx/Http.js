import axios from 'axios'

class Http {
  /**
   * GET方式调用
   * @method
   * @static
   * @param   {String}    url         URL
   * @param   {Object}    params      参数
   * @param   {Object}    [urlConf]   请求的配置数据
   * @returns {Promise}
   */
  static get(url, params, urlConf) {
    return new Promise((resolve, reject) => {
      axios
        .get(url, urlConf)
        .then(ret => {
          resolve(ret.data)
        })
        .catch(reject)
    })
  }
}

export default Http
