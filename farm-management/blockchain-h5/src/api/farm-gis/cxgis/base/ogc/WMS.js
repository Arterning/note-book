import X2js from 'x2js'
import axios from 'axios'
import * as olProj from 'ol/proj'
// import Arcgis from '../mapServer/Arcgis'

const x2js = new X2js()

/**
 * cxgis.base.ogc.WMS
 * 标准ogc wms处理类
 * @class
 * @memberof cxgis.base.ogc
 * @author zhang qing
 * @since  2020-07-01
 */
class WMS {
  /**
   * 根据wms服务url获取wfs服务路径（arcgis、geoserver等平台的wms、wfs路径关系不同）
   * @param {Object}          opts                                        wms服务配置对象
   * @param {String}          opts.url                                    wms服务地址，如:http://192.168.39.218:6080/arcgis/services/swdz/SWDZ_QG_500W/MapServer/WMSServer
   * @returns {Promise}
   */
  static getCapabilities(opts) {
    return new Promise((resolve, reject) => {
      axios
        .get(opts.url + '?request=GetCapabilities&service=WMS')
        .then(ret => {
          resolve(
            (() => {
              const data = x2js.xml2js(ret.data)
              return {
                getData: () => data,
                getVersion: () => data.WMS_Capabilities._version,
                getLayer: () =>
                  this._getLayer(data.WMS_Capabilities.Capability.Layer)
              }
            })()
          )
        })
        .catch(err => reject(err))
    })
  }

  /**
   * 根据wms参数信息获取FeatureInfo数据
   * @param {Object}          opts                                        wms服务配置对象
   * @param {String}          opts.url                                    wms服务地址，如http://192.168.39.239:6080/arcgis/services/swdz/ZBDXSLX2019/MapServer/WMSServer
   * @param {String}          opts.layers                                 查询的目标图层
   * @param {String}          opts.srsName                                wfs服务投影，如EPSG:4326
   * @param {Array.<Number>}  opts.bbox                                   bbox空间坐标[minx,miny,maxx,maxy]，坐标点位与srcname一致
   * @param {String}          [opts.format='image/png']                   返回地图格式
   * @param {String}          [opts.infoFormat='text/html']               返回数据格式
   * @param {String}          [opts.version='1.3.0']                      wms服务版本
   * @param {Boolean}         [opts.transparent=true]                     是否透明
   * @param {String}          [opts.styles]                               样式
   * @returns {Promise}
   */
  static getFeatureInfo(opts) {
    opts.version = opts.version || '1.3.0'
    opts.transparent = window.cx.base.defVal(opts.transparent, true)
    const axisOrientation = olProj.get(opts.srsName).getAxisOrientation()
    if (
      window.cx.base.compareVersions(opts.version, '1.3') >= 0 &&
      axisOrientation.substr(0, 2) === 'ne'
    ) {
      // 根据wms版本 和 轴方向设置bbox值顺序
      opts.bbox = [opts.bbox[1], opts.bbox[0], opts.bbox[3], opts.bbox[2]]
    }

    const wmsUrl =
      `${opts.url}?SERVICE=WMS&REQUEST=GetFeatureInfo` +
      `&VERSION=${opts.version}` +
      `&FORMAT=${opts.format || 'image/png'}` +
      `&QUERY_LAYERS=${opts.layers}` +
      `&LAYERS=${opts.layers}` +
      `&CRS=${opts.srsName}` +
      `&INFO_FORMAT=${opts.infoFormat || 'text/html'}` +
      `&TRANSPARENT=${opts.transparent}` +
      `&TILED=true` +
      `&WIDTH=256` +
      `&HEIGHT=256` +
      `&STYLES=${opts.styles || ''}` +
      `&BBOX=${opts.bbox}`
    return axios.get(wmsUrl).then(ret => {
      if (ret.data && ret.data.features) {
        // 返回json数据
        return ret.data.features.map(item => item.properties)
      } else {
        // 返回html，并解析
        const temp = x2js.xml2js(ret.data)
        return temp.html.body.table.map(item => {
          const flds = item.tbody.tr[0].th
          const vals = item.tbody.tr[1].td
          const ret = {}
          flds.forEach((col, idx) => {
            ret[col] = vals[idx]
          })
          return ret
        })
      }
    })
  }

