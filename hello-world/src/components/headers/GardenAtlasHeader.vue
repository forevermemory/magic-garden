

<template>
  <div class="index">
        <div class="atlas">
            【花种图谱】
            <br/>
            花园世界目前已知花种535种，您成功培育出了21种。
            <br/>
            <router-link :to="{path:'/game/garden/atlas-common'}" >普通(6)</router-link>
            <router-link :to="{path:'/game/garden/atlas-unique'}">|独特(15)|</router-link>
            <router-link :to="{path:'/game/garden/atlas-rare'}">珍稀(0)</router-link>
            <br/>
        </div>
  </div>
</template>

<script>



export default {
    name: 'GardenAtlasHeader',
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
