import Vue from 'vue'
import App from './App.vue'
import vuetify from './plugins/vuetify'
import router from './router'
import store from './store/store'
import 'material-design-icons-iconfont/dist/material-design-icons.css'
import VueAxios from './plugins/axios'

Vue.use(VueAxios)

Vue.config.productionTip = false

new Vue({
  vuetify,
  router,
  store,
  render: h => h(App)
}).$mount('#app')

