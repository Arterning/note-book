import X2js from 'x2js'
import * as olProj from 'ol/proj'
import { GeoJSON } from 'ol/format'
import OgcWFS from '../../base/ogc/WFS'
import VectorLayer from 'ol/layer/Vector'
import InteractionFactory from '../factory/InteractionFactory'

/**
 * cxgis.ol.util.SelFeaByWFS
 * 基于wfs服务选择feature的工具类
 * @class
 * @memberof cxgis.ol.util
 * @author zhang qing
 * @since  2020-07-01
 */
class SelFeaByWFS {
  /**
   * 根据mapSvrUrl指定的arcgis服务，获取点击位置的feature对象
   * @param {Object}            opts                         IdentifyTask需要的配置参数对象,更多参数见esri.tasks.support.IdentifyParameters
   * @param {ol/Map}            opts.map                     view对象，默认使用gMap.view
   * @param {Function}          [opts.onSelect]              选中回调函数
   * @param {Object|null}       [opts.onSelect.features]     选中feture的对象
   * @param {Object}            [opts.onSelect.mapEvt]       原始点击事件对象
   * @param {String}            [opts.onSelect.mapSvrName]   地图服务名(wms表中的name)，通过wfs选中时有值
   * @param {Function}          [opts.filterVecLyr]          矢量图层过滤器，有它时优先从矢量图层上选择feature，没有在设置的瓦片图层上选择
   * @param {ol.layer.Layer}    [opts.filterVecLyr.lyr]      矢量图层对象
   * @param {Boolean}           [opts.filterVecLyr.return]   判断结果
   */
  constructor(opts = {}) {
    this._map = opts.map
    this._map.on('singleclick', evt => this.onClickMap(evt))
    this.x2js = new X2js()
    this.onSelect = opts.onSelect
    this.filterVecLyr = opts.filterVecLyr
    this.lyrSource = null
    this.lyrEtt = null
    this.disable = false

    // 创建一个交互工具触发互斥
    this.selectIntact = InteractionFactory.createSelect({ layers: [] })
    this._map.addInteraction(this.selectIntact)
  }

  destroy() {
    this._map.un('singleclick', this.onClickMap)
    this._map.removeInteraction(this.selectIntact)
  }

  /**
   * 根据wms图层和坐标位置获取feature数据
   * @param {ol.Source.TileWMS}  wmsSource               待获取的wms source对象
   * @param {ol.Coordinate}      coor                    查询位置坐标，投影必须与wms source一致
   * @param {Number}             resolution              获取
   * @param {Object}             param                   getFeatureInfo参数，如INFO_FORMAT、QUERY_LAYERS等
   * @param {String}             [param.INFO_FORMAT]     默认 application/geojson
   * @returns {Promise}
   */
  getFeatureInfoByWms(wmsSource, coor, resolution, param = {}) {
    if (!param.INFO_FORMAT) {
      param.INFO_FORMAT = 'application/geojson'
    }

    const url = wmsSource.getFeatureInfoUrl(
      coor,
      resolution,
      wmsSource.getProjection(),
      param
    )
    return new Promise((resolve, reject) => {
      // 有wms的getFeatureInfo服务路径
      if (url) {
        window.cx.http
          .get(url)
          .then(ret => resolve(ret))
          .catch(err => reject(err))
      } else {
        resolve()
      }
    })
  }

  /**
   * 根据wms服务url获取wfs服务路径（arcgis、geoserver等平台的wms、wfs路径关系不同）
   * @param {String} wmsUrl      wms服务url
   * @returns {String|null}      wfs服务url
   */
  getWfsUrlByWms(wmsUrl) {
    let ret = null
    const url = wmsUrl.split('?')[0]
    const arrs = url.split('/')
    const lastItem = arrs[arrs.length - 1]
    if (lastItem && lastItem.toLowerCase() === 'wmsserver') {
      // arcgis服务
      ret = url.replace(lastItem, 'WFSServer')
    }
    return ret
  }

