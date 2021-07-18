

<template>
  <div class="index">
        <!-- <br/> -->
        <!-- <div class="note">
        <img src="/image/site/001.gif" alt="."/><a href="/ad/ad/1.html">[幻想西游]古典神话网游！</a><br/>
        </div> -->

    <div class="bar">
         
        <router-link :to="{name: 'Index'}"> 家园 </router-link>
        <router-link :to="{name: 'IndexFriends'}"> 好友 </router-link>
        <router-link :to="{name: 'IndexFamily'}"> 家族 </router-link>
        <router-link :to="{name: 'IndexSquare'}"> 广场 </router-link>
        <router-link :to="{name: 'IndexGames'}"> 游戏 </router-link>
        <br/>
    </div>
  </div>
</template>

<script>



export default {
    name: 'IndexBarHeader',
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
    },
    created(){
    },
    mounted(){

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
