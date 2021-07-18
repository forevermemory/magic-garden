

<template>
  <div class="index">
        <div class="header">
            <router-link :to="{path:'/game/garden/index'}">我的花园|</router-link>
            <router-link :to="{path:'/game/garden/friend'}">好友|</router-link>
            <router-link :to="{path:'/game/garden/flowers'}">花房|</router-link>
            <router-link :to="{path:'/game/garden/magic-house'}">魔法屋</router-link>
            <br/>
        </div>
  </div>
</template>

<script>



export default {
    name: 'GardenNavHeader',
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



</style>
