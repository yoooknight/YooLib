import request from '@/utils/request'

export function list(token) {
  console.log('11111111111')
  return request({
    url: '/permission/list',
    method: 'get'
  })
}
