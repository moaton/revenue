/*
 =========================================================
 * Vue Black Dashboard - v1.1.3
 =========================================================

 * Product Page: https://www.creative-tim.com/product/black-dashboard
 * Copyright 2024 Creative Tim (http://www.creative-tim.com)

 =========================================================

 * The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

 */
import Vue from "vue";
import VueRouter from "vue-router";
import RouterPrefetch from "vue-router-prefetch";
import App from "./App";
// TIP: change to import router from "./router/starterRouter"; to start with a clean layout
import router from "./router/index";
import store from './store'
import axios from 'axios'
import VueAxios from 'vue-axios'
import VueCookie from 'vue-cookie'

import { BootstrapVue, IconsPlugin } from 'bootstrap-vue'

// import 'bootstrap/dist/css/bootstrap.css'
// import 'bootstrap-vue/dist/bootstrap-vue.css'

import BlackDashboard from "./plugins/blackDashboard";
import i18n from "./i18n";
import "./registerServiceWorker";


Vue.config.productionTip = false
axios.defaults.baseURL = 'http://localhost:8081/api/';


Vue.use(BootstrapVue)
Vue.use(IconsPlugin)
Vue.use(BlackDashboard);
Vue.use(VueRouter);
Vue.use(RouterPrefetch);
Vue.use(VueAxios, axios);
new Vue({
  store,
  router,
  i18n,
  async created() {
    let token = VueCookie.get('token')
    if(token){
      axios.interceptors.request.use(request => {
        request.headers['Authorization'] = `Bearer ${token}`
        return request
      })
    }
  },
  render: (h) => h(App),
}).$mount("#app");
