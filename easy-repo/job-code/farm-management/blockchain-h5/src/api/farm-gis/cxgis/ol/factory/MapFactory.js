import Map from 'ol/Map'
import View from 'ol/View'
import { get as projGet } from 'ol/proj'
import InteractionMan from '../control/InteractionMan'

/**
 * cxgis.ol.factory.MapFactory
 * 地图处理工厂
 * @class
 * @memberof cxgis.ol.factory
 * @author zhang qing
 * @since  2020-07-01
 */
class MapFactory {
  /**
   * 创建feature
   * @static
   * @method
   * @param {Object}            opts                     地图配置对象
   * @param {String|Element}    opts.target              地图容器对象
   * @param {Array}             opts.layers              图层对象序列
   * @param {Object}            opts.viewData            视图配置对象
   * @param {String}            opts.srsSys              地图的投影编码，如4326
   * @returns {Object}                                   创建好的Map对象
   * @returns {ol/Map}          returns.map              地图对象
   * @returns {ol/View}         returns.view             视图对象
   */
  static create(opts) {
    const proj = projGet('EPSG:' + opts.srsSys)

    // 创建地图视图对象
    const view = new View({
      center: opts.viewData.center,
      projection: proj,
      multiWorld: false,
      extent: opts.viewData.extent || proj.getExtent(),
      zoom: opts.viewData.zoom,
      maxZoom: opts.viewData.zoome,
      minZoom: opts.viewData.zooms,
      //1.设置缩放级别为整数
      constrainResolution: true,
      //2.关闭无级缩放地图
      smoothResolutionConstraint: false
    })

    // 创建地图对象
    const map = new Map({
      layers: opts.layers || [],
      target: opts.target,
      controls: [],
      view: view
    })

    if (opts.viewData.initextent) {
      view.fit(opts.viewData.initextent)
    }

    // 为ol.map对象扩展interact管理对象
    map.getInteractionMan = (() => {
      const _interactionMan = new InteractionMan({ map: map })
      return () => _interactionMan
    })()

    return { map: map, view: view }
  }

  /**
   * 根据系统配置扩展投影
   * @method
   * @param {Object}    参考系配置对象，key是参考系名称，如：EPSG:4326，value是配置对象，包含def和extent
   *   如 {"EPSG:2345":{"def":"+proj=tmerc +lat_0=0 +lon_0=117 +k=1 +x_0=500000 +y_0=0 +a=6378140 +b=6356755.288157528 +units=m +no_defs","extent":[190416.42,2452360.40,708209.04,5714208.92]}}
   */
  static extendProj() {
    return
  }

  /**
   * 定位到要素范围
   * @static
   * @method
   * @param {ol.Map}                  map            地图对象
   * @param {Array.<ol.Feature>}      feas           待定位的feature序列
   * @param {Number}                  [padding=230]  定位图形与地图边界的像素距离
   */
  static locateView(map, feas, padding = 230) {
    if (feas) {
      const lyr = window.cxgis.ol.factory.LyrVectorFactory.createVector()
      lyr.getSource().addFeatures(feas)
      const extent = lyr.getSource().getExtent()

      if (window.cx.base.isNum(extent[0] + extent[1] + extent[2] + extent[3])) {
        map
          .getView()
          .fit(extent, { padding: [padding, padding, padding, padding] })
      }
    }
  }
}

export default MapFactory
