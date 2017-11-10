import Vue from 'vue'
import App from './App'
import Vuetify from 'vuetify'
import VueParticles from 'vue-particles'
import 'vuetify/dist/vuetify.min.css'
import 'clipboard/dist/clipboard.min'

Vue.use(Vuetify)
Vue.use(VueParticles)

Vue.config.productionTip = false

/* eslint-disable no-new */
new Vue({
  el: '#app',
  template: '<App/>',
  components: { App }
})
