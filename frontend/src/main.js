import Vue from 'vue'
import App from './App'
import router from './router'

Vue.config.productionTip = false

import '../semantic/dist/semantic.css'
window.$ = window.jQuery = require('jquery')

new Vue({
  el: '#app',
  router,
  components: { App },
  template: '<App/>'
})
