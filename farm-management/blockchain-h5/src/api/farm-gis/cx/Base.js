/**
 * cx.base
 * 类-Base
 *
 * @class
 * @memberof    cx
 * @author      Zhang Fayong
 * @since       2018-04-28
 */
class Base {
  /**
   * @constructor
   */
  constructor() {
    this._newId = 0
    this._nextZIndex = 1000
    this.debug = false
  }

  /**
   * 比较v1和v2两个版本的大小
   * @param {string|number} v1 第一个版本值，如：'1.1.2' 或数值
   * @param {string|number} v2 第二个版本值
   * @returns {number}         判断的值，0是相等，1是v1大于v2，-1是v1小于v2
   */
  static compareVersions(v1, v2) {
    var s1 = ('' + v1).split('.')
    var s2 = ('' + v2).split('.')
    for (var i = 0; i < Math.max(s1.length, s2.length); i++) {
      var n1 = parseInt(s1[i] || '0', 10)
      var n2 = parseInt(s2[i] || '0', 10)

      if (n1 > n2) {
        return 1
      }

      if (n2 > n1) {
        return -1
      }
    }

    return 0
  }

  /**
   * 产生一个新ID
   * @static
   * @method
   * @returns {Number}
   */
  static nextId() {
    return this._newId++
  }

  /**
   * 产生一个新 z-index
   * @static
   * @method
   * @returns {Number}
   */
  static nextZIndex() {
    this._nextZIndex += 2 // 对话框，对话框下面的模态对话
    return this._nextZIndex
  }

  /**
   * 判断是否Boolean类型，是时返回true，否则false
   * @static
   * @method
   * @param   {Object}    val
   * @returns {Boolean}
   */
  static isBoolean(val) {
    return typeof val === 'boolean'
  }

  /**
   * 判断是否定义，不等于undefined时返回true，否则false
   * @static
   * @method
   * @param   {Object}    val
   * @returns {Boolean}
   */
  static isDef(val) {
    return val !== undefined
  }

  /**
   * 判断val是否为空，等于null时返回true，否则false
   * @static
   * @method
   * @param   {Object}    val
   * @returns {Boolean}
   */
  static isNull(val) {
    return val === null
  }

  /**
   * 判断val已定义且不为null，即有赋值时返回true，否则false
   * @static
   * @method
   * @param   {Object}    val
   * @returns {Boolean}
   */
  static isNNull(val) {
    return val !== undefined && val !== null
  }

  /**
   * 判断val是否为数值，是返回true，不是返回false
   * @static
   * @method
   * @param   {Object}    val
   * @returns {Boolean}
   */
  static isNum(val) {
    return typeof val === 'number' && !isNaN(val)
  }

  /**
   * 判断val是否为字符串，是返回true，不是返回false
   * @static
   * @method
   * @param   {Object}    val
   * @returns {Boolean}
   */
  static isStr(val) {
    return typeof val === 'string' || val instanceof String
  }

  /**
   * 判断val是否为数组，是返回true，不是返回false
   * @static
   * @method
   * @param   {Object}    val
   * @returns {Boolean}
   */
  static isArr(val) {
    return val instanceof Array
  }

  /**
   * 判断val是否为对象，是返回true，不是返回false
   * @static
   * @method
   * @param   {Object}    val
   * @returns {Boolean}
   */
  static isObj(val) {
    return val && (typeof val === 'object' || typeof val === 'function')
  }

  /**
   * 判断val是否为函数，是返回true不是返回false
   * @static
   * @method
   * @param   {Object}    val
   * @returns {Boolean}
   */
  static isFunc(val) {
    return typeof val === 'function'
  }

  /**
   * 如果给出的值为未定义，则给出缺省值
   * @static
   * @method
   * @param   {Object}  value            值
   * @param   {Object}  defValue         缺省值
   * @returns {Object}
   */
  static defVal(value, defValue) {
    return value !== undefined ? value : defValue
  }

  /**
   * 判断val(对象或者属性)是否在obj(数组或者对象)中
   * @static
   * @method
   * @param   {Object|Array}  obj
   * @param   {Object}        val
   * @returns {Boolean}
   */
  static contains(obj, val) {
    let i
    if (this.isArr(obj)) {
      for (i = 0; i < obj.length; i++) {
        if (obj[i] === val) {
          return true
        }
      }
    } else if (this.isObj(obj)) {
      for (i in obj) {
        if (obj[i] === val) {
          return true
        }
      }
    }
    return false
  }

  /**
   * 目标对象(target)继承对象(args可以多个参数)
   * @static
   * @method
   * @param   {Object}    target          目标对象
   * @param   {Object}    argvs           可变参数
   * @returns {Object}                    target
   */
  static extend(target, ...argvs) {
    for (const argv of argvs) {
      if (typeof argv === 'object') {
        for (const key in argv) {
          target[key] = argv[key]
        }
      }
    }
    return target
  }

