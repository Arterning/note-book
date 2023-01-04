import Feature from 'ol/Feature'
import Point from 'ol/geom/Point'
import Polygon from 'ol/geom/Polygon'
import MultiPolygon from 'ol/geom/MultiPolygon'
import LineString from 'ol/geom/LineString'

/**
 * cxgis.ol.factory.FeatureFactory
 * feature处理工厂
 * @class
 * @memberof cxgis.ol.factory
 * @author zhang qing
 * @since  2020-07-01
 */
class FeatureFactory {
  /**
   * 创建feature
   * @static
   * @method
   * @param {Object}         gra       图形对象
   * @param {Array.<Number>} gra.coor  图形的坐标序列
   * @param {Number}         geoType   图形类型，entity中type类型值：2-点，3-线，4-面
   * @returns {ol/Feature}             创建好的feature对象
   */
  static createFea(gra, geoType) {
    let fea = null
    switch (geoType) {
      case window.cx.consts.ENTITY_PNT:
        fea = this.createPoint(gra)
        break
      case window.cx.consts.ENTITY_LIN:
        fea = this.createLineString(gra)
        break
      case window.cx.consts.ENTITY_REG:
        fea =
          gra.coor.length > 1
            ? this.createMultiPolygon(gra)
            : this.createPolygon(gra)
        break
    }

    return fea
  }

  /**
   * 创建feature
   * @static
   * @method
   * @param {Object}         gra       图形对象
   * @param {Array.<Number>} gra.coor  图形的坐标序列
   * @param {String}         type      图形类型，如 Point、LineString、Polygon
   * @returns {ol/Feature}             创建好的feature对象
   */
  static createByType(gra, type) {
    let fea = null
    switch (type) {
      case 'Point':
        fea = this.createPoint(gra)
        break
      case 'LineString':
        fea = this.createLineString(gra)
        break
      case 'Polygon':
        fea = this.createPolygon(gra)
        break
      case 'MultiPolygon':
        fea = this.createMultiPolygon(gra)
        break
    }

    return fea
  }

  /**
   * 创建点feature
   * @static
   * @method
   * @param {Object}         gra       图形对象
   * @param {Array.<Number>} gra.coor  图形的坐标序列
   * @returns {ol/Feature}             创建好的feature对象
   */
  static createPoint(gra) {
    const fea = new Feature({
      geometry: new Point(gra.coor)
    })
    fea.setProperties(gra)

    return fea
  }

  /**
   * 创建面feature
   * @static
   * @method
   * @param {Object}         gra       图形对象
   * @param {Array.<Number>} gra.coor  图形的坐标序列
   * @returns {ol/Feature}             创建好的feature对象
   */
  static createMultiPolygon(gra) {
    const fea = new Feature({
      geometry: new MultiPolygon(gra.coor)
    })
    fea.setProperties(gra)

    return fea
  }

  /**
   * 创建面feature
   * @static
   * @method
   * @param {Object}         gra       图形对象
   * @param {Array.<Number>} gra.coor  图形的坐标序列
   * @returns {ol/Feature}             创建好的feature对象
   */
  static createPolygon(gra) {
    const fea = new Feature({
      geometry: new Polygon(gra.coor)
    })
    fea.setProperties(gra)

    return fea
  }

  static createPolygonGeometry(coor) {
    return new Polygon(coor)
  }

  /**
   * 创建线feature
   * @static
   * @method
   * @param {Object}         gra       图形对象
   * @param {Array.<Number>} gra.coor  图形的坐标序列
   * @returns {ol/Feature}             创建好的feature对象
   */
  static createLineString(gra) {
    const fea = new Feature({
      geometry: new LineString(gra.coor)
    })
    fea.setProperties(gra)

    return fea
  }
}

export default FeatureFactory
