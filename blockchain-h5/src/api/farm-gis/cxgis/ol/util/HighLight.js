import VectorLayer from 'ol/layer/Vector'
import VectorSource from 'ol/source/Vector'
import Style from 'ol/style/Style'
import CircleStyle from 'ol/style/Circle'
import Fill from 'ol/style/Fill'
import Stroke from 'ol/style/Stroke'

/**
 * cxgis.ol.util.HightLight
 * 高亮工具类
 * @class
 * @memberof cxgis.ol.util
 * @author zhang qing
 * @since  2020-07-01
 */
class HightLight {
  /**
   * @constructor
   * @param {Object} opts                       参数对象
   * @param {ol/Map} opts.map                   地图对象
   * @param {String} [opts.cLight='#ffcc33']    高亮颜色，默认亮黄色
   * @param {String} [opts.cFlicker1='#ffcc33'] 闪烁颜色1，默认亮黄色
   * @param {String} [opts.cFlicker2='#f45c24'] 闪烁颜色2，默认亮红色
   */
  constructor(opts = {}) {
    this._map = opts.map
    this._lyr = new VectorLayer({
      map: this._map,
      source: new VectorSource({
        overlaps: true
      }),
      declutter: true
    })
    this._lyr.set('name', 'light')

    this._cLight = opts.cLight || '#ffcc33'
    this._cFlicker1 = opts.cFlicker1 || '#ffcc33'
    this._cFlicker2 = opts.cFlicker2 || '#f45c24'

    this._lights = {} // 保存已经高亮特效的对象
    this._lightsNoClone = {} // 保存已经高亮特效的非克隆原始对象
    this._locateBuff = 100 // 定位时，地图边界的缓冲大小，单位像素
  }

  destroy() {
    this.clear()
    this._lyr.setMap(null)
  }

  clear() {
    this._lyr.getSource().clear()
    this._lights = {}

    // 还原非克隆fea样式
    this._setOldStyle(this._lightsNoClone)
    this._lightsNoClone = {}
    if (this.timeout) {
      window.clearTimeout(this.timeout)
    }
  }

  flikerFea(fea, time = 4000) {
    if (this.timeout) {
      window.clearTimeout(this.timeout)
    }
    this.fea = fea
    this.setFlicker(fea)
    this.timeout = setTimeout(() => {
      window.clearTimeout(this.timeout)
      this.timeout = null
    }, time)
  }

  /**
   * 清除一个feature的高亮和闪烁效果
   * @method
   * @param {ol/Feature} feature 目标对象
   */
  remove(feature) {
    if (!feature) {
      return
    }

    const fid = this._getFid(feature)
    const del = fea => {
      // 有闪烁，停止闪烁
      const handle = fea.get('flicker')
      if (handle) {
        clearInterval(handle)
      }

      this._lyr.getSource().removeFeature(fea)
    }

    // 拷贝feature
    const ftCopy = this._lights[fid]
    if (ftCopy) {
      del(ftCopy)
      delete this._lights[fid]
    }

    // 非拷贝feature
    const ftNoCopy = this._lightsNoClone[fid]
    if (ftNoCopy) {
      del(ftNoCopy.feature)
      delete this._lightsNoClone[fid]
      feature.setStyle(ftNoCopy.oldStyle)
    }
  }

  /**
   * 闪烁一个feature
   * @method
   * @param {ol/Feature}      feature               目标对象
   * @param {ol/style.Style}  [style1]              高亮的样式
   * @param {ol/style.Style}  [style2]              高亮的样式
   * @param {Number}          [time]                闪烁间隔时间，毫秒
   * @param {Boolean}         [isLocate=false]      是否定位闪烁的对象（在可视窗口显示，缩放级别适中）
   * @param {Boolean}         [isClone=true]        是否采用克隆源fea方式
   */
  setFlicker(feature, style1, style2, time, isLocate, isClone = true) {
    if (!feature) {
      return
    }

    // 先移除对象的特效
    this.remove(feature)
    const fid = this._getFid(feature)
    let tarFea = feature

    if (isClone) {
      // 获取特效对象
      tarFea = feature.clone()
      this._lights[fid] = tarFea
    } else {
      // 获取特效对象
      const fea = this._lightsNoClone[fid]
      if (!fea) {
        this._lightsNoClone[fid] = {}
        this._lightsNoClone[fid].feature = feature
        this._lightsNoClone[fid].oldStyle = feature.getStyle()
      }
    }
    this._lyr.getSource().addFeature(tarFea)

    // 获取特效对象
    time = window.cx.base.isNNull(time) ? time : 500
    style1 = window.cx.base.isNNull(style1)
      ? style1
      : this._createDefaultStyles(this._cFlicker1)
    style2 = window.cx.base.isNNull(style2)
      ? style2
      : this._createDefaultStyles(this._cFlicker2)

    let turn = true
    const idFlicker = setInterval(() => {
      const style = turn ? style1 : style2
      turn = !turn
      tarFea.setStyle(style)
    }, time)

    tarFea.set('flicker', idFlicker) // 保存闪烁句柄

    if (isLocate) {
      this._locate(tarFea)
    }
  }

