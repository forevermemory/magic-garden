<template>
  <div class="index">
        {{error}}
        
        <IndexHeader />

                
        <a href="/home/profile/35798847.html">{{userinfo.username}}</a> 
        {{userinfo.LevelObject.level_name}}
        <img :src='`/static${userinfo.LevelObject.level_img}`' alt="."/>
        (
            <router-link :to="{name: 'Userinfo'}">详情 </router-link>
        )
        <br/>

        <img src="~@/assets/images/site/002.gif" alt="."/>
        连续登录，今日+2活跃天.
        <a href="/bbs/topic/1257511.html">加速升级>>></a>
        <br/>

        <div class="write-mood">
            暂无心情. 
            <a href="/home/my_mood.html">&gt;&gt;</a>
            <br/>
        </div>

        ☆我的恋人:<a href="/bbs/marriage/room/">进入洞房</a> 
        ☆超Q 特权:<a href="/home/xy.html">每日许愿</a>
        <br/>

        ☆今日必作:
        <a href="/home/7t.html">签到好礼</a>☆
        <a href="/home/promo.html">推荐豪礼</a>☆
        <a href="/bbs/my_family.html">快速回家</a>☆
        
        <div class="module-title">
            【我的游戏】 <a href="/bbs/topic_active.html">活动</a> |
            <a href="/home/my_topic.html">帖</a>  |
            <a href="/book/user/">书</a>
            <br/>
        </div>

        <div class="module-content">
            正在玩:
            <router-link :to="{path:'/game/garden/index'}">魔法花园</router-link>
            <br/>

            <div v-for="game in games" :key="game._id">
                <img src="~@/assets/images/site/new.gif" alt="."/>
                <!-- <router-link :to="game.g_url"> -->
                <router-link :to="{ path: '/game', query: { game_id: game._id ,g_url:game.g_url}}">
                    {{game.g_name}}
                </router-link>
                <br/>
            </div>
           
        </div>
        <div class="module-content">
            <router-link to="/gameManager">
                游戏管理
            </router-link>
            <!-- <router-link to="/gameManager/updateGame">
                管理游戏
            </router-link> -->
        </div>

        
        <div class="module-title">
            【<a href="/home/my_news.html">我的新鲜事</a>】
            <br/>
        </div>
        <br/>

        <div class="module-title">
            【<a href="/home/friend_news.html">好友新鲜事</a>】
            <br/>
        </div>
        <div class="events">
            1.(2天前)<a href="/home/profile/35792031.html"><font color="#ff0000">　晴天哥哥　</font></a>发表帖子：
            <a href="/bbs/topic/1298042.html">墨明</a><br/>
            2.(2天前)<a href="/home/profile/35792031.html"><font color="#ff0000">　晴天哥哥　</font></a>回复帖子：
            <a href="/bbs/topic/1298033.html">这位 云涧 離歌</a><br/>
            3.(3天前)<a href="/home/profile/35792031.html"><font color="#ff0000">　晴天哥哥　</font></a>发表帖子：
            <a href="/bbs/topic/1297957.html">早上好</a><br/>
        </div>

        <div class="module-title">
            【<a href="/home/my_visitor.html">来访客人(共1次)</a>】<br/>
        </div>
        <br/>

        <div class="module-title">
            【<a href="/home/my_message.html">家园留言(共0条)</a>】
            <br/>
        </div>

        <br/>
    </div>
</template>

<script>
import {login,listGames} from '@/api/request.js'
import IndexHeader from '@/components/headers/IndexHeader.vue'

export default {
    data() {
        return {
            userinfo:{},
            username:'未知用户',
            error:'',
            page:{
                pageSize:100000,
                pageNo:1,
            },
            games:[],
        }
    },
    components: {
        IndexHeader,
        // Header
    },
    created(){
        let userinfo = this.$userinfo.get()
        if(userinfo){
            this.userinfo = userinfo
            this.username = this.userinfo.username
        }
    },
    mounted(){
        // VueEvent.$emit('toHeader',"【家园首页】")
        // 从vuex缓存取数据
        // if(this.$store.state.UserInfo){
        //     // 从注册进来的
        //     this.formValue.username = this.$store.state.UserInfo.username
        //     this.formValue.password = this.$store.state.UserInfo.password
        // }

        
        // 查询游戏列表
        this.getGamesList()
    },
    methods: {
        async loginBtn(){
            // 校验
            if (this.validateRegisterInfo()!=null){
                this.error = this.validateRegisterInfo()
                return
            }
            this.error = ''
            // 调用注册
            let res = await login(this.formValue)
            if(res.code==0){
                this.error = '登陆成功'
                // 跳转index

            }
        },

        async getGamesList() {
            let res = await listGames(this.page)
            this.games = res.data.data
        }
        
    }
}
</script>

<style >
/* // <style lang="scss"> */

.module-title {
    margin: 8px 0 0;
    padding: 0 5px;
    height: 20px;
    line-height: 20px;
    border-bottom: 4px solid #9FC6EC;
    color: #000;
    font-weight: bold;
}



</style>