import VectorLayer from 'ol/layer/Vector'
import VectorSource from 'ol/source/Vector'

/**
 * cxgis.ol.factory.LyrVectorFactory
 * 矢量图层处理工厂
 * @class
 * @memberof cxgis.ol.factory
 * @author zhang qing
 * @since  2020-07-01
 */
class LyrVectorFactory {
  /**
   * 创建vector图层
   * @static
   * @method
   * @param {Object}                   opts                             参数对象，,是ol/layer/vector构造函数的参数，不需要传source
   * @param {Function}                 [opts.style]                     图形类型
   * @param {Boolean}                  [opts.declutter=false]           是否避免图形压盖
   * @returns {ol/layer/vector}                                         图层对象
   */
  static createVector(opts = {}) {
    opts.source = new VectorSource()
    return new VectorLayer(opts)
  }
}

export default LyrVectorFactory