  /**
   * 字符串调用
   * @static
   * @method
   * @param  {String}   func            函数字符串,可以'()'结束，也可以没有'()' (如 'test' 或 'test()')
   * @param  {Array}    argvs           参数数组
   * @returns {*}
   */
  static execute(func, argvs) {
    if (func) {
      let parent = window
      let obj = window
      let strs = func
      if (strs.indexOf('()') >= 0) {
        // 去掉()
        strs = strs.slice(0, -2)
      }
      strs = strs.split('.') // 查找包名
      for (let i = 0; i < strs.length; i++) {
        parent = obj
        obj = obj[strs[i]]
      }
      if (obj) {
        // 执行函数
        return obj.apply(parent, argvs)
      }
    }
  }

  /**
   * 取字词表符串
   * @static
   * @method
   * @param   {String}    key           字符串key
   * @param   {*}         argvs         变参，要被替换的目标字符串
   * @returns {String}                  词表字符串
   */
  static loc(key, ...argvs) {
    let ret = window.vm.$t(key)
    if (argvs.length > 0) {
      for (let i = 0; i < argvs.length; i++) {
        ret = ret.replace('%' + (i + 1), '' + argvs[i])
      }
    }
    return ret
  }

  /**
   * 根据组件的配置项，获取其对应的虚拟节点
   * @method
   * @param   {Object}    component    组件的配置对象
   * @param   {Object}    option       组件接收的参数，如{props:{a:1},attrs:{},on:{}....}。具体参见Vue官方API中的【渲染函数 & JSX】
   * @returns {VNode}     虚拟节点对象
   */
  static getVNode(component, option) {
    return window.vm.$createElement(component, option)
  }

  /**
   * 显示信息
   * @static
   * @method
   * @param {String}    message       消息文本
   * @param {String}    [type]        类型(success/warning/info/error)
   */
  static showMsg(message, type) {
    if (!type) {
      window.vm.$message(message)
    } else {
      if (type === 'error') {
        window.vm.$message.error(message)
      } else {
        window.vm.$message({ type: type, message: message })
      }
    }
  }

  /**
   * 弹框显示自定义组件
   * @method
   * @param   {Object}    param                参数对象
   * @param   {Function}  param.render         render函数。优先级最高，会传入构造函数$createElement作为参数。示例：h => h('span', null, '这是一段内容 ')
   * @param   {Object}    param.component      组件的配置对象。优先级比render低。
   * @param   {Object}    [param.option]       组件接收的参数，如{props:{a:1},attrs:{},on:{}....}
   * @param   {String}    [param.option.key]   每次弹出层打开后，Vue 会对新老 VNode 节点进行比对，然后将根据比较结果进行最小单位地修改视图。这也许会造成弹出层内容区域的组件没有重新渲染。当这类问题出现时，解决方案是给 VNode 加上一个不相同的 key。
   * @param   {Object}    [param.param]        messageBox的原生属性对象
   * @param   {Object}    [param.title]       自定义标题
   */
  static showMsgBox({ component, option, param, render, title }) {
    window.vm.$msgbox({
      title: title || '消息',
      message: render
        ? render(window.vm.$createElement)
        : this.getVNode(component, { key: new Date().getTime(), ...option }),
      showConfirmButton: false,
      ...param
    })
  }

  /**
   * 输出调试信息
   * @static
   * @method
   * @param  {String}     msg
   */
  static log(msg) {
    if (this.debug) {
      console.log(msg)
    }
  }

  /**
   * 设置计时器，来监测条件满足时执行回调
   * @static
   * @method
   * @param   {Function}    isReady     是否满足条件的判断函数
   * @param   {Function}    callBack    满足条件后的回调函数
   * @param   {Number}      [time=600] 间隔时间(ms)。
   */
  static watch(isReady, callBack, time = 200) {
    const func = () => {
      if (isReady()) {
        if (myInterval) {
          window.clearInterval(myInterval)
        }
        callBack()
      }
    }
    // 设置定时任务，每1000ms执行一次watch方法。
    const myInterval = window.setInterval(func, time)
  }
}

Base._newId = 0
Base._nextZIndex = 1000
Base.debug = false
Base.processing = {
  _loads: [],
  _getDom(target) {
    if (target) {
      return Base.isStr(target)
        ? document.getElementById(target) || document.body
        : target
    } else {
      return document.body
    }
  },
  /**
   * 显示控件
   * @method
   * @param {Element|String|null}  [target]       正在加载显示的容器对象或id，默认是body
   */
  show(target) {
    const dom = this._getDom(target)
    this.loading = window.vm.$loading({
      lock: true,
      text: window.vm.$loc('cx.processing'),
      spinner: 'el-icon-loading',
      background: 'rgba(255, 255, 255, 0.5)',
      customClass: 'cx-processing',
      target: dom
    })

    this._loads.push({ dom: dom, load: this.loading })
  },
  /**
   * 关闭控件
   * @method
   * @param {Element|String|null}  [target]       正在加载显示的容器对象或id，默认是body
   */
  close(target) {
    const dom = this._getDom(target)
    let idx = this._loads.findIndex(item => item.dom === dom)
    let loadObj = this._loads[idx]

    if (!loadObj && this._loads.length === 1) {
      loadObj = this._loads[0]
      idx = 0
    }

    if (loadObj && loadObj.load) {
      loadObj.load.close()
      this._loads.splice(idx, 1)
    }
  }
}
export default Base
