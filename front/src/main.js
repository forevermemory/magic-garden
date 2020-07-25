import Vue from "vue";
import App from "./App.vue";
import router from "./router";

Vue.config.productionTip = false;

new Vue({
    router,
    render: function(h) {
        return h(App);
    },
}).$mount("#app");

// 全局引入基础scss  main.js
// 	import './assets/css/basic.scss'
