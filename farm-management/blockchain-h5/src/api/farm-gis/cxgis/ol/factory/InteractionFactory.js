import Select from 'ol/interaction/Select'
import Draw from 'ol/interaction/Draw'

/**
 * cxgis.ol.factory.InteractionFactory
 * ol的interaction交互处理工厂
 * @class
 * @memberof cxgis.ol.factory
 * @author zhang qing
 * @since  2020-07-01
 */
class InteractionFactory {
  /**
   * 创建图层选择交互工具
   * @static
   * @method
   * @param {Object}                           opts                             参数对象,是select构造函数的参数
   * @param {String}                           [opts.mutexType='mapMutexAct']   交互空间的互斥类型，有值时才在InteractionMan中进行同类互斥管理
   * @param {Array.<ol/layers/Vector>}         opts.layers                      选择操作的目标图层序列
   * @param {Function}                         opts.onSelect                    选中回调函数
   * @param {Object}                           opts.onSelect.evt                选中的事件对象
   * @param {Array.<ol/Feature>}               opts.onSelect.evt.selected       选中的feature序列
   * @param {ol/events/condition}              [opts.condition='singleClick']   执行选中的鼠标操作方式
   * @param {Function}                         [opts.style]                     样式函数，默认使用ol内置样式，设为false是不高亮
   * @param {Function}                         [opts.filter]                    过滤函数
   * @param {Function}                         [opts.hitTolerance=3]            点击的容差值，单位px
   * @returns {ol/interaction/Select}                                           控件对象
   */
  static createSelect(opts = {}) {
    const act = new Select(opts)
    act.set('mutexType', opts.mutexType || 'mapMutexAct')

    act.on('select', evt => {
      if (opts.onSelect) {
        opts.onSelect(evt)
      }
    })
    return act
  }

  /**
   * 创建多边形绘制交互工具
   * @static
   * @method
   * @param {Object}                           opts                             参数对象，是ol/interaction/Draw构造函数的参数
   * @param {String}                           [opts.mutexType='mapMutexAct']   交互空间的互斥类型，有值时才在InteractionMan中进行同类互斥管理
   * @param {ol.soure.Vectory}                 [opts.source]                    图形类型
   * @returns {ol/interaction/Draw}                                             图层对象
   */
  static createDrawPolygon(opts = {}) {
    opts.type = 'Polygon'
    const act = new Draw(opts)
    act.set('mutexType', opts.mutexType || 'mapMutexAct')

    act.on('drawstart', evt => {
      if (opts.onDrawStart) {
        opts.onDrawStart(evt)
      }
    })
    act.on('drawend', evt => {
      if (opts.onDrawEnd) {
        opts.onDrawEnd(evt)
      }
    })
    return act
  }
  /**
   * 创建点绘制交互工具
   * @static
   * @method
   * @param {Object}                           opts                             参数对象，是ol/interaction/Draw构造函数的参数
   * @param {String}                           [opts.mutexType='mapMutexAct']   交互空间的互斥类型，有值时才在InteractionMan中进行同类互斥管理
   * @param {ol.soure.Vectory}                 [opts.source]                    图形类型
   * @returns {ol/interaction/Draw}                                             图层对象
   */
  static createDrawPoint(opts = {}) {
    opts.type = 'Point'
    const act = new Draw(opts)
    act.set('mutexType', opts.mutexType || 'mapMutexAct')

    act.on('drawstart', evt => {
      if (opts.onDrawStart) {
        opts.onDrawStart(evt)
      }
    })
    act.on('drawend', evt => {
      if (opts.onDrawEnd) {
        opts.onDrawEnd(evt)
      }
    })
    return act
  }
  /**
   * 创建线绘制交互工具
   * @static
   * @method
   * @param {Object}                           opts                             参数对象，是ol/interaction/Draw构造函数的参数
   * @param {String}                           [opts.mutexType='mapMutexAct']   交互空间的互斥类型，有值时才在InteractionMan中进行同类互斥管理
   * @param {ol.soure.Vectory}                 [opts.source]                    图形类型
   * @returns {ol/interaction/Draw}                                             图层对象
   */
  static createDrawLineString(opts = {}) {
    opts.type = 'LineString'
    const act = new Draw(opts)
    act.set('mutexType', opts.mutexType || 'mapMutexAct')

    act.on('drawstart', evt => {
      if (opts.onDrawStart) {
        opts.onDrawStart(evt)
      }
    })
    act.on('drawend', evt => {
      if (opts.onDrawEnd) {
        opts.onDrawEnd(evt)
      }
    })
    return act
  }
  /**
   * 创建圆绘制交互工具
   * @static
   * @method
   * @param {Object}                           opts                             参数对象，是ol/interaction/Draw构造函数的参数
   * @param {String}                           [opts.mutexType='mapMutexAct']   交互空间的互斥类型，有值时才在InteractionMan中进行同类互斥管理
   * @param {ol.soure.Vectory}                 [opts.source]                    图形类型
   * @returns {ol/interaction/Draw}                                             图层对象
   */
  static createDrawCircle(opts = {}) {
    opts.type = 'Circle'
    const act = new Draw(opts)
    act.set('mutexType', opts.mutexType || 'mapMutexAct')

    act.on('drawstart', evt => {
      if (opts.onDrawStart) {
        opts.onDrawStart(evt)
      }
    })
    act.on('drawend', evt => {
      if (opts.onDrawEnd) {
        opts.onDrawEnd(evt)
      }
    })
    return act
  }
}

export default InteractionFactory
