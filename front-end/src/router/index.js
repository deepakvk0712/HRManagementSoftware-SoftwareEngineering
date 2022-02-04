import Vue from 'vue'
import VueRouter from 'vue-router'
import LandingPage from '../components/LandingPage.vue'
import LeavePage from '../components/LeavePage.vue'
import PayslipPage from '../components/PayslipPage.vue'

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'landingPage',
    component: LandingPage
  },
  {
    path: '/leaves',
    name: 'leaves',
    component: LeavePage
  },
  {
    path: '/payslip',
    name: 'payslip',
    component: PayslipPage
  }
]

const router = new VueRouter({
  routes
})

export default router
