import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'
import LoginView from '../views/LoginView.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: HomeView,
      children: [
        {
          path: 'ui/interface/ethernet',
          name: 'ethernet',
          component: () => import("../views/InterfaceView.vue"),
          props: {
            ifaceType: 'ethernet'
          }
        }, {
          path: 'ui/vpn/l2tp',
          name: 'l2tp',
          component: () => import("../views/L2TPView.vue"),
        }, {
          path: 'ui/group/ipv4-address-group',
          name: 'ipv4-address-group',
          component: () => import("../views/AddressGroupView.vue"),
        }, {
          path: "ui/group/ipv6-address-group",
          name: "ipv6-address-group",
          component: () => import("../views/Ipv6AddressGroupView.vue"),
        }, {
          path: 'ui/group/ipv4-network-group',
          name: 'ipv4-network-group',
          component: () => import("../views/NetworkGroupView.vue"),
        }, {
          path: 'ui/group/ipv6-network-group',
          name: 'ipv6-network-group',
          component: () => import("../views/Ipv6NetworkGroupView.vue"),
        }, {
          path: 'ui/group/port-group',
          name: 'port-group',
          component: () => import("../views/PortGroupView.vue"),
        }
      ]
    },
    {
      path: '/ui/login',
      name: 'login',
      component: LoginView
    },
    {
      path: '/about',
      name: 'about',
      // route level code-splitting
      // this generates a separate chunk (About.[hash].js) for this route
      // which is lazy-loaded when the route is visited.
      component: () => import('../views/AboutView.vue')
    }
  ]
})

export default router
