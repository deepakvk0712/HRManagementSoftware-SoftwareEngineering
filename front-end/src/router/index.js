import Vue from 'vue'
import VueRouter from 'vue-router'
import LandingPage from '../components/LandingPage.vue'
import AttendancePage from '../components/AttendancePage.vue'
// import PayslipPage from '../components/PayslipPage.vue'
import UserProfileView from '../views/UserProfileView.vue'
import AboutView from '../views/AboutView.vue'
import LoginPage from '../components/LoginPage.vue'
import SignupPage from '../components/SignupPage.vue'
<<<<<<< Updated upstream
// import store from '../store/store'

// (function() {
//   if (store.state.accessToken) {
//       this.$axios.defaults.headers.common['Authorization'] = "Bearer " + this.$store.state.accessToken;
//   } else {
//       // this.$axios.defaults.headers.common['Authorization'] = null;
//   }
// })();
=======
import ChangePassword from '../components/ChangePassword.vue'
>>>>>>> Stashed changes

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    redirect: '/login'
  },
  {
    path: '/landing',
    name: 'landingPage',
    component: LandingPage
  },
  {
    path: '/attendance',
    name: 'attendance',
    component: AttendancePage
  },
  // {
  //   path: '/payslip',
  //   name: 'payslip',
  //   component: PayslipPage
  // },
  {
    path: '/about',
    name: 'about',
    component: AboutView
  },
  {
    path: '/user-profile',
    name: 'profile',
    component: UserProfileView
  },
  {
    path: '/login',
    name: 'login',
    component: LoginPage
  },
  {
    path: '/signup',
    name: 'signup',
    component: SignupPage
  },
  {
    path: '/ChangePassword',
    name: 'ChangePassword',
    component: ChangePassword
  },
]

const router = new VueRouter({
  routes
})

export default router
