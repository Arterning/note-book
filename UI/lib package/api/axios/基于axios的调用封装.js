/**
 * axios配置好后就可以export
 * 然后我们引入为request
 * 在很多项目中可以看到这样定义API的形式 感觉这种风格还是不错的 很清晰 值得学习
 */
import request from '@/utils/request'

export function login(username, password) {
  return request({
    url: '/user/login',
    method: 'post',
    data: {
      username,
      password
    }
  })
}

export function getInfo(token) {
  return request({
    url: '/user/info',
    method: 'get',
    params: { token }
  })
}