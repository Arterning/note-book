/**
 * cxgis.ol.control.InteractionMan
 * 地图交互控件的管理类。监听map的interaction添加删除操作，对有mutexType属性的控件进行同类控件互斥管理,并触发mapInteractionMutex消息
 * @class
 * @memberof cxgis.ol.control
 * @author zhang qing
 * @since  2020-07-01
 */
class InteractionMan {
  /**
   * @constructor
   * @param {Object}                  options                      参数对象
   * @param {ol.Map}                  options.map                  控件管理对象所属的map对象
   */
  constructor(options) {
    options = options || {}

    /**
     * 控件管理对象所属的map对象
     * @private
     * @property {ol.Map} _map
     */
    this._map = options.map

    /**
     * map对象的interaction集合对象
     * @private
     * @property {ol.Map} _map
     */
    this._interactCollection = this._map.getInteractions()
    this._interactCollection.on('add', evt => this.add(evt.element))
    this._interactCollection.on('remove', evt => this.remove(evt.element))

    /**
     * 控件的对象存储中心,结构为{type1:[],type2:[]}
     * @private
     * @property {Object} _interacts
     */
    this._interacts = {}

    /**
     * 标记是否由添加控件，导致交互控件清空
     * @private
     * @property {Boolean} _isClearByAdd
     */
    this._isClearByAdd = false
  }
  /**
   * 销毁管理对象
   * @method
   */
  destroy() {
    this._interactCollection.un('add', evt => this.add(evt.element))
    this._interactCollection.un('remove', evt => this.remove(evt.element))
  }
  /**
   * 清空控件管理中心存储的所有控件数据
   * @method
   */
  clear() {
    for (var i in this._interacts) {
      this.clearByType(i)
    }

    this._interacts = {}
  }
  /**
   * 从控件管理中心清空type类别的所有控件数据
   * @method
   * @param {String} type 待清空控件的类别
   */
  clearByType(type) {
    this._interacts[type] = this._interacts[type] || []
    for (var i = this._interacts[type].length - 1; i >= 0; i--) {
      this.remove(this._interacts[type][i])
    }

    // 还原标记
    this._isClearByAdd = false
  }
  /**
   * 获取控件中心所属的地图对象
   * @method
   * @returns {ol.Map} 地图对象
   */
  getMap() {
    return this._map
  }
  getMutexType(interaction) {
    return interaction.get('mutexType')
  }
  /**
   * 在管理中心添加交互控件
   * @method
   * @param {cx2.core.control.IActControl} interaction 待添加的控件对象。清空同类控件再添加
   */
  add(interaction) {
    var type = this.getMutexType(interaction)
    if (type) {
      this._interacts[type] = this._interacts[type] || []
      // 同类互斥，先清空
      this._isClearByAdd = true
      this.clearByType(type)

      // 添加新控件
      interaction.setActive(true)
      this._interacts[type].push(interaction)
    }
  }
  /**
   * 从控件中心删除指定的控件
   * @method
   * @param {cx2.core.control.IActControl} interaction 待删除的控件对象
   */
  remove(interaction) {
    var type = this.getMutexType(interaction)
    if (type) {
      this._interacts[type] = this._interacts[type] || []
      for (var i = this._interacts[type].length - 1; i >= 0; i--) {
        if (this._interacts[type][i] === interaction) {
          window.cx.array.removeAt(this._interacts[type], i, 1)
          interaction.setActive(false)

          if (this._isClearByAdd) {
            // 添加情况下 互斥删除
            window.vm.$bus.$emit(
              'mapInteractionMutex',
              { type: type, interaction: interaction },
              this
            )
          } else if (this._interacts[type].length <= 0) {
            // 非添加情况下 清空一类交互控件，广播消息
            /**
             * 交互管理中某类控件清空的消息事件
             * @event cx2.core.control.ActMan#actManClear
             * @type {Event}                              消息事件
             * @property {String}                  type   清空的交互控件类型
             * @property {cx2.core.control.ActMan} actMan 控件管理对象
             * @example
             * window.cx.events.listen(this, 'actManClear', (evt)=>{}, actManInstance)
             */
            window.vm.$bus.$emit(
              'actManClear',
              { type: type, actMan: this },
              this
            )
          }
          break
        }
      }
    }
  }
}

export default InteractionMan
