import Element from 'element-ui'
import '@/styles/element-variables.scss'

export default {
  install(Vue) {
    Element.Dialog.props.closeOnClickModal.default = false
    Element.MessageBox.setDefaults({
      closeOnClickModal: false
    })

    Vue.use(Element, { size: 'medium', zIndex: 2000 })
  }
}
