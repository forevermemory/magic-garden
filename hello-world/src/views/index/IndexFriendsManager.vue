<template>
  <div class="index">
        {{error}}
        
        <IndexHeader />

        <div class="module-title">
            <router-link :to="{name: 'IndexFriends'}"> 我的好友 </router-link> |
            <router-link :to="{name: 'IndexFriendsBlack'}"> 黑名列表 </router-link>
        </div>
        <router-view></router-view>
    </div>
</template>

<script>
// import {login,listGames} from '@/api/request.js'
import IndexHeader from '@/components/headers/IndexHeader.vue'

export default {
    data() {
        return {
            error:'',
        }
    },
    components: {
        IndexHeader,
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
        // 去好友列表
        this.$router.push({ name: 'IndexFriends'})
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