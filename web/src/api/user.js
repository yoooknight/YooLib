import request from '@/utils/request'

export function login(data) {
  return request({
    url: '/user/login',
    method: 'post',
    headers: {
      'Content-Type': 'application/json; charset=UTF-8'
    },
    data
  })
}

export function getInfo(token) {
  console.log(token)
  return request({
    url: '/user/info',
    method: 'get'
  })
}
