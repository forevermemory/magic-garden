import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'


// VUEX
import Vuex from 'vuex'
Vue.use(Vuex)


import VueRouter from 'vue-router'
import ElementUI from 'element-ui';
import 'element-ui/lib/theme-chalk/index.css';

Vue.use(VueRouter)
Vue.use(ElementUI);

//////////////////////
var $userinfo = {}
$userinfo.get = function () { 
    var info = localStorage.getItem("userinfo")
    if (!info){
        return null
    }
    return JSON.parse(info)
}
$userinfo.set = function (info) { 
    localStorage.setItem("userinfo",JSON.stringify(info))
}

Vue.prototype.$userinfo = $userinfo;

new Vue({
  router,
  store,
  render: h => h(App)
}).$mount('#app')
