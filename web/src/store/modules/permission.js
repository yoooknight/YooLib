import { asyncRoutes, constantRoutes } from '@/router'
import Layout from '@/layout/index'

/**
 * Use meta.role to determine if the current user has permission
 * @param roles
 * @param route
 */
// function hasPermission(roles, route) {
//   if (route.meta && route.meta.roles) {
//     return roles.some(role => route.meta.roles.includes(role))
//   } else {
//     return true
//   }
// }

/**
 * Filter asynchronous routing tables by recursion
 * @param routes asyncRoutes
 * @param roles
 */
export function filterAsyncRoutes(routes, menu) {
  const res = []

  // routes.forEach(route => {
  //   const tmp = { ...route }
  //   if (hasPermission(roles, tmp)) {
  //     if (tmp.children) {
  //       tmp.children = filterAsyncRoutes(tmp.children, roles)
  //     }
  //     res.push(tmp)
  //   }
  // })
  menu.forEach(smenu => {
    const tmp = {
      path: smenu.web_router,
      component: Layout,
      name: smenu.name,
      redirect: 'noRedirect',
      meta: { title: smenu.name },
      children: smenu.children
    }

    console.log(tmp)
    console.log(smenu)

    if (tmp.children) {
      console.log('22222222222222')
      tmp.children = filterAsyncRoutes(routes, tmp.children)
    } else {
      console.log('333333333333')
      tmp.component = (resolve) => require(['@/views' + smenu.web_router], resolve)
      delete tmp.redirect
      delete tmp.children
    }
    res.push(tmp)
  })

  return res
}

const state = {
  routes: [],
  addRoutes: []
}

const mutations = {
  SET_ROUTES: (state, routes) => {
    state.addRoutes = routes
    state.routes = constantRoutes.concat(routes)
  }
}

const actions = {
  generateRoutes({ commit }, menu) {
    console.log(menu)
    return new Promise(resolve => {
      let accessedRoutes
      if (menu.includes('root')) {
        accessedRoutes = asyncRoutes || []
      } else {
        accessedRoutes = filterAsyncRoutes(asyncRoutes, menu)
      }

      console.log(accessedRoutes)
      commit('SET_ROUTES', accessedRoutes)
      resolve(accessedRoutes)
      // list().then(res => {
      //   // console.log(res)
      //   // 将获取到的数据转化
      //   const accessedRouters = arrayToMenu(res.data)
      //   accessedRouters.concat([{ path: '*', redirect: '/404', hidden: true }])
      //   commit('SET_ROUTERS', accessedRouters)
      //   resolve
      // })
    })
  }
}

// export function arrayToMenu(array) {
//   console.log(array)
// }

export default {
  namespaced: true,
  state,
  mutations,
  actions
}
