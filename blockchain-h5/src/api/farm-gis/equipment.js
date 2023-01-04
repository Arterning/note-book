/*
 * @Author: your name
 * @Date: 2021-09-30 15:52:05
 * @LastEditTime: 2021-09-30 17:00:02
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: \agriculture-app\src\api\farm-gis\equipment.js
 */
import request from '@/utils/request'

// 查询瑞仪卡监控设备信息
export function getRykEquipmentData(payload) {
  return request({
    url: '/agriculture/v1/deviceMessage/getLastMessage',
    method: 'get',
    data: payload
  })
}
