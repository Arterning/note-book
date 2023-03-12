import request from '@/http'

export function getReportList(data) {
  return request({
    url: '/blockchainadmin/srvMachPacking/list',
    method: 'post',

    data
  })
}

export function uploadReport(data) {
  return request({
    url: '/blockchainadmin/srvMachPacking/uploadReport',
    method: 'post',

    data
  })
}
