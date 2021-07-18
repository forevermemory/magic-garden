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

【收花记录】<br/>
<div class="list">
<div class="row">
1.<a href="/game/garden/garden.html?gid=35792031"><font color="#ff0000">　暮笙姐姐　</font></a>送红色野花(1朵)，赠言：希望你开心快乐！ [2020-08-20 15:08:33] [<a href="/game/garden/user/gift_flower.html?fid=35792031">再送</a>]<br/>
</div>
<div class="row">
2.<a href="/game/garden/garden.html?gid=35798849">@yet</a>送红色野花(1朵)，赠言：希望你开心快乐！ [2020-08-20 15:07:27] [<a href="/game/garden/user/gift_flower.html?fid=35798849">再送</a>]<br/>
</div>
<div class="row">
</div>
(第<b>1</b>/1页/共2条记录)<br/>
</div>
----------<br/>
<a href="/game/garden/user/my_garden.html">花园</a>><a href="/game/garden/user/basket_list.html">花瓶</a>>收花<br/>


    </div>
</template>

<script>
// import {getGameByID} from '@/api/request.js'
import GardenNavHeader from '@/components/headers/GardenNavHeader.vue'

export default {
    name: 'GardenFlowersReceiveRecord',
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