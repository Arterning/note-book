<template>
  <div style="width: 100%;height: 100%;">
    <div id="map" ref="mapContOl" style="width: 100%;height: 100%;"></div>
    <slot></slot>
  </div>
</template>

<script>
/**
 * openlayers的通用地图组件
 */
import 'ol/ol.css'
import { query } from '@/api/farm-gis/AttBase'

export default {
  props: {
    defLyrs: { type: Array, default: () => ['imgMap'] }, // 默认加载的图层名序列
    topLyrs: { type: Array, default: () => [] }, // 总是在最上层的图层名序列
    actSelType: { type: String, default: 'wms' } // 选择工具类型，wms和wfs
  },
  mounted() {
    // window.cx = require('@/api/farm/cx')
    this.makeSvrData().then(() => {
      // 延迟创建地图，根据地图容器尺寸初始化地图显示级别，否则dom对象的style还没渲染完毕
      setTimeout(() => this.createMap())
    })
  },
  beforeDestroy() {
    if (this.light) {
      this.light.destroy()
    }
  },
  methods: {
    createMap() {
      window.cxgis.ol.factory.MapFactory.extendProj(
        this.syscfg.find(item => item.keyno === 'srsAddDef')
      )
      const mapObj = window.cxgis.ol.factory.MapFactory.create({
        target: this.$refs.mapContOl,
        layers: this.defLyrs.map(name =>
          window.cxgis.ol.factory.LyrTileFactory.createByType(
            this.wms.find(item => item.name === name)
          )
        ),
        srsSys: this.syscfg.find(item => item.keyno === 'srsSys').value,
        viewData: this.viewport[0]
      })

      this.map = mapObj.map
      this.view = mapObj.view

      // 创建高亮对象
      this.light = new window.cxgis.ol.util.HighLight({
        map: this.map,
        cLight: '#D40095'
      })

      // 创建图层鼠标选中工具对象
      const ConstClass =
        this.actSelType === 'wms'
          ? window.cxgis.ol.util.SelFeaByWMS
          : window.cxgis.ol.util.SelFeaByWFS
      this.actSelectFeature = new ConstClass({
        map: this.map,
        onSelect: this.onSelFeature
      })
      // this.$bus.$emit(
      //   'mapInteractionMutex',
      //   evt => this.onInteractionMutex(evt),
      //   this.map.getInteractionMan()
      // )

      this.$emit('created', { component: this, map: this.map, view: this.view })
    },

    /**
     * 构造服务查询参数对象
     * @method
     * @param   {Number}    major   主类型
     * @param   {Number}    minor   子类型
     * @param   {Object}    [exp]   查询条件
     * @returns {Object}    查询参数对象
     */
    makeParam(major, minor, exp) {
      return {
        _major: major,
        _minor: minor,
        _exp: exp || '1=1'
      }
    },

    /**
     * 查询服务获取数据
     * @method
     * @returns {Promise}
     */
    makeSvrData() {
      return new Promise((resolve, reject) => {
        Promise.all([
          query(this.makeParam(99, 2)),
          query(this.makeParam(99, 34)),
          query(this.makeParam(99, 31))
        ])
          .then(([syscfg, viewport, wms]) => {
            this.syscfg = syscfg
            this.viewport = viewport
            this.wms = wms
            window.cx.cache = window.cx.cache || {}
            window.cx.cache.wms = wms
            resolve([viewport, syscfg, wms])
          })
          .catch(reject)
      })
    },

    /**
     * 重新计算地图视口大小
     */
    resize() {
      this.map.updateSize()
    },

    /**
     * 获取右键菜单对象
     * @method
     * @returns {Vue}  右键菜单对象
     */
    getContextMenu() {
      return this._contextmenu
    },

    /**
     * 获取地图对象
     * @returns {ol/Map}      地图对象
     */
    getMap() {
      return this.map
    },

    /**
     * 获取视图对象
     * @returns {ol/View}     视图对象
     */
    getView() {
      return this.view
    },

    /**
     * 获取通过wms选择feature的交互工具对象
     * @returns {map/ol/util/SelFeaByWMS}      地图对象
     */
    getActSelectFeature() {
      return this.actSelectFeature
    },

    /**
     * 获取高亮工具对象
     * @returns {map/ol/util/HighLight}      地图对象
     */
    getCtrLight() {
      return this.light
    },

    onInteractionMutex(evt) {
      if (evt.interaction === this.actSelectFeature) {
        this.actSelectFeature.setDisable(true)
      }
    },
    /**
     * 地图点击响应函数
     * @method
     * @param {Object|null}      features       选中的feature对象
     * @param {Object}           evt            始点击事件对象
     */
    onSelFeature(features) {
      this.light.clear()
      if (features[0]) {
        const styles = window.cxgis.ol.factory.StyleFactory.createLightFlicker()
        this.light.setFlickerFea(features[0], {
          style1: styles[0],
          style2: styles[1]
        })
      }

      this.$emit('selectFeature', features[0])
    }
  }
}
</script>

<style lang="scss" scoped>
* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}
</style>
