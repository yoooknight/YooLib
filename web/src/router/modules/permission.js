/** When your routing table is too long, you can split it into small modules **/

import Layout from '@/layout'

const permissionRouter = {
  path: '/permission',
  component: Layout,
  redirect: '/permission/manager',
  name: 'Permission',
  meta: {
    title: '权限管理'
    // icon: 'table'
  },
  children: [
    {
      path: 'manager',
      component: () => import('@/views/permission/manager'),
      name: 'manager',
      meta: { title: '管理员管理' }
    },
    {
      path: 'role',
      component: () => import('@/views/permission/role'),
      name: 'role',
      meta: { title: '角色管理' }
    },
    {
      path: 'menu',
      component: () => import('@/views/permission/menu'),
      name: 'menu',
      meta: { title: '菜单管理' }
    }
  ]
}
export default permissionRouter
