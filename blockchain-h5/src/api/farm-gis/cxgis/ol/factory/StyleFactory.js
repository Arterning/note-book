import { Fill, Stroke, Icon, Text, Circle, Style } from 'ol/style'

/**
 * cxgis.ol.factory.StyleFactory
 * 样式处理工厂
 * @class
 * @memberof cxgis.ol.factory
 * @author zhang qing
 * @since  2020-07-01
 */
class StyleFactory {
  /**
   * 创建圆形图标样式
   * @static
   * @method
   * @param {Object}        opts                           参数对象
   * @param {Number}        [opts.radius=5]                原型半径
   * @param {String}        [opts.fillColor='#231DCF']     填充色
   * @param {Number}        [opts.strokeWidth=1]           边线宽度
   * @param {String}        [opts.strokeColor='#FFFFFF']   边线颜色
   * @returns {ol/style/Style}                             样式对象
   */
  static createCircle(opts = {}) {
    opts.radius = opts.radius || 5
    opts.fillColor = opts.fillColor || '#231DCF'
    opts.strokeWidth = opts.strokeWidth || 1
    opts.strokeColor = opts.strokeColor || '#FFFFFF'

    return new Style({
      image: new Circle({
        radius: opts.radius,
        fill: new Fill({
          color: opts.fillColor
        }),
        stroke: new Stroke({
          color: opts.strokeColor,
          width: opts.strokeWidth
        })
      })
    })
  }

  /**
   * 创建通用样式
   * @static
   * @method
   * @param {Object}        opts                                         参数对象
   * @param {String}        [opts.fillColor='rgba(51, 51, 204, 0.3)]     填充色
   * @param {Number}        [opts.strokeWidth=1]                         边线宽度
   * @param {String}        [opts.strokeColor='rgba(4, 90, 141)']       边线颜色
   *
   * @param   {Object}      [opts.image]              完整图标样式对象
   *
   * @param   {Object}      [opts.icon]               图标样式
   * @param   {Object}      [opts.icon.color]         图标填充色
   * @param   {Object}      [opts.icon.scale]         图标缩放比例
   * @param   {Object}      [opts.icon.opacity]       图标不透明度，默认1
   * @param   {Object}      [opts.icon.src]           图标路径，默认dot
   *
   * @param   {Object}      [opts.text]               文本样式
   * @param   {String}      [opts.text.label]         文本内容
   * @param   {String}      [opts.text.fontSize]      文本字体大小，默认12px
   * @param   {String}      [opts.text.color]         文本字体颜色，默认黑色
   * @param   {String}      [opts.text.offsetX]       文本X偏移值，默认0
   * @param   {String}      [opts.text.offsetY]       文本Y偏移值，默认15
   * @param   {String}      [opts.text.textAlign]     文本对齐方式，默认居中
   * @param   {String}      [opts.text.strokeColor]   字体边框颜色，默认白色
   * @param   {String}      [opts.text.strokeWidth]   字体边框宽度，默认2
   * @returns {ol/style/Style}                                           样式对象
   */
  static createStyle(opts = {}) {
    const fillColor = opts.fillColor || 'rgba(51, 51, 204, 0.3)'
    opts.strokeWidth = opts.strokeWidth || 1
    opts.strokeColor = opts.strokeColor || 'rgba(4, 90, 141)'

    return new Style({
      fill: opts.fillColor
        ? new Fill({
            color: fillColor
          })
        : null,
      stroke: new Stroke({
        color: opts.strokeColor,
        width: opts.strokeWidth
      }),
      image:
        opts.image ||
        (opts.icon
          ? new Icon({
              color: opts.icon.color,
              crossOrigin: 'anonymous',
              scale: opts.icon.scale || 1,
              opacity: opts.icon.opacity || 1,
              src: opts.icon.src || 'static/img/pic/dot.png'
            })
          : null),
      text: opts.text
        ? new Text({
            font: opts.text.fontSize || '12' + 'px Calibri,sans-serif',
            text: opts.text.label || '',
            offsetX: opts.text.offsetX || 0,
            offsetY: opts.text.offsetY || 15,
            textAlign: opts.text.textAlign || 'center',
            fill: new Fill({ color: opts.text.color || '#000000' }),
            stroke: new Stroke({
              color: opts.text.strokeColor || '#ffffff',
              width: opts.text.strokeWidth || 2
            })
          })
        : null
    })
  }

  /**
   * 创建通用样式
   * @static
   * @method
   * @param {Object}        opts                                         参数对象
   * @param {String}        [opts.fillColor='rgba(51, 51, 204, 0.3)]     填充色
   * @returns {ol/style/Style}                                           样式对象
   */
  static createStyleFill(opts = {}) {
    opts.fillColor = opts.fillColor || 'rgba(51, 51, 204, 0.3)'

    return new Style({
      fill: new Fill({
        color: opts.fillColor
      })
    })
  }

  /**
   * 创建高亮闪烁的样式
   * @static
   * @method
   * @param {Object}        [opts]                                       高亮配置参数对象1，见createStyle函数参数
   * @returns {ol/style/Style}                                           两个样式对象
   */
  static createLight(opts = {}) {
    opts.fillColor = opts.fillColor || 'rgba(222,34,5,0.2)'
    opts.strokeColor = opts.strokeColor || 'rgba(222,34,1)'
    opts.strokeWidth = opts.strokeWidth || 2
    opts.image = opts.image || this.createCircle(opts).getImage()

    return this.createStyle(opts)
  }

  /**
   * 创建高亮闪烁边线的样式
   * @static
   * @method
   * @param {Object}        [opts]                                       高亮配置参数对象1
   * @param {Number}        [opts.strokeWidth=1]                         边线宽度
   * @param {String}        [opts.strokeColor='rgba(4, 90, 141)']        边线颜色
   * @returns {ol/style/Style}                                           两个样式对象
   */
  static createLightStroke(opts = {}) {
    opts.fillColor = opts.fillColor || 'rgba(222,34,5,0)'
    opts.strokeColor = opts.strokeColor || 'rgba(222,34,1)'
    opts.strokeWidth = opts.strokeWidth || 2

    return this.createStyle(opts)
  }

  /**
   * 创建高亮闪烁的样式
   * @static
   * @method
   * @param {Object}        [opts1]                                      闪烁配置参数对象1，见createStyle函数参数
   * @param {Object}        [opts2]                                      闪烁配置参数对象2，见createStyle函数参数
   * @returns {Array.<ol/style/Style>}                                   两个样式对象
   */
  static createLightFlicker(opts1 = {}, opts2 = {}) {
    opts1.fillColor = opts1.fillColor || 'rgba(222,34,5,0.2)'
    opts1.strokeColor = opts1.strokeColor || 'rgba(222,34,1)'
    opts1.strokeWidth = opts1.strokeWidth || 2
    opts1.image = opts1.image || this.createCircle(opts1).getImage()

    opts2.fillColor = opts2.fillColor || 'rgba(49,49,50,0.2)'
    opts2.strokeColor = opts2.strokeColor || 'rgba(49,49,50,1)'
    opts2.strokeWidth = opts2.strokeWidth || 2
    opts2.image = opts2.image || this.createCircle(opts2).getImage()

    return [this.createStyle(opts1), this.createStyle(opts2)]
  }
}

export default StyleFactory
