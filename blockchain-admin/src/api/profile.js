import request from '@/utils/request'
/**
 * @description: 获取用户信息
 * @author: Hemingway
 */
export function getPersonalCenter(data) {
  return request({
    url: '/blockchainadmin/user/v1/getPersonalCenter',
    method: 'post',
    data
  })
}
