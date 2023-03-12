import newRequest from '@/utils/newRequest'

// --------------------------------------- 数据字典 ---------------------------------------

export function queryBusinessDic() {
  return newRequest({
    url: '/agriculturepub/v1/dict/query',
    method: 'post'
  })
}
