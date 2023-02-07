/** 绘制矢量feature的基类 */
export default {
  beforeDestroy() {
    this.clearMapVec()
  },
  methods: {
    /**
     * 清空图层上的Features
     */
    clearMapVec() {
      window.gMap.curLyr.getSource().clear()
      window.gMap.light.clear()
    },

    /**
     * 根据传入的Features在图层上渲染地块/点/线
     * @param data
     */
    drawFeas(data) {
      this.clearMapVec()
      const feas = []
      const clickfeas = []
      for (const item of data) {
        item.coor = item.coor || []
        // 将区域对应的公报数据放入进去
        const fea = this.makeFea({
          ...item,
          att: item,
          geoType: item.geoType || window.cx.consts.ENTITY_REG
        })
        // 将大屏的虫情灯单独放一个图层，不点击
        if (item.ly && item.ly === 'onclick') {
          clickfeas.push(fea)
        } else {
          feas.push(fea)
        }
      }
      this.feas = feas
      setTimeout(() => {
        window.gMap.curLyr.getSource().addFeatures(feas)
      })
    },

    /**
     * Feature样式中text的label
     * @param fea
     * @returns {*}
     */
    getFeaLabel(fea) {
      return fea.get('code')
    },

    /**
     * Feature的填充色
     * @param fea
     * @returns {string}
     */
    getFillColor() {
      return 'rgb(38, 202, 101, 0.6)'
    },

    locateView(feas = []) {
      window.cxgis.ol.factory.MapFactory.locateView(this.map, feas)
    },

    /**
     * 根据查询到的数据构造feature
     * @method
     * @param   {Object}     item                   查询到的实体数据对象
     * @param   {String || Number}  item.geoType    feature类型，可以是数字1、2/3/4，也可以是字符串: Point、Polygon、MultiPolygon
     * @returns {ol.Feature} 构造好的feature
     */
    makeFea(item) {
      const fea =
        item.geoType && !Number(item.geoType)
          ? window.cxgis.ol.factory.FeatureFactory.createByType(
              item,
              item.geoType
            )
          : window.cxgis.ol.factory.FeatureFactory.createFea(item, item.geoType)
      this.styleFea(fea, item)
      fea.setId(fea.get('id'))
      return fea
    },

    /**
     * 渲染feature对象
     * @method
     * @param {ol/Feature}   feature     feature对象
     */
    styleFea(feature) {
      //const fontColor = '#fff'
      feature.setStyle(() =>
        window.cxgis.ol.factory.StyleFactory.createStyle({
          fillColor: this.getFillColor(feature) || 'rgb(253, 46, 199, 0.3)',
          strokeWidth: 1,
          strokeColor: '#fff',
          // text: {
          //   label: feature.getProperties().label || this.getFeaLabel(feature),
          //   fontSize: '18px Calibri,sans-serif',
          //   color: fontColor,
          //   strokeColor: fontColor,
          //   strokeWidth: 1,
          //   offsetY: -3
          // },
          icon:
            feature
              .getGeometry()
              .getType()
              .toLowerCase() !== 'point'
              ? null
              : {
                  crossOrigin: 'anonymous',
                  scale: feature.getProperties().scale || 1,
                  opacity: feature.getProperties().opacity || 1,
                  src: feature.getProperties().pic
                }
        })
      )
    }
  }
}
