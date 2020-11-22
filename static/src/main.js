import Vue from 'vue'
import App from '@/App'
import '@/plugins/axios'
import router from '@/plugins/router'
import vuetify from '@/plugins/vuetify'

Vue.config.productionTip = false

new Vue({
  router,
  vuetify,
  render: h => h(App)
}).$mount('#app')
