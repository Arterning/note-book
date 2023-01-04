import request from '@/utils/request'

// --------------------------------------- 数据字典 ---------------------------------------

export function getSystemDic() {
  return request({
    url: '/agriculturedbo/dictionary/getAllDictionaryKeyValue/w/v1',
    method: 'post'
  })
}

export function queryBusinessDic() {
  return request({
    url: '/agriculturedbo/v1/dict/query',
    method: 'post'
  })
}
