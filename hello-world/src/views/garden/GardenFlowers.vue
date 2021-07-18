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

您收获和采摘的花均存放于您的花篮中,共有鲜花20种,50朵<br/>
<div class="row">
<img src="/data/game/garden/m_s_203.gif" alt="."/><a href="/game/garden/user/basket/26258.html">橙色圣诞之吻</a> 1<br/>
</div>
<div class="row">
<img src="/data/game/garden/m_s_179.gif" alt="."/><a href="/game/garden/user/basket/26257.html">橙色花样年华</a> 1<br/>
</div>
<div class="row">
<img src="/data/game/garden/m_s_224.gif" alt="."/><a href="/game/garden/user/basket/26256.html">蓝色雪花纷飞</a> 1<br/>
</div>
<div class="row">
<img src="/data/game/garden/m_s_136.gif" alt="."/><a href="/game/garden/user/basket/26105.html">孤芳自赏</a> 1<br/>
</div>
<div class="row">
<img src="/data/game/garden/m_s_191.gif" alt="."/><a href="/game/garden/user/basket/26104.html">紫色奥运之花</a> 1<br/>
</div>
<div class="row">
<img src="/data/game/garden/m_s_313.gif" alt="."/><a href="/game/garden/user/basket/26103.html">武之王者</a> 2<br/>
</div>
<div class="row">
<img src="/data/game/garden/m_s_197.gif" alt="."/><a href="/game/garden/user/basket/15698.html">粉色水之精灵</a> 1<br/>
</div>
<div class="row">
<img src="/data/game/garden/m_s_172.gif" alt="."/><a href="/game/garden/user/basket/15622.html">白色岁月如歌</a> 1<br/>
</div>
<div class="row">
<img src="/data/game/garden/m_s_8.gif" alt="."/><a href="/game/garden/user/basket/15621.html">红梅花</a> 7<br/>
</div>
<div class="row">
<img src="/data/game/garden/m_s_139.gif" alt="."/><a href="/game/garden/user/basket/1790.html">红友谊花</a> 1<br/>
</div>
<div class="row">
<a href="/game/garden/user/basket_list-2.html">下页</a>
<form action="/game/garden/user/basket_list.html" method="get">
第<input type="text" name="page" size="2" emptyok="true" format="*N" value="1" maxlength="10" />页
<input type="submit" value="前往" />
</form>
</div>
(第<b>1</b>/2页/共20条记录)<br/>
<br/>


<div class="send-receive">
    <router-link :to="{name:'GardenFlowersSend'}" >送花.</router-link>
    <router-link :to="{name:'GardenMagicHouse'}" >合成.</router-link>
    <router-link :to="{name:'GardenFlowersSendRecord'}" >送花记录</router-link>
    <br/>
</div>

----------<br/>
<a href="/game/garden/user/my_garden.html">花园</a>>花篮<br/>
        
    </div>
</template>

<script>
// import {getGameByID} from '@/api/request.js'
import GardenNavHeader from '@/components/headers/GardenNavHeader.vue'

export default {
    name: 'GardenFlowers',
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