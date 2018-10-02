import Vue from 'vue'
import Router from 'vue-router'
import About from '@/view/About'
import PageNoFound from '@/view/PageNoFound'

Vue.use(Router)

export default new Router({
  mode: 'history',
  routes: [
    { path: '/', alias: '/about',  name: 'about', component: About },
    { path: '/*', name: 'pageNoFound', component: PageNoFound }
  ]
})
