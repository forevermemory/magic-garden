<template>
  <div class="index">
        {{error}}
        
        <IndexHeader />

        <h2>个人中心 </h2>

                
        <a href="/home/profile/35798847.html">{{userinfo.username}}</a> 
        {{userinfo.LevelObject.level_name}}
        <img :src='`/static${userinfo.LevelObject.level_img}`' alt="."/>
        (<a href="/home/level/35798847.html">详情</a>)
        <br/>

        <img src="~@/assets/images/site/002.gif" alt="."/>
        连续登录，今日+2活跃天.
        <a href="/bbs/topic/1257511.html">加速升级>>></a>
        <br/>

        <br/>
    </div>
</template>

<script>
import {login,listGames} from '@/api/request.js'
import IndexHeader from '@/components/headers/IndexHeader.vue'

export default {
    name:'Userinfo',
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