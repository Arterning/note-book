import request from '@/utils/request'

// --------------------------------------- 投诉商品 ---------------------------------------

export function getList(data) {
  return request({
    url: '/blockchainadmin/v2/complaintDetails/list',
    method: 'get',
    data
  })
}

export function submitInfo(data) {
  return request({
    url: '/blockchainadmin/v2/complaintDetails/handle',
    method: 'post',
    data,
    headers: {
      'Content-Type': 'application/json'
    }
  })
}
