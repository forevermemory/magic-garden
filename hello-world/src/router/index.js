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
        component: () => import('@/views/index/Index.vue')
    },
    {
        path: '/qqkongjian',
        name: 'QQkongjian',
        component: () => import('@/views/QQkongjian.vue')
    },
    {
        path: '/userinfo',
        name: 'Userinfo',
        component: () => import('@/views/Userinfo.vue')
    },
    /////////////// 二级菜单
    {
        path: '/index-friends',
        name: 'IndexFriendsManager',
        children:[
            {
                path: 'index-friends-good',
                name: 'IndexFriends',
                component: () => import('@/views/index/IndexFriends.vue')
            },
            {
                path: 'index-friends-black',
                name: 'IndexFriendsBlack',
                component: () => import('@/views/index/IndexFriendsBlack.vue')
            },
        ],
        component: () => import('@/views/index/IndexFriendsManager.vue')
    },
    {
        path: '/index-family',
        name: 'IndexFamily',
        component: () => import('@/views/index/IndexFamily.vue')
    },
    {
        path: '/index-square',
        name: 'IndexSquare',
        component: () => import('@/views/index/IndexSquare.vue')
    },
    {
        path: '/index-games',
        name: 'IndexGames',
        component: () => import('@/views/index/IndexGames.vue')
    },

  //////////////////////////////////////////////
  // gameManager
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
  ///////////////////////////////////////
    {
        path: '/game',
        name: 'Game',
        children:[
        ],
        component: () => import('@/views/garden/EnterGame.vue')
    },
    {
        path: '/game/garden',
        name: 'GardenManager',
        children:[
            {
                path: '/',
                name: 'GardenMyIndex',
                component: () => import('@/views/garden/GardenMyIndex.vue')
            },
            {
                path: 'index',
                name: 'GardenMyIndex2',
                component: () => import('@/views/garden/GardenMyIndex.vue')
            },
            {
                path: 'friend',
                name: 'GardenMyFriend',
                component: () => import('@/views/garden/GardenMyFriend.vue')
            },
            {
                path: 'flowers',
                name: 'GardenFlowers',
                component: () => import('@/views/garden/GardenFlowers.vue')
            },
            {
                path: 'flowers-send',
                name: 'GardenFlowersSend',
                component: () => import('@/views/garden/GardenFlowersSend.vue')
            },
            {
                path: 'flowers-send-record',
                name: 'GardenFlowersSendRecord',
                component: () => import('@/views/garden/GardenFlowersSendRecord.vue')
            },
            {
                path: 'flowers-receive-record',
                name: 'GardenFlowersReceiveRecord',
                component: () => import('@/views/garden/GardenFlowersReceiveRecord.vue')
            },
            {
                path: 'flowers-vase',
                name: 'GardenFlowersVase',
                component: () => import('@/views/garden/GardenFlowersVase.vue')
            },
            {
                path: 'magic-house',
                name: 'GardenMagicHouse',
                component: () => import('@/views/garden/GardenMagicHouse.vue')
            },
            // tools /////////////////////////////////////////////////////////
            // 背包 
            {
                path: 'backpack',
                name: 'GardenToolsBackpack',
                component: () => import('@/views/garden/tools/GardenToolsBackpack.vue')
            },
            {
                path: 'backpack-props',
                name: 'GardenToolsBackpackProps',
                component: () => import('@/views/garden/tools/GardenToolsBackpackProps.vue')
            },
            // 商店
            {
                path: 'shop',
                name: 'GardenToolsShop',
                component: () => import('@/views/garden/tools/GardenToolsShop.vue')
            },
            // 道具
            {
                path: 'props',
                name: 'GardenToolsProps',
                component: () => import('@/views/garden/tools/GardenToolsProps.vue')
            },
            // 签到
            {
                path: 'signin',
                name: 'GardenToolsSignin',
                component: () => import('@/views/garden/tools/GardenToolsSignin.vue')
            },
            // 帮助 
            {
                path: 'help',
                name: 'GardenToolsHelp',
                component: () => import('@/views/garden/tools/GardenToolsHelp.vue')
            },
            {
                path: 'help-detail',
                name: 'GardenToolsHelpDetail',
                component: () => import('@/views/garden/tools/GardenToolsHelpDetail.vue')
            },
            // 神秘商店 
            {
                path: 'shop-sterious',
                name: 'GardenToolsShopSterious',
                component: () => import('@/views/garden/tools/GardenToolsShopSterious.vue')
            },
            // 花种交易市场
            {
                path: 'market',
                name: 'GardenToolsMarket',
                component: () => import('@/views/garden/tools/GardenToolsMarket.vue')
            },
            // 花种交易市场-出售闲置花种
            {
                path: 'market-sale',
                name: 'GardenToolsMarketSale1',
                component: () => import('@/views/garden/tools/GardenToolsMarketSale1.vue')
            },
            // 花种交易市场-出售闲置花种-2
            {
                path: 'market-sale2',
                name: 'GardenToolsMarketSale2',
                component: () => import('@/views/garden/tools/GardenToolsMarketSale2.vue')
            },
            // 花种交易市场-出售闲置花种-3
            {
                path: 'market-sale3',
                name: 'GardenToolsMarketSale3',
                component: () => import('@/views/garden/tools/GardenToolsMarketSale3.vue')
            },
            // 花种交易市场-我的橱窗
            {
                path: 'market-onsale',
                name: 'GardenToolsMarketOnSale',
                component: () => import('@/views/garden/tools/GardenToolsMarketOnSale.vue')
            },
            // 花种交易市场-我的橱窗-下架
            {
                path: 'market-onsale2',
                name: 'GardenToolsMarketOnSale2',
                component: () => import('@/views/garden/tools/GardenToolsMarketOnSale2.vue')
            },

            //////////////////////////////////////////////////////////////
            // 图谱
            {
                path: 'atlas-common',
                name: 'GardenAtlasCommon',
                component: () => import('@/views/garden/atlas/GardenAtlasCommon.vue')
            },
            {
                path: 'atlas-unique',
                name: 'GardenAtlasUnique',
                component: () => import('@/views/garden/atlas/GardenAtlasUnique.vue')
            },
            {
                path: 'atlas-rare',
                name: 'GardenAtlasRare',
                component: () => import('@/views/garden/atlas/GardenAtlasRare.vue')
            },

        ],
        component: () => import('@/views/garden/GardenManager.vue')
    },

]

const router = new VueRouter({
  routes
})

export default router