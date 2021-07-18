

<template>
  <div class="index">
        <router-link :to="{name: 'Userinfo'}">
            <img src="~@/assets/images/site/id.gif" alt="."/>{{this.username}}
        </router-link>
        <router-link :to="{name: 'Index'}">
            <img src="~@/assets/images/site/home.gif" alt="."/>家园
        </router-link>
        <router-link :to="{name: 'QQkongjian'}">
            <img src="~@/assets/images/site/blog.gif" alt="."/>空间
        </router-link>
        <br/>
        <!-- <div class="note">
        <img src="/image/site/001.gif" alt="."/><a href="/ad/ad/1.html">[幻想西游]古典神话网游！</a><br/>
        </div> -->

    
  </div>
</template>

<script>
export default {
    name: 'IndexBaseHeader',
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



</style>
