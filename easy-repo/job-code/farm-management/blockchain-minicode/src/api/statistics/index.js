import request from '@/http'

export function getEntiretyStatis(data) {
  return request({
    url: '/blockchainadmin/srvH5AccRecReport/entiretyStatis',
    method: 'GET',

    data
  })
}

export function getMultidimenStatis(data) {
  return request({
    url: '/blockchainadmin/srvH5AccRecReport/multidimenStatis',
    method: 'post',

    data
  })
}

export function getPages(data) {
  return request({
    url: '/blockchainadmin/sysDict/list ',
    method: 'post',
    headers: {
      'Content-Type': 'application/json',
      sessionId: uni.getStorageSync('sessionId')
    },
    data
  })
}
