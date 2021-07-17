import Vue from 'vue';
import Vuex from 'vuex';

Vue.use(Vuex);

export default new Vuex.Store({
  state: {
    UserInfo: null,
    UserInfo2: null,
  },
  mutations: {
    setUserInfo(state, info) {
      state.UserInfo = info;
    },
    setUserInfo2(state, info) {
      state.UserInfo2 = info;
    },
  
  },
  actions: {

  },
});