  /**
   * 闪烁一个feature
   * @method
   * @param {ol/Feature}      feature                     目标对象
   * @param {Object}          [param]                     目标对象
   * @param {ol/style.Style}  [param.style1]              高亮的样式
   * @param {ol/style.Style}  [param.style2]              高亮的样式
   * @param {Number}          [param.time]                闪烁间隔时间，毫秒
   * @param {Boolean}         [param.isLocate=false]      是否定位闪烁的对象（在可视窗口显示，缩放级别适中）
   * @param {Boolean}         [param.isClone=true]        是否采用克隆源fea方式
   * @param {Number}          [param.clearTime]           自动清除闪烁的延迟时间，没有值时不自动清除
   */
  setFlickerFea(feature, param = {}) {
    this.fea = feature
    this.setFlicker(
      feature,
      param.style1,
      param.style2,
      param.time,
      param.isLocate,
      param.isClone
    )

    if (param.clearTime) {
      if (this.timeout) {
        window.clearTimeout(this.timeout)
      }

      this.timeout = setTimeout(() => {
        this.setLight(feature)
        window.clearTimeout(this.timeout)
        this.timeout = null
      }, param.clearTime)
    }
  }

  /**
   * 高亮一个feature
   * @method
   * @param {ol/Feature}      feature          目标对象
   * @param {ol/style/Style}  [style]          高亮的样式
   * @param {Boolean}         [isLocate=false] 是否定位闪烁的对象（在可视窗口显示，缩放级别适中）
   * @param {Boolean}         [isClone=true]   是否采用克隆源fea方式
   */
  setLight(feature, style, isLocate, isClone = true) {
    if (!feature) {
      return
    }

    // 先移除对象的特效
    this.remove(feature)
    const fid = this._getFid(feature)
    let tarFea = feature
    style = window.cx.base.isNNull(style)
      ? style
      : this._createDefaultStyles(this._cLight, feature.getStyle())

    if (isClone) {
      // 获取特效对象
      tarFea = feature.clone()
      this._lights[fid] = tarFea
    } else {
      // 获取特效对象
      const fea = this._lightsNoClone[fid]
      if (!fea) {
        this._lightsNoClone[fid] = {}
        this._lightsNoClone[fid].feature = feature
        this._lightsNoClone[fid].oldStyle = feature.getStyle()
      }
    }
    this._lyr.getSource().addFeature(tarFea)
    tarFea.setStyle(style)

    if (isLocate) {
      this._locate(feature)
    }
  }

  /**
   * 根据目标颜色，和参考style对象，创建默认style
   * @private
   * @method
   * @param {String} color 目标颜色
   * @param {ol/style/Style} 参考style对象
   * @returns {ol/style/Style} 创建的样式
   */
  _createDefaultStyles(color, oldStyle) {
    var textStyle = null
    if (oldStyle && !window.cx.base.isFunc(oldStyle)) {
      oldStyle = window.cx.base.isArr(oldStyle) ? oldStyle[0] : oldStyle
      var oldTexStyle = oldStyle.getText()
      if (oldTexStyle) {
        textStyle = oldTexStyle.clone()
        textStyle.setStroke(
          new Stroke({
            color: color,
            width: 1
          })
        )
      }
    }
    return new Style({
      fill: new Fill({ color: color }),
      stroke: new Stroke({
        color: color,
        width: 2
      }),
      image: new CircleStyle({
        radius: 5,
        fill: new Fill({
          color: color
        })
      }),
      text: textStyle
    })
  }

  /**
   * 根据feature做fid
   * @private
   * @method
   * @param {ol/Feature} feature 目标feature对象
   * @returns {String} fid值
   */
  _getFid(feature) {
    return feature.get('id')
  }

  /**
   * 定位一个feature。当对象不在可视窗口时居中显示，然后把地图级别缩放到适中，可以清楚发现该对象
   * @private
   * @method
   * @param {ol/Feature} feature 目标对象
   */
  _locate(feature) {
    var mapExtent = this._map.getView().calculateExtent(this._map.getSize())
    var feaExtent = feature.getGeometry().getExtent()
    var mapBuff = this._locateBuff * this._map.getView().getResolution()

    // 对定位的边界按缓冲区缩小
    mapExtent = [
      mapExtent[0] + mapBuff,
      mapExtent[1] + mapBuff,
      mapExtent[2] - mapBuff,
      mapExtent[3] - mapBuff
    ]

    // 定位的目标在显示区域外时，移动到地图进行定位
    if (
      feaExtent[0] < mapExtent[0] ||
      feaExtent[1] < mapExtent[1] ||
      feaExtent[2] > mapExtent[2] ||
      feaExtent[3] > mapExtent[3]
    ) {
      this._map
        .getView()
        .setCenter([
          (feaExtent[2] + feaExtent[0]) / 2,
          (feaExtent[1] + feaExtent[3]) / 2
        ])
    }
  }

  /**
   * 还原原始fea样式
   * @private
   * @method
   * @param {Object}  lightsNoClone          非克隆数据对象
   */
  _setOldStyle(lightsNoClone) {
    for (const i in lightsNoClone) {
      lightsNoClone[i].feature.setStyle(lightsNoClone[i].oldStyle)
    }
  }
}

export default HightLight