  /**
   * 根据地图服务url获取wms的getFeatureInfo信息
   * @param {Object}          opts                                        wms服务配置对象
   * @param {String}          opts.url                                    地图服务地址，会格式化为标准wms地址
   * @param {Object}          opts.bbox                                   bbox空间坐标
   * @param {String}          opts.bbox.srsName                           bbox坐标点位投影，如EPSG:4326
   * @param {Array.<Number>}  opts.bbox.coor                              bbox空间坐标 [minx,miny,maxx,maxy]
   * @returns {Promise}
   */
  static getFeatureInfoByUrl(opts) {
    const wmsUrl = this.getFormatUrl(opts.url)
    return new Promise((resolve, reject) => {
      this.getCapabilities({ url: wmsUrl })
        .then(handle => {
          const layer = handle.getLayer()

          // 确保bbox坐标投影与地图一致
          let bbox = opts.bbox.coor
          if (layer.crs.epsg !== opts.bbox.srsName) {
            bbox = olProj.transformExtent(
              opts.bbox.coor,
              opts.bbox.srsName,
              layer.crs.epsg
            )
          }
          const param = {
            url: wmsUrl,
            srsName: layer.crs.epsg,
            layers: layer.layers.map(lyr => lyr.name).join(),
            bbox: bbox,
            version: handle.getVersion()
          }
          this.getFeatureInfo(param)
            .then(ret => {
              resolve(ret)
            })
            .catch(reject)
        })
        .catch(reject)
    })
  }
  /**
   * 根据url获取标准wms地址(各服务器的地址不同)
   * @method
   * @static
   * @param {String}          url        待格式处理的地图服务地址
   * @returns {String}                   发布的wms服务地址，如https://192.168.2.68:6443/arcgis/services/Taixing3D/dmt_wfs/MapServer/WMSServer
   */
  static getFormatUrl(url) {
    if (window.Arcgis.isArcMapSvr(url)) {
      return window.Arcgis.getUrlWMS(url)
    }
  }

  static _getLayer(data) {
    const layers = []
    // eslint-disable-next-line
    const findLayer = item => {
      if (item.Name) {
        layers.push(item)
      }
      if (item.Layer && item.Layer.Name !== undefined) {
        layers.push(item.Layer)
      }
      if (
        item.Layer &&
        window.cx.base.isArr(item.Layer) &&
        item.Layer.length !== 0
      ) {
        item.Layer.forEach(findLayer)
      }
    }
    const funcObj = {
      getCRS: function(data) {
        const ret = {}
        data.forEach(item => {
          ret[item.split(':')[0].toLowerCase()] = item
        })
        return ret
      },
      getBoundingBox: function(data) {
        const ret = []
        data.forEach(item => {
          ret.push({
            crs: item._CRS,
            extent: [
              Number(item._minx),
              Number(item._miny),
              Number(item._maxx),
              Number(item._maxy)
            ]
          })
        })
        return ret
      },
      getGeoBoundingBox: function(data) {
        return [
          Number(data.westBoundLongitude),
          Number(data.northBoundLatitude),
          Number(data.eastBoundLongitude),
          Number(data.southBoundLatitude)
        ]
      },
      getLayers(data) {
        const list = window.cx.base.isArr(data) ? data : [data]
        const getStyle = function(style) {
          return {
            name: style.Name,
            title: style.Title,
            legendURL: {
              format: style.LegendURL.Format,
              width: Number(style.LegendURL._width),
              height: Number(style.LegendURL._height),
              url: style.LegendURL.OnlineResource['_xlink:href']
            }
          }
        }
        return list.map(item => {
          return {
            name: item.Name,
            abstract: item.Abstract,
            title: item.Title,
            crs: this.getCRS(item.CRS),
            boundingBox: this.getBoundingBox(item.BoundingBox),
            ex_GeographicBoundingBox: this.getGeoBoundingBox(
              item.EX_GeographicBoundingBox
            ),
            style: getStyle(item.Style)
          }
        })
      }
    }
    return {
      crs: funcObj.getCRS(data.CRS),
      boundingBox: funcObj.getBoundingBox(data.BoundingBox),
      ex_GeographicBoundingBox: funcObj.getGeoBoundingBox(
        data.EX_GeographicBoundingBox
      ),
      layers: funcObj.getLayers(data.Layer)
    }
  }
}

export default WMS
