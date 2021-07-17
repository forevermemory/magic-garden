

<template>
  <div class="index">
        <IndexBaseHeader/>
        <!-- <br/> -->
        <!-- <div class="note">
        <img src="/image/site/001.gif" alt="."/><a href="/ad/ad/1.html">[幻想西游]古典神话网游！</a><br/>
        </div> -->

    <div class="bar">
        家园 
        <router-link to=""> 好友 </router-link>
        <router-link to=""> 家族 </router-link>
        <router-link to=""> 广场 </router-link>
        <router-link to=""> 游戏 </router-link>
        <br/>
    </div>
  </div>
</template>

<script>

import IndexBaseHeader from '@/components/headers/IndexBaseHeader.vue'


export default {
    name: 'IndexHeader',
    props: {
        // username:String,
    },
    data() {
        return {
            userinfo:{},
            username:'未知用户',
            error:'',
        }
    },
    components: {
        IndexBaseHeader,
        // Header
    },
    created(){
    },
    mounted(){
        // VueEvent.$emit('toHeader',"【家园首页】")
        // 从vuex缓存取数据
        // if(this.$store.state.UserInfo){
        //     // 从注册进来的
        //     this.formValue.username = this.$store.state.UserInfo.username
        //     this.formValue.password = this.$store.state.UserInfo.password
        // }

        if(localStorage.getItem('userinfo')){

            this.userinfo = JSON.parse(localStorage.getItem('userinfo'))
            this.username = this.userinfo.username
        }

    },
    methods: {
        
    }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.bar {
    height: 25px;
    padding: 0 5px;
    background: #71afe3 url(~@/assets/images/bar.gif) repeat-x;
    line-height: 25px;
    color: #fff;
}


</style>
