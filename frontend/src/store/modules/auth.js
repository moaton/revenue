import axios from 'axios'
import VueCookie from 'vue-cookie'


export default {
  namespaced: true,
  state: {
    user: {},
    token: null,
    isLoading: false,
  },
  actions: {
    async login({ commit }, payload){
      try {
        commit('SET_IS_LOADING', true)
        const r = await axios.post('users', payload)
        console.log(r);
        commit('SET_TOKEN', r.data.payload.token)
        return true
      } catch (error) {
        console.log(error);
        return false
      } finally {
        commit('SET_IS_LOADING', false)
      }
    }
  },
  mutations: {
    SET_IS_LOADING(state, value){
      state.isLoading = value
    },
    SET_TOKEN(state, token){
      state.token = token
      VueCookie.set('token', token, '1m', null, 'localhost')
      axios.interceptors.request.use(request => {
        request.headers['Authorization'] = `Bearer ${token}`
        return request
      })
    },
    SET_USER(state, user){
      state.user = user
    }
  },
}