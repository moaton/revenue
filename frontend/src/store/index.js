import Vue from 'vue'
import Vuex from 'vuex'
import auth from './modules/auth'
import revenue from './modules/revenue'

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
  },
  mutations: {
  },
  actions: {
  },
  modules: {
    auth,
    revenue
  }
})
