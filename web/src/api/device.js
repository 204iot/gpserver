import request from '@/utils/request'

export function devices() {
  return request({
    url: '/devices',
    method: 'get'
  })
}
