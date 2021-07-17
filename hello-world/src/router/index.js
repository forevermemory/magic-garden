// import Vue from 'vue'
import VueRouter from 'vue-router'
// import ElementUI from 'element-ui';
// import 'element-ui/lib/theme-chalk/index.css';

// Vue.use(VueRouter)
// Vue.use(ElementUI);

const originalPush = VueRouter.prototype.push
VueRouter.prototype.push = function push(location) {
  return originalPush.call(this, location).catch(err => err)
}
const routes = [
  {
    path: '',
    name: 'login2',
    component: () => import('@/views/Login.vue'),
  },
  {
    path: '/login',
    name: 'loginlogin1',
    component: () => import('@/views/Login.vue'),
  },
  {
    path: '/register',
    name: 'Register',
    component: () => import('@/views/Register.vue')
  },
  {
    path: '/index',
    name: 'Index',
    component: () => import('@/views/Index.vue')
  },
  
  // games
  {
    path: '/gameManager',
    name: 'GameManager',
    children:[
        {
            path: 'addGame',
            name: 'AddGame',
            component: () => import('@/views/games/AddGame.vue')
        },
        {
            path: 'updateGame',
            name: 'UpdateGame',
            component: () => import('@/views/games/UpdateGame.vue')
        },
        {
            path: 'addGameResult',
            name: 'AddGameResult',
            component: () => import('@/views/games/AddGameResult.vue')
        },
    ],
    component: () => import('@/views/games/GameManager.vue')
  },
//   {
//     path: '/gameManager/addGame',
//     name: 'AddGame',
//     component: () => import('@/views/games/AddGame.vue')
//   },
//   {
//     path: '/gameManager/updateGame',
//     name: 'AddGame',
//     component: () => import('@/views/games/UpdateGame.vue')
//   },

]

const router = new VueRouter({
  routes
})

export default router