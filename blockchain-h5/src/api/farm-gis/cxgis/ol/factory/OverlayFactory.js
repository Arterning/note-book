import Overlay from 'ol/Overlay'

/**
 * cxgis.ol.factory.OverlayFactory
 * 地图overlayer处理工厂
 * @class
 * @memberof cxgis.ol.factory
 * @author zhang qing
 * @since  2020-07-01
 */
class OverlayFactory {
  static createOverlay(opts) {
    return new Overlay(opts)
  }
}

export default OverlayFactory
