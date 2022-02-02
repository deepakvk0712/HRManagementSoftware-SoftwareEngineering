import { createRouter, createWebHashHistory } from 'vue-router'
import HomeTab from '../NavigationTabs/HomeTab.vue'
import PaySlipTab from '../NavigationTabs/PaySlipTab.vue'
import LeaveTab from '../NavigationTabs/LeaveTab.vue'
import MyInfoTab from '../NavigationTabs/MyInfoTab.vue'

const routes = [{
    path: '/',
    name: 'home',
    component: HomeTab
  },
  {
    path: '/myinfo',
    name: 'myinfo',
    component: MyInfoTab
  },
  {
    path: '/leave',
    name: 'leaveinfo',
    component: LeaveTab
  },
  {
    path: '/payslip',
    name: 'payslip',
    component: PaySlipTab
  }
]

const router = createRouter({
  history: createWebHashHistory(),
  routes
})

export default router
