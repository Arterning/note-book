// 2-ol
import InteractionMan from './ol/control/InteractionMan'

import ExtentFactory from './ol/factory/ExtentFactory'
import FeatureFactory from './ol/factory/FeatureFactory'
import InteractionFactory from './ol/factory/InteractionFactory'
import LyrTileFactory from './ol/factory/LyrTileFactory'
import LyrVectorFactory from './ol/factory/LyrVectorFactory'
import MapFactory from './ol/factory/MapFactory'
import OverlayFactory from './ol/factory/OverlayFactory'
import StyleFactory from './ol/factory/StyleFactory'

import HighLight from './ol/util/HighLight'
import SelFeaByWFS from './ol/util/SelFeaByWFS'
import SelFeaByWMS from './ol/util/SelFeaByWMS'
// import TileClip from './ol/util/TileClip'
// import QuickSort from './ol/util/QuickSort'
import cx from '../cx/index'

window.cx = cx
window.cx.colorManage = new cx.ColorRangeMaker({
  color: [
    [255, 0, 0],
    [251, 128, 0],
    [246, 255, 0],
    [123, 255, 72],
    [0, 255, 144]
  ],
  value: [0.8, 1.2]
})
const cxgis = {
  ol: {
    control: {
      InteractionMan
    },
    factory: {
      ExtentFactory,
      FeatureFactory,
      InteractionFactory,
      LyrTileFactory,
      LyrVectorFactory,
      MapFactory,
      OverlayFactory,
      StyleFactory
    },
    util: {
      HighLight,
      SelFeaByWFS,
      SelFeaByWMS
      // TileClip,
      // QuickSort
    }
  }
}
window.cxgis = cxgis

export default cxgis
