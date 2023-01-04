/** 公共方法*/
import axios from 'axios'

window.svrCache = window.svrCache || {}

/**
 * 向数据库批量添加图形
 * @static
 * @method
 * @param {Number}             major           图形表主类型
 * @param {Number}             minor           图形表子类型
 * @param {Array.<Object>}     graDatas        图形表数据对象序列
 * @returns {Promise}
 */
const addGras = (major, minor, graDatas) => {
  const params = {
    _major: major,
    _minor: minor,
    json: JSON.stringify(graDatas)
  }
  return query({ ...params, cmd: 'addGras' })
}

const getFlds = (tabname, cacheKey) => {
  const param = {
    _major: 99,
    _minor: 5,
    _exp: `tabname='${tabname}'`
  }
  return query(param, cacheKey)
}

const makeParam = (param, server = 'AreaSvr') => {
  if (param && param.method === 'post') {
    return {
      method: 'post',
      url: `${process.env.VUE_APP_GIS_BASE_URL}/gis/${param.server || server}`,
      data: { cmd: 'query', ...param },
      headers: {
        'content-type': 'application/json'
      },
      ...(param.axiosCfg || {})
    }
  } else {
    return {
      method: 'get',
      url: `${process.env.VUE_APP_GIS_BASE_URL}/gis/${server}`,
      params: { cmd: 'query', ...param, axiosCfg: '' },
      ...(param.axiosCfg || {})
    }
  }
}

/**
 * 基础的查询服务
 * @method
 * @param   {Object}    param         查询参数对象
 * @param   {String}    [cacheKey]    缓存键，如果传入的话会进行缓存，下次查询直接从缓存中返回
 */
const query = (param, cacheKey) => {
  return new Promise((resolve, reject) => {
    if (cacheKey && window.svrCache[cacheKey]) {
      return resolve(window.svrCache[cacheKey])
    }
    axios(makeParam(param, param.server))
      .then(ret => {
        if (cacheKey) {
          window.svrCache[cacheKey] = ret.data
        }
        resolve(ret.data)
      })
      .catch(reject)
  })
}

/**
 * 根据条件查询空间实体数据,主表包含geom字段
 * @static
 * @method
 * @param   {Object}    params                    参数
 * @param   {Number}    params._major             实体主类型
 * @param   {Number}    params._minor             实体子类型
 * @param   {String}    params.tables             联合查询的表名(例'a;b:5,7' 1.'a'代表主表别名(只有主表不需要指定主子类型) 2.'b:5,7'表别名b,对应的主类型5,子类型7)
 * @param   {String}    params._exp               SQL条件表达式(1.列名前需要带表别名;2.参数使用?代替))
 * @param   {String}    params.cols               需要取的列(例'a:*;b:fldb1,fldb2;c:fldc1 as TTT' 1.'a.*'代表a表的全部列;2.'b:fldb1,fldb2'代表b表的fldb1,fldb2字段;3.'c:fldc1 as TTT'代表c表的fldc1字段, 返回时用名字TTT)
 * @param   {String}    params.types              参数类型(用,分隔)(1/i/long-整数,2/r/double-实数,3/s/string-字符串,4/d/date-日期,5/t/datetime-日期时间)
 * @param   {String}    params.vals               参数值(用,分隔)(与参数一一对应)
 * @param   {String}    [params.orderby]          排序（可以不传）
 * @param   {String}    [params.separator = ',']  数值中的分隔符（可以不传,默认是,）(考虑到字符串中可能有,可以设置特别分隔符)
 * @param   {String}    [params.init]             0/1  首次调用为0, 非第一次为1
 * @param   {String}    [params.pageno]           页码(1开始)
 * @param   {String}    [params.pagesize]         每页数据条数
 * @param   {Number}    [params.geometry=1]       是否返回图形数据，1-返回，0-不返回
 * @param   {String}    [params.xy]               空间查询的范围坐标串，xy有值时开启空间查询模式，gtype必须要传
 * @param   {Number}    [params.gtype]            空间范围的几何对象类型，3-矩形，4-多边形，5-圆心半径圆，1-缓冲区圆
 * @param   {Number}    [params.gr=0]             圆形空间查询时的半径，gtype是圆时必须传
 * @param   {Number}    [params.r=0]              缓冲区圆半径，gtype是圆时必须传
 * @param   {String}    [cacheKey]                缓存键
 * @returns {Promise}
 */
const queryUnionGra = (params, cacheKey) => {
  params.geometry = params.geometry || 1
  params.gr = params.gr || 0
  params.r = params.r || 0

  return query({ cmd: 'queryUnionGra', ...params }, cacheKey)
}

/**
 * 查询usersql表中的sql并执行返回结果
 * @method
 * @param {Object}  params        参数对象
 * @param {String}  params.keyno         usersql表中的唯一标识符
 * @param {String}  [params.vals]        值，多个值中间用分号;隔开
 * @param {String}  [params.exp]         自定义sql语句串，多个中间用分号;隔开
 * @param {String}  [cacheKey]    缓存键，如果传入的话会进行缓存，下次查询直接从缓存中返回
 */
const querySql = (params, cacheKey) => {
  return query({ ...params, cmd: 'querySQL', _major: 99, _minor: 47 }, cacheKey)
}

/**
 * 批量更新服务。
 * @method
 * @param {Object}  params          参数对象
 * @param {Number}  params._major    待更新的实体主类型
 * @param {Number}  params._minor    待更新的实体子类型
 * @param {String}  params.ids      待更新的数据id串，中间用逗号隔开
 * ...其它参数为待更新的列名和对应的值串，如：col1:'val1,val2,val3'
 */
const updateAtts = params => {
  return query({ ...params, cmd: 'updateAtts' })
}

/**
 * 新增数据（单条）
 * @method
 * @param {Number} _major 实体主类型
 * @param {Number} _minor 实体子类型
 * @param {Object} att    待插入的数据对象
 */
const add = (_major, _minor, att) => {
  return query({ cmd: 'add', _major, _minor, retAtt: 0, ...att })
}

/**
 * 新增数据（批量）
 * @method
 * @param {Number}        _major  实体主类型
 * @param {Number}        _minor  实体子类型
 * @param {Array<Object>} atts    待插入的数据对象集合
 */
const adds = (_major, _minor, atts) => {
  return query({ cmd: 'adds', _major, _minor, json: JSON.stringify(atts) })
}

/**
 * 设置计时器，来监测条件满足时执行回调
 * @static
 * @method
 * @param   {Function}    isReady     是否满足条件的判断函数
 * @param   {Function}    callBack    满足条件后的回调函数
 * @param   {Number}      [time=600] 间隔时间(ms)。
 */
const watch = function(isReady, callBack, time = 200) {
  let func = () => {
    if (isReady()) {
      if (myInterval) window.clearInterval(myInterval)
      callBack()
    }
  }
  //设置定时任务，每1000ms执行一次watch方法。
  let myInterval = window.setInterval(func, time)
}

export {
  getFlds,
  query,
  querySql,
  queryUnionGra,
  updateAtts,
  add,
  adds,
  addGras,
  watch
}
