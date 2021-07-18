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

        
【送花】<br/>
选择想要送花的好友:<br/>
<div class="list">
<div class="row">
1.<a href="/game/garden/user/gift_flower.html?fid=35792031">111</a><br/>
</div>
<div class="row">
2.<a href="/game/garden/user/gift_flower.html?fid=35798849">倪永炜</a><br/>
</div>
<div class="row">
</div>
(第<b>1</b>/1页/共2条记录)<br/>
</div>
----------<br/>
<form action="/game/garden/user/gift_flower.html" method="get">
好友帐号:<br/>
<input type="text" name="fid" maxlength="10" value="" /><br/>
<input type="submit" value="确定赠送" />
</form>
----------<br/>
    </div>
</template>

<script>
// import {getGameByID} from '@/api/request.js'
import GardenNavHeader from '@/components/headers/GardenNavHeader.vue'

export default {
    name: 'GardenFlowersSend',
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