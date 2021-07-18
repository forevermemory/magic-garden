<template>
    <div>
        {{error}}
【花园商店】<br/>
花种|价格|等级|普通<br/>
<div class="list">
<div class="row">
<img src="/data/game/garden/s_s_1.gif" alt="."/><a href="/game/garden/user/seed/1.html">菊花</a> 1级 200GB <a href="/game/garden/user/seed_buy/1.html">[购买]</a><br/>
</div>
<div class="row">
<img src="/data/game/garden/s_s_2.gif" alt="."/><a href="/game/garden/user/seed/2.html">野花</a> 1级 50GB <a href="/game/garden/user/seed_buy/2.html">[购买]</a><br/>
</div>
<div class="row">
<img src="/data/game/garden/s_s_3.gif" alt="."/><a href="/game/garden/user/seed/3.html">梅花</a> 2级 250GB <a href="/game/garden/user/seed_buy/3.html">[购买]</a><br/>
</div>
<div class="row">
<img src="/data/game/garden/s_s_4.gif" alt="."/><a href="/game/garden/user/seed/4.html">桃花</a> 3级 250GB <a href="/game/garden/user/seed_buy/4.html">[购买]</a><br/>
</div>
<div class="row">
<img src="/data/game/garden/s_s_5.gif" alt="."/><a href="/game/garden/user/seed/5.html">兰花</a> 4级 300GB <a href="/game/garden/user/seed_buy/5.html">[购买]</a><br/>
</div>
<div class="row">
<img src="/data/game/garden/s_s_6.gif" alt="."/><a href="/game/garden/user/seed/6.html">樱花</a> 5级 400GB <a href="/game/garden/user/seed_buy/6.html">[购买]</a><br/>
</div>
<div class="row">
<img src="/data/game/garden/s_s_7.gif" alt="."/><a href="/game/garden/user/seed/7.html">丁香</a> 6级 600GB <a href="/game/garden/user/seed_buy/7.html">[购买]</a><br/>
</div>
<div class="row">
<img src="/data/game/garden/s_s_8.gif" alt="."/><a href="/game/garden/user/seed/8.html">百合</a> 7级 500GB <a href="/game/garden/user/seed_buy/8.html">[购买]</a><br/>
</div>
<div class="row">
<img src="/data/game/garden/s_s_9.gif" alt="."/><a href="/game/garden/user/seed/9.html">玫瑰</a> 8级 700GB <a href="/game/garden/user/seed_buy/9.html">[购买]</a><br/>
</div>
<div class="row">
<img src="/data/game/garden/s_s_10.gif" alt="."/><a href="/game/garden/user/seed/10.html">芙蓉</a> 9级 800GB <a href="/game/garden/user/seed_buy/10.html">[购买]</a><br/>
</div>
<div class="row">
<a href="/game/garden/user/seed_list-2.html">下页</a>
<form action="/game/garden/user/seed_list.html" method="get">
第<input type="text" name="page" size="2" emptyok="true" format="*N" value="1" maxlength="10" />页
<input type="submit" value="前往" />
</form>
</div>
(第<b>1</b>/4页/共36条记录)<br/>
</div>
----------<br/>
<a href="/game/garden/user/my_garden.html">花园</a>>商店<br/>
        
    </div>
</template>

<script>
// import {getGameByID} from '@/api/request.js'
// import IndexBaseHeader from '@/components/headers/IndexBaseHeader.vue'

export default {
    name: 'GardenToolsShop',
    data() {
        return {
            // games:[],
            error:'',
            game:{},
        }
    },
    components: {
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