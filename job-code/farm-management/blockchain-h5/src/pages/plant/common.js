import cxgis from '@/api/farm-gis/cxgis'
import { query } from '@/api/farm-gis/AttBase'
import gcoord from 'gcoord'

const common = {
  data() {
    return {
      defLyrs: ['tianImg_c', 'tianCia_c', 'wxyx'],
      show: false,
      selArea: null
    }
  },

  beforeCreate() {
    console.log(JSON.stringify(new Date()))
  },

  methods: {
    /**
     * 创建地图
     * @method
     */
    onCreatedMap({ component, map, view }) {
      this.lyrs = {}
      this.map = map
      this.view = view
      this.mapCom = component
      this.light = this.mapCom.getCtrLight()
      this.actSelFeaByWMS = this.mapCom.getActSelectFeature()

      // 创建矢量图层
      this.curLyr = cxgis.ol.factory.LyrVectorFactory.createVector({
        map: this.map,
        name: 'lyr'
      })

      this.map.getView().setZoom(14)
      this.map.getView().setCenter([118.29133310661416, 31.162736074366197])
      window.gMap = this
      // this.init()
    },

    /**
     * 重构地块坐标
     */
    coorFilter(arr) {
      if (!arr || arr.length === 0) return []

      //首先识别是否polygon情况
      let filter = []
      arr.forEach(data => {
        let geoJson = JSON.parse(data.geometryJson)

        if (geoJson.type === 'Polygon') {
          data.coor = geoJson.coordinates
          filter.push(data)
        } else if (geoJson.type === 'MultiPolygon') {
          data.coor = geoJson.coordinates[0]
          filter.push(data)
        } else if (geoJson.type === 'LineString') {
          data.coor = [geoJson.coordinates] //linestring情况
          if (Array.isArray(data.coor[0][0])) {
            filter.push(data)
          }
        }
      })

      return filter
    },

    /**
     * 为图层添加互动(选中事件)
     */
    setActSel() {
      this.actSel = cxgis.ol.factory.InteractionFactory.createSelect({
        onSelect: this.onSelectFea,
        layers: [this.curLyr],
        style: feature =>
          window.cxgis.ol.factory.StyleFactory.createStyle({
            fillColor: this.getActFeaColor(feature) || 'rgb(253, 46, 199, 0.3)',
            strokeWidth: 1,
            strokeColor: '#fff',
            text: {
              label: feature.getProperties().code,
              fontSize: '18px Calibri,sans-serif',
              color: '#fff',
              strokeColor: '#fff',
              strokeWidth: 1,
              offsetY: -3
            }
          })
      })
      this.map.addInteraction(this.actSel)
    },

    /**
     * 选中feature响应
     * @method
     * @param  {Object}    evt      事件对象
     */
    async onSelectFea(evt) {
      this.show = false
      this.setOverLay()
      this.popup.setPosition()
      this.light.clear()
      this.selFea = evt.selected[0]
      if (!this.selFea) {
        return false
      }
      this.$bus.$emit('onSelectFeature', this.selFea, att => {
        this.selArea = att
        this.popup.setPosition(this.getPopupPosition(evt.selected[0]))
      })
    },

    /**
     * 设置地图的overlay
     * @method
     */
    setOverLay() {
      this.popup = cxgis.ol.factory.OverlayFactory.createOverlay({
        element: this.$refs.overlay.$el,
        positioning: 'bottom-center',
        offset: [0, -25]
      })
      this.map.addOverlay(this.popup)
    },

    /**
     * 获取featrue的伪中心点坐标
     * @param fea
     * @returns {number[]}
     */
    getPopupPosition(fea) {
      const extent = fea.getGeometry().getExtent()
      return [(extent[0] + extent[2]) / 2, (extent[1] + extent[3]) / 2]
    },

    /**
     * 获取对应农场的坐标
     * @method
     */
    getFarmCoors(farmID) {
      const param = {
        _major: 10,
        _minor: 16,
        _exp: 'id>? and farmid=?',
        types: 'i,s',
        vals: `0,${farmID}`
      }
      return query(param)
    },

    /**
     * @description: 将百度坐标转换成天地图坐标
     * @param {Array} position 经纬度，如：[112, 23]
     * @return {Array} 经纬度，如：[112, 23]
     * @author: Hemingway
     */
    transformLocate1(position) {
      return gcoord.transform(position, gcoord.BD09, gcoord.WGS84)
    },

    /**
     * @description: 将GCJ02坐标转换成天地图坐标
     * @param {Array} position 经纬度，如：[112, 23]
     * @return {Array} 经纬度，如：[112, 23]
     * @author: Hemingway
     */
    transformLocate(position) {
      return gcoord.transform(position, gcoord.GCJ02, gcoord.WGS84)
    },

    /**
     * @description: 将百度坐标坐标转换成GCJ02坐标
     * @param {Array} position 经纬度，如：[112, 23]
     * @return {Array} 经纬度，如：[112, 23]
     * @author: Hemingway
     */
    transformLocate2(position) {
      return gcoord.transform(position, gcoord.BD09, gcoord.GCJ02)
    }
  }
}

export default common
