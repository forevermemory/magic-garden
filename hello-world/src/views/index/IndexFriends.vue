<template>
  <div class="index">
        {{error}}
        


[<a href="/home/label_config/01.html">-</a>]我的好友(2)<br/>
<img src="/image/site/home.gif" alt="."/><a href="/home/friend/35792031.html">111</a>[<a href="/home/chat/35792031.html">发家信</a>]<br/>
<img src="/image/site/home.gif" alt="."/><a href="/home/friend/35798849.html">倪永炜</a>[<a href="/home/chat/35798849.html">发家信</a>]<br/>
<div class="module-title">功能导航</div>
<div class="module-content">
<a href="/home/chat_list.html">全部消息</a>|<a href="/home/message_list.html">系统消息</a>|<a href="/home/friend_news.html">好友动态</a><br/>
<a href="/home/apply_config.html">交友设置</a>|<a href="/home/friend_find.html">好友查找</a>|<a href="/home/label_list.html">分组管理</a><br/>
</div>

    </div>
</template>

<script>
// import {login,listGames} from '@/api/request.js'
// import IndexHeader from '@/components/headers/IndexHeader.vue'

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
        // IndexHeader,
        // Header
    },
    created(){
        // let userinfo = this.$userinfo.get()
        // if(userinfo){
        //     this.userinfo = userinfo
        //     this.username = this.userinfo.username
        // }
    },
    mounted(){
      
        
        // 查询游戏列表
        // this.getGamesList()
    },
    methods: {
       
        
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