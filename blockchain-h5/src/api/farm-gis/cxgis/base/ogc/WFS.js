import X2js from 'x2js'
import axios from 'axios'
import WMS from './WMS'
// import Arcgis from '../mapServer/Arcgis'

const x2js = new X2js()

/**
 * cxgis.base.ogc.WFS
 * 标准ogc wfs处理类
 * @class
 * @memberof cxgis.base.ogc
 * @author zhang qing
 * @since  2020-07-01
 */
class WFS {
  /**
   * 根据wms服务url获取wfs服务路径（arcgis、geoserver等平台的wms、wfs路径关系不同）
   * @param {Object}          opts                                        wfs服务配置对象
   * @param {String}          opts.url                                    wfs服务地址，如https://192.168.2.68:6443/arcgis/services/Taixing3D/dmt_wfs/MapServer/WFSServer
   * @returns {Promise}
   */
  static getCapabilities(opts) {
    return new Promise((resolve, reject) => {
      axios
        .get(opts.url + '?request=GetCapabilities&service=WFS')
        .then(ret => {
          WMS.getCapabilities({ url: window.Arcgis.getUrlWMS(opts.url) }).then(
            retWms => {
              const layers = retWms ? retWms.getLayers() : []
              resolve(
                (() => {
                  const data = x2js.xml2js(ret.data)
                  const getFeatureType = () => {
                    // 多个type时时数字，单个时是对象
                    const types =
                      data.WFS_Capabilities.FeatureTypeList.FeatureType
                    const func = item => {
                      const coors = (
                        item.WGS84BoundingBox.LowerCorner.toString() +
                        ' ' +
                        item.WGS84BoundingBox.UpperCorner.toString()
                      ).split(' ')
                      return {
                        name: item.Name.toString(),
                        title: item.Title.toString(),
                        crs: item.DefaultCRS.toString(),
                        srs:
                          'EPSG:' +
                          item.DefaultCRS.toString()
                            .split('::')
                            .pop(),
                        wgs84BoundingBox: coors.map(coor => Number(coor))
                      }
                    }

                    return window.cx.base.isArr(types)
                      ? types.map(item => func(item))
                      : [func(types)]
                  }
                  return {
                    getData: () => data,
                    getVersion: () =>
                      data.WFS_Capabilities.ServiceIdentification.ServiceTypeVersion.toString(),
                    getLayers: () => layers,
                    getFeatureType,
                    getFeatureTypeByLayer: layer => {
                      const selLayer = layers.find(
                        item => Number(item.Name) === Number(layer)
                      )
                      const selFeatreType = selLayer
                        ? getFeatureType().find(
                            item => item.title === selLayer.Title
                          )
                        : null
                      return selFeatreType ? selFeatreType.name : null
                    },
                    getOperGetFeature: () => {
                      const ret = {}
                      const conf = data.WFS_Capabilities.OperationsMetadata.Operation.find(
                        item => item._name === 'GetFeature'
                      )
                      ret.http = {
                        get: conf.DCP.HTTP.Get['_xlink:href'],
                        post: conf.DCP.HTTP.Post['_xlink:href']
                      }

                      const outConf = conf.Parameter.find(
                        item => item._name === 'outputFormat'
                      )
                      ret.outputFormat = outConf.AllowedValues.Value.map(item =>
                        item.toString()
                      )
                      return ret
                    }
                  }
                })()
              )
            }
          )
        })
        .catch(err => reject(err))
    })
  }

  // /**
  //  * 构造AND属性过滤的xml过滤条件
  //  * @static
  //  * @param {String}  filter1                     条件1
  //  * @param {String}  filter2                     条件2
  //  * @returns {String}                            构造好的and条件
  //  */
  // static filerAnd(filter1, filter2)
  // {
  //   return `<ogc:And>${filter1}${filter2}</ogc:And>`
  // }

  // /**
  //  * 构造属性过滤的xml过滤条件
  //  * @static
  //  * @param {Object}  opts                        属性过滤参数对象
  //  * @param {String}  opts.fldName                条件字段名
  //  * @param {String}  opts.value                  条件值
  //  * @param {String}  opts.operType               过滤操作类型
  //  *    operType值有：
  //  *        EqualTo
  //  *        NotEqualTo
  //  *        GreaterThan
  //  *        LessThanOrEqualTo
  //  *        GreaterThanOrEqualTo
  //  *        Like
  //  *        Null
  //  *        Nil
  //  *        Between
  //  * @returns {String}                            构造好的xml过滤条件
  //  */
  // static makeGmlFilterByProp(opts)
  // {
  //   switch (opts.operType)
  //   {
  //   case 'Like':
  //     return `<ogc:PropertyIs${opts.operType} wildCard="*" singleChar="." escapeChar="!"><ogc:PropertyName>${opts.fldName}</ogc:PropertyName><ogc:Literal>*${opts.value}*</ogc:Literal></ogc:PropertyIs${opts.operType}>`
  //   default:
  //     return `<ogc:PropertyIs${opts.operType}><ogc:PropertyName>${opts.fldName}</ogc:PropertyName><ogc:Literal>${opts.value}</ogc:Literal></ogc:PropertyIs${opts.operType}>`
  //   }
  // }

