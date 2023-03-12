/**
 * cxgis.ol.factory.ExtentFactory
 * extent处理工厂
 * @class
 * @memberof cxgis.ol.factory
 * @author zhang qing
 * @since  2020-07-01
 */
class ExtentFactory {
  /**
   * 根据点坐标和容差距离创建extent
   * @static
   * @method
   * @param {Object}             opts
   * @param {Coordinates}        opts.coor             点坐标
   * @param {Number}             opts.resolution       地图分辨率
   * @param {Number}             [opts.tolerance=3]    屏幕容差范围,单位 px
   * @returns {ol/Extent}                              创建extent，格式[minx, miny, maxx, maxy]
   */
  static createByPntCoor(opts) {
    opts.tolerance = opts.tolerance || 3
    const delta = opts.tolerance * opts.resolution
    return [
      opts.coor[0] - delta,
      opts.coor[1] - delta,
      opts.coor[0] + delta,
      opts.coor[1] + delta
    ]
  }

  /**
   * 根据extent生成多边形的坐标序列
   * @static
   * @method
   * @param {ol/Extent}                  extent         extent值
   * @returns {Array.<Coordinates>}                     构造好的多边形坐标序列
   */
  static getCoorPol(extent) {
    return [
      [
        [extent[0], extent[1]],
        [extent[2], extent[1]],
        [extent[2], extent[3]],
        [extent[0], extent[3]],
        [extent[0], extent[1]]
      ]
    ]
  }
}

export default ExtentFactory