  /**
   * 根据wfs获取feature数据
   * @param {Object}  opts                        参数对象
   * @param {String}  opts.url                    wms服务url
   * @param {String}  [opts.geomFld='Shape']      空间字段名
   * @param {String}  opts.operType               过滤操作类型，区分大小写
   * @param {Object}  opts.geom                   条件值
   * @param {String}  opts.geom.type              空间图形的类型
   *    type值有：
   *        Box -坐标范围
   *        Point - 点
   *        LineString - 线串
   *        LinearRing - 线环
   *        Polygon - 多边形
   * @param {Array}   opts.geom.coors             各类图形的坐标序列
   * @param {String}  opts.geom.srsName           所提供坐标的投影
   * @returns {Promise}
   */
  getFeaturesByWfs(opts) {
    return new Promise((resolve, reject) => {
      OgcWFS.getCapabilities({ url: opts.url })
        .then(wfsHandle => {
          const lyrIds = this.lyrSource.getParams().LAYERS.split(',')
          const param = {
            url: opts.url,
            srsName: this.lyrSource.getProjection().getCode(),
            featureType: wfsHandle.getFeatureType()[lyrIds[0] || 0].name,
            version: wfsHandle.getVersion(),
            gmlFilter: OgcWFS.makeGmlFilterByGeom({
              geomFld: opts.geomFld || 'Shape',
              operType: opts.operType || 'BBox',
              geom: {
                type: opts.geom.type || 'Box',
                coors: opts.geom.coors,
                srsName: opts.geom.srsName
              }
            })
          }

          OgcWFS.getFeature(param)
            .then(ret => {
              resolve(ret)
            })
            .catch(err => reject(err))
        })
        .catch(err => reject(err))
    })
  }

  onClickMap(mapEvt) {
    const coor = mapEvt.coordinate

    if (!this.disable) {
      let feas = []
      // 1 先在非高亮的矢量图层上选择
      if (this.filterVecLyr) {
        // 优先矢量选择
        feas = this._map.getFeaturesAtPixel(mapEvt.pixel, {
          hitTolerance: 3,
          layerFilter: lyr => {
            return lyr.get('name') !== 'light' && lyr instanceof VectorLayer
              ? this.filterVecLyr(lyr)
              : false
          }
        })
      }

      if (feas.length > 0) {
        this.onSelect(feas, mapEvt)
      } else if (this.lyrSource) {
        // 如果没有选择就在wms图层选择
        const mapSrsName = this._map
          .getView()
          .getProjection()
          .getCode()
        const urls = this.lyrSource.getUrls()
        if (urls && urls[0]) {
          window.cx.base.processing.show()
          const url = this.getWfsUrlByWms(urls[0])

          // let bbox = ol.extent.buffer([coor[0], coor[1], coor[0], coor[1]], 3 * this._map.getView().getResolution())
          // let opts = {
          //   url,
          //   operType: 'BBOX',
          //   geom: {
          //     type: 'Box',
          //     srsName: 'EPSG:4326',
          //     coors: olProj.transformExtent(bbox, mapSrsName, 'EPSG:4326')
          //   }
          // }
          const opts = {
            url,
            operType: 'Intersects',
            geom: {
              type: 'Point',
              srsName: 'EPSG:4326',
              coors: olProj.transform(coor, mapSrsName, 'EPSG:4326')
            }
          }
          this.getFeaturesByWfs(opts)
            .then(ret => {
              window.cx.base.processing.close()
              if (this.onSelect) {
                const feas = new GeoJSON().readFeatures(ret.data, {
                  dataProjection: this.lyrSource.getProjection().getCode(),
                  featureProjection: mapSrsName
                })
                this.onSelect(feas, mapEvt, this.lyrSource.get('svrName'))
              }
            })
            .catch(() => window.cx.base.processing.close())
        }
      } else {
        this.onSelect([], mapEvt)
      }
    }
  }

  /**
   * 设置工具的禁止和激活状态
   * @param {Boolean}           disabled             是否禁用工具
   * @param {Function}          [onSelect]           选中回调函数
   * @param {Object}            onSelect.geoJson     选中空间对象的geoJson
   * @param {Object}            onSelect.mapEvt      原始点击事件对象
   */
  setDisable(disable, onSelect) {
    this.disable = disable

    // 根据激活状态修改控制选中控件，触发互斥
    !disable
      ? this._map.addInteraction(this.selectIntact)
      : this._map.removeInteraction(this.selectIntact)

    if (onSelect) {
      this.onSelect = onSelect
    }
  }

  /**
   * 设置选择图层元
   * @method
   * @param {ol.source.TileImage} lyrSource 待选择操作的图层source，主要为了获取投影和服务地质
   * @param {*} lyrEtt
   */
  setLyrSource(lyrSource, lyrEtt) {
    this.lyrSource = lyrSource
    this.lyrEtt = lyrEtt
  }
}

export default SelFeaByWFS
