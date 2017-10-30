import Vue from 'vue'
import VueLogger from 'vuejs-logger'
import VueRouter from 'vue-router'
import VueResource from 'vue-resource'
import App from './components/App.vue'
import PageHeader from './components/PageHeader.vue'
import store from './store'
import config from '../config.json'

import { routes } from './routes'

import './assets/app.scss'

Vue.use(VueRouter)
Vue.use(VueResource)
Vue.use(VueLogger, {
    logLevel : config.logLevel, 
    stringifyArguments : false,
    showLogLevel : false,
    showMethodName : false,
    separator: '|',
    showConsoleColors: true
})

Vue.http.options.root = config.apiBaseUrl

const router = new VueRouter({
  mode: 'history',
  routes
})

Vue.component('page-header', PageHeader)

new Vue({ // eslint-disable-line no-new
  el: '#app',
  router,
  store,
  render: h => h(App)
})
