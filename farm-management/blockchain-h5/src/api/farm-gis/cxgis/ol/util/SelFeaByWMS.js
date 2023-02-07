import axios from 'axios'
import VectorLayer from 'ol/layer/Vector'
import { TileImage, TileWMS } from 'ol/source'
import OgcWMS from '../../base/ogc/WMS'
import InteractionFactory from '../factory/InteractionFactory'

/**
 * cxgis.ol.util.SelFeaByWMS
 * 基于wms服务选择feature的工具类
 * @class
 * @memberof cxgis.ol.util
 * @author zhang qing
 * @since  2020-07-01
 */
class SelFeaByWMS {
  /**
   * 根据mapSvrUrl指定的arcgis服务，获取点击位置的feature对象
   * @param {Object}            opts                         IdentifyTask需要的配置参数对象,更多参数见esri.tasks.support.IdentifyParameters
   * @param {ol/Map}            opts.map                     view对象，默认使用gMap.view
   * @param {Function}          [opts.onSelect]              选中回调函数，可通过setDisable修改
   * @param {Object|null}       [opts.onSelect.features]     选中feture的对象
   * @param {Object}            [opts.onSelect.mapEvt]       原始点击事件对象
   * @param {Function}          [opts.filterVecLyr]          矢量图层过滤器，有它时优先从矢量图层上选择feature，没有在设置的瓦片图层上选择
   * @param {ol.layer.Layer}    [opts.filterVecLyr.lyr]      矢量图层对象
   * @param {Boolean}           [opts.filterVecLyr.return]   判断结果
   */
  constructor(opts = {}) {
    this._map = opts.map
    this._map.on('singleclick', evt => this.onClickMap(evt))
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

  onClickMap(mapEvt) {
    const coor = mapEvt.coordinate

    // 异常和成果处理函数
    const onErr = () => {
      window.cx.base.processing.close()
      this.onSelect([], mapEvt)
    }
    const onSuccess = ret => {
      window.cx.base.processing.close()
      this.onSelect(ret, mapEvt)
    }

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
        window.cx.base.processing.show()

        // 2 如果没有选择 且 指定待选图层
        let wmsUrl = null
        const projection = this._map.getView().getProjection()
        if (this.lyrSource instanceof TileWMS) {
          // 2.1 wms图层，直接获取调用url
          wmsUrl = this.lyrSource.getFeatureInfoUrl(
            coor,
            this._map.getView().getResolution(),
            projection,
            { INFO_FORMAT: 'text/html' }
          )
          axios
            .get(wmsUrl)
            .then(onSuccess)
            .catch(onErr)
        } else {
          // 2.2 其他图层，根据url构造请求参数
          const urls = this.lyrSource.getUrls()
          if (urls && urls[0]) {
            const bbox = window.ol.extent.buffer(
              [coor[0], coor[1], coor[0], coor[1]],
              3 * this._map.getView().getResolution()
            )
            OgcWMS.getFeatureInfoByUrl({
              url: urls[0],
              bbox: {
                coor: bbox,
                srsName: projection.getCode()
              }
            })
              .then(onSuccess)
              .catch(onErr)
          }
        }
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
   * 设置选择图层
   * @method
   * @param {ol.source.TileImage} lyrSource 待选择操作的图层source，主要为了获取投影和服务地址
   * @param {*} lyrEtt
   */
  setLyrSource(lyrSource, lyrEtt) {
    // 只有瓦片图层才有效
    if (lyrSource instanceof TileImage) {
      this.lyrSource = lyrSource
    }

    this.lyrEtt = lyrEtt
  }
}

export default SelFeaByWMS
