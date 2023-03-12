import TileLayer from 'ol/layer/Tile'
import TileGrid from 'ol/tilegrid/TileGrid'
import XYZ from 'ol/source/XYZ'
import { TileWMS } from 'ol/source'
// import { get as getProjection, getTransform } from 'ol/proj'
// import { getWidth, getTopLeft, applyTransform } from 'ol/extent'

// import WMTSOGC from '../../base/ogc/WMTS'
// import TileClip from '../util/TileClip'

/**
 * cxgis.ol.factory.LyrTileFactory
 * 瓦片图层处理工厂
 * @class
 * @memberof cxgis.ol.factory
 * @author zhang qing
 * @since  2020-07-01
 */
class LyrTileFactory {
  /**
   * 根据maptype创建瓦片图层
   * @static
   * @method
   * @param {Object}                options                        图层的配置参数
   * @param {String}                options.url                    tile服务url
   * @param {String}                options.layer                  tile从服务端请求的发布图层名，如"vopp:cqGroup"
   * @param {Number}                options.maptype                图层的类型，0-wms,1-天地图，2-百度地图，3-谷歌地图，4-wmts，5-ArcRest，6-高德
   * @param {String}                options.version                tile请求版本
   * @param {String}                options.format                 tile请求数据返回格式
   * @param {Number}                options.major                  图层主类型
   * @param {String}                [options.name]                 图层名
   * @param {Boolean}               options.visible                图层初始是否可见
   * @param {Number}                options.opacity                不透明度，1为不透明
   * @param {String}                options.mapSrsUnit             地图的投影单位，如：m,degree等
   * @param {ol.Extent}             [options.mapExtent]            地图的实体范围(百度地图的tilegrid.extent没值是就必须配置)
   * @param {Object}                [options.params]               附加参数对象
   * @param {String}                [options.crs]                  投影的编码代号，如 EPSG:4326
   * @param {Number}                [options.cacheSize=2048]       缓存瓦片的大小，undefined时缓存2048
   * @param {Object}                [options.tilegrid]             瓦片网格策略对象（属性与ol中要求一致）.如果没有，就用地图投影对应的切片方案
   * @param {Array.<Number>}        options.tilegrid.resolutions   网格的各级分辨率
   * @returns {esri/layers/TileLayer}                              瓦片source对象
   */
  static createByType(options = {}) {
    let lyr = null
    options = options || {}
    switch (options.maptype) {
      // 标准ogc服务
      case 'wms':
        lyr = this.createWMS(options)
        break
      case 'xyz':
        lyr = this.createXYZ(options)
        break
    }

    return lyr
  }

  /**
   * 创建标准wms图层，具体配置信息依赖服务能力说明
   * @static
   * @method
   * @param {Object}                options                        图层的配置参数
   * @param {String}                options.url                    wms的服务url
   *    arcgis的wms：http://192.168.39.239:6080/arcgis/services/swdz/xzq1/MapServer/WMSServer
   *           wms能力说明：http://192.168.39.239:6080/arcgis/services/swdz/xzq1/MapServer/WMSServer?request=GetCapabilities&service=WMS
   *
   * @param {String}                options.layer                  tile从服务端请求的发布图层名，如"vopp:cqGroup"
   * @param {Object}                [options.params]               标注wms的其他参数对象，如SLD_BODY自定义样式，STYLES要和SLD_BODY中的名称一致
   * @param {String}                [options.name]                 图层名
   * @param {String}                [options.version='1.3.0']      tile请求版本
   * @param {String}                [options.format='image/png']   tile请求数据返回格式
   * @param {String}                [options.crs]                  投影的编码代号，如 EPSG:4326
   * @param {Number}                [options.cacheSize]            缓存瓦片的大小
   * @param {Object}                [options.tilegrid]             瓦片网格策略对象（属性与ol中要求一致）.如果没有，就用地图投影对应的切片方案
   * @param {Array.<Number>}        options.tilegrid.resolutions   网格的各级分辨率
   * @returns {ol.layer.Tile}                                      瓦片source对象
   */
  static createWMS(options) {
    var tilegrid
    if (window.cx.base.isObj(options.tilegrid)) {
      tilegrid = new TileGrid(options.tilegrid)
    }

    const source = new TileWMS({
      url: options.url,
      crossOrigin: '*',
      projection: options.crs,
      tileGrid: tilegrid,
      cacheSize: options.cacheSize,
      wrapX: false,
      params: Object.assign(
        {
          VERSION: options.version || '1.3.0',
          LAYERS: options.layer,
          FORMAT: options.format || 'image/png'
        },
        options.params
      )
    })

    return new TileLayer({
      source: source,
      name: options.name
    })
  }

  /**
   * 使用XYZ的方式加载瓦片地图，（难度上WMTS更繁琐，效率上XYZ更优）
   * @static
   * @method
   * @param {Object}                options                         图层的配置参数
   * @param {String}                options.url                     tile服务url
   * @param {String}                [options.name]                  图层名
   * @param {String}                [options.crs]                   投影的编码代号，如 EPSG:4326
   * @param {Number}                [options.cacheSize]             缓存瓦片的大小，根据屏幕大小确定，太小的话会自动扩大
   * @param {Object}                [options.tileGrid]              瓦片网格策略对象（属性与ol中要求一致）.如果没有，就用地图投影对应的切片方案
   * @param {ol.coordinate}         [options.tileGrid.origin]       网格的裁切原点
   * @param {ol.extent}             [options.tileGrid.extent]       瓦片的裁剪范围，配置后不请求该范围外的瓦片。没有配置origin默认用左上角为原点
   * @param {Array.<Number>}        [options.tileGrid.resolutions]  网格的各级分辨率
   * @param {Object}                [options.tileLoadFunction]      瓦片加载处理函数
   * @param {String}                [options.tileUrlFunction]       瓦片url处理函数
   * @param {Object}                [options.clipOpts]              瓦片裁剪配置对象，设置后会内置tileUrlFunction、tileLoadFunction函数，优先级低于传入的函数
   * @param {Boolean}               [options.clipOpts.isClip=false] 是否根据范围精准裁剪（有clipOpts.geom时才生效）。值为true时利用tileLoadFunction精准裁剪，否则只用tileUrlFunction过滤到没有相交的瓦片
   * @param {ol.geom.Polygon}       [options.clipOpts.geom=geom]    瓦片裁剪范围，设置后用内置tileUrlFunction按范围过滤瓦片
   * @returns {ol.layer.Tile}                                       瓦片layer对象
   * @example
   */
  static createXYZ(options = {}) {
    var tilegrid
    // const rest = 0.703125
    // const tyem = []
    // for (let i = 0; i < 21; i++) { tyem.push(rest / Math.pow(2, i)) }
    // options.tilegrid = {
    //   origin: [-180, 90],
    //   resolutions: tyem
    // }
    if (window.cx.base.isObj(options.tilegrid)) {
      tilegrid = new TileGrid(options.tilegrid)
    }

    const source = new XYZ({
      url: options.url,
      crossOrigin: '*',
      projection: options.crs,
      tileGrid: tilegrid,
      wrapX: false,
      cacheSize: options.cacheSize,
      tileLoadFunction: options.tileLoadFunction,
      tileUrlFunction: options.tileUrlFunction
    })

    return new TileLayer({
      source: source,
      name: options.name
    })
  }
}

export default LyrTileFactory
