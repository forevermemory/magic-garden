<template>
    <div>
        {{error}}
        <GardenNavHeader/>
        <div class="flower">

            【
            <router-link :to="{path:'/game/garden/flowers'}">我的花篮</router-link>
            |
            <router-link :to="{path:'/game/garden/flowers-vase'}">我的花瓶</router-link>
            】
            <br/>
        </div>
        
您收到的鲜花会存放在您的花瓶中,共有鲜花1种,1朵.<br/>
<div class="row">
<img src="/data/game/garden/m_s_9.gif" alt="."/><a href="/game/garden/user/bottle/18.html">黄梅花</a> 1<br/>
</div>
<div class="row">
</div>
(第<b>1</b>/1页/共1条记录)<br/>
<br/>


<div class="send-receive">
    <router-link :to="{name:'GardenFlowersSend'}" >送花.</router-link>
    <router-link :to="{name:'GardenMagicHouse'}" >合成.</router-link>
    <router-link :to="{name:'GardenFlowersSendRecord'}" >收花记录</router-link>
    <br/>
</div>


----------<br/>
<!-- <div class="bar">
<a href="/game/garden/user/my_garden.html">魔法花园</a>>我的花瓶<br/>
</div> -->

    </div>
</template>

<script>
// import {getGameByID} from '@/api/request.js'
import GardenNavHeader from '@/components/headers/GardenNavHeader.vue'

export default {
    name: 'GardenFlowersVase',
    data() {
        return {
            // games:[],
            error:'',
            game:{},
        }
    },
    components: {
        GardenNavHeader,
    },
    created(){
    },
    mounted(){
        
        // this.$router.push({ name: 'AddGame'})

    },
    methods: {
        
    }
}
</script>

<style >



/* // <style lang="scss"> */

</style>