  // /**
  //  * 构造空间过滤的xml过滤条件
  //  * @static
  //  * @param {Object}  opts                        属性过滤参数对象
  //  * @param {String}  [opts.geomFld='Shape']      空间字段名
  //  * @param {Object}  opts.geom                   条件值
  //  * @param {String}  opts.geom.type              空间图形的类型
  //  *    type值有：
  //  *        Box -坐标范围
  //  *        Point - 点
  //  *        LineString - 线串
  //  *        LinearRing - 线环
  //  *        Polygon - 多边形
  //  * @param {Array}   opts.geom.coors             各类图形的坐标序列
  //  * @param {String}  opts.geom.srsName           所提供坐标的投影
  //  * @param {String}  opts.operType               过滤操作类型，区分大小写
  //  *    operType值有：
  //  *        BBOX  仅适用于 Envelope或Box图形
  //  *        Equals
  //  *        Disjoint
  //  *        Intersects
  //  *        Touches
  //  *        Crosses
  //  *        Within
  //  *        Contains
  //  *        Overlaps
  //  *        Beyond
  //  *        DWithin
  //  * @returns {String}                            构造好的xml过滤条件
  //  */
  // static makeGmlFilterByGeom(opts)
  // {
  //   let gmlGeom = ''
  //   opts.geomFld = opts.geomFld || 'Shape'
  //   switch (opts.geom.type)
  //   {
  //   case 'Box':
  //     gmlGeom = `<gml:coordinates>${opts.geom.coors.join(',')}</gml:coordinates>`
  //     break
  //   case 'Point':
  //     gmlGeom = `<gml:coordinates>${opts.geom.coors.join(',')}</gml:coordinates>`
  //     break
  //   case 'LineString':
  //     gmlGeom = `<gml:coordinates>${opts.geom.coors.map(coor => coor.toString()).join(' ')}</gml:coordinates>`
  //     break
  //   case 'LinearRing':
  //     gmlGeom = `<gml:coordinates>${opts.geom.coors.map(coor => coor.toString()).join(' ')}</gml:coordinates>`
  //     break
  //   case 'Polygon':
  //     gmlGeom = `<gml:outerBoundaryIs><gml:LinearRing><gml:coordinates>${opts.geom.coors[0].map(coor => coor.toString()).join(' ')}</gml:coordinates></gml:LinearRing></gml:outerBoundaryIs>`
  //     break
  //   }
  //   let xmlFilter = `<ogc:${opts.operType}>` +
  //     `<ogc:PropertyName>${opts.geomFld}</ogc:PropertyName>` +
  //     `<gml:${opts.geom.type} srsName="urn:x-ogc:def:crs:${opts.geom.srsName}">` +
  //         gmlGeom +
  //     `</gml:${opts.geom.type}>` +
  //   `</ogc:${opts.operType}>`

  //   return xmlFilter
  // }

  /**
   * 根据wms服务url获取wfs服务路径（arcgis、geoserver等平台的wms、wfs路径关系不同）
   * @param {Object}          opts                                        wfs服务配置对象
   * @param {String}          opts.url                                    wfs服务地址，如https://192.168.2.68:6443/arcgis/services/Taixing3D/dmt_wfs/MapServer/WFSServer
   * @param {String}          opts.featureType                            featuretype名，Taixing3D_zq_dr_wfs:dr
   * @param {String}          opts.srsName                                wfs服务投影，如EPSG:4326
   * @param {String}          [opts.outputFormat='GEOJSON']               返回数据格式，默认geojson
   * @param {String}          [opts.version='2.0.0']                      wfs服务版本
   * @param {Number}          [opts.pageno=1]                             开始索引号。指定结果集在响应文档中开始呈现的索引号。
   * @param {Number}          [opts.pagesize=25]                          返回最大数据量，默认全部
   * @param {Array.<Object>}  [opts.gmlFilter]                            过滤条件配置
   * @returns {Promise}
   */
  static getFeature(opts) {
    opts.version = opts.version || '2.0.0'
    opts.geomFld = opts.geomFld || 'Shape'
    opts.outputFormat = opts.outputFormat || 'GEOJSON'
    const count = (opts.pageno || 1) * (opts.pagesize || 25)
    opts.startIndex = count - (opts.pagesize || 25)
    const data =
      `<GetFeature ` +
      `xmlns:ogc="http://www.opengis.net/ogc" ` +
      `xmlns:wfs="http://www.opengis.net/wfs" ` +
      `xmlns:fes="http://www.opengis.net/fes/2.0" ` +
      `xmlns:gml="http://www.opengis.net/gml" ` +
      `xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" ` +
      `xmlns:myns="http://www.someserver.com/myns" ` +
      `xsi:schemaLocation="http://www.opengis.net/wfs/2.0 http://schemas.opengis.net/wfs/2.0/wfs.xsd" ` +
      // `xsi:schemaLocation="http://www.opengis.net/wfs http://schemas.opengis.net/wfs/1.1.0/wfs.xsd" ` +
      `service="WFS" version="${opts.version}" outputFormat="${
        opts.outputFormat
      }" startIndex="${
        opts.startIndex
      }" count="${count}" maxFeatures="${opts.pagesize || 25}">` +
      `<wfs:Query typeName="${opts.featureType}" srsName="${opts.srsName}">` +
      `<ogc:Filter>` +
      opts.gmlFilter +
      `</ogc:Filter>` +
      `</wfs:Query>` +
      `</GetFeature>`
    const conf = {
      headers: { 'Content-Type': 'application/xml' }
    }
    return axios.post(opts.url, data, conf)
  }
}

export default WFS
