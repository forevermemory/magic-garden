<template>
    <div>
        {{error}}【花园商店】<br/>
道具名|价格<br/>
<div class="list">
<div class="row">
1.
星耀光辉（缩短成长14小时） 80元宝<a href="/game/garden/user/dz_buy/95.html">[购买]</a><br/>
</div>
<div class="row">
2.
荣耀圣水（缩短成长9小时30分钟） 55元宝<a href="/game/garden/user/dz_buy/93.html">[购买]</a><br/>
</div>
<div class="row">
3.
闪耀神水（缩短成长7小时30分钟） 45元宝<a href="/game/garden/user/dz_buy/92.html">[购买]</a><br/>
</div>
<div class="row">
4.
染色剂 10元宝<a href="/game/garden/user/dz_buy/87.html">[购买]</a><br/>
</div>
<div class="row">
5.
神圣甘露（缩短成长6小时） 35元宝<a href="/game/garden/user/dz_buy/86.html">[购买]</a><br/>
</div>
<div class="row">
6.
神奇甘露（缩短成长4小时40分钟） 20元宝<a href="/game/garden/user/dz_buy/85.html">[购买]</a><br/>
</div>
<div class="row">
7.
神秘甘露（缩短成长3小时30分钟） 10元宝<a href="/game/garden/user/dz_buy/84.html">[购买]</a><br/>
</div>
<div class="row">
8.
神奇甘露（缩短成长4小时40分钟） 5600GB<a href="/game/garden/user/dz_buy/83.html">[购买]</a><br/>
</div>
<div class="row">
9.
神秘甘露（缩短成长3小时30分钟） 4100GB<a href="/game/garden/user/dz_buy/82.html">[购买]</a><br/>
</div>
<div class="row">
10.
稀有魔力营养液（缩短成长2小时30分钟） 2900GB<a href="/game/garden/user/dz_buy/81.html">[购买]</a><br/>
</div>

<div class="row">
11.
高级魔力营养液（缩短成长1小时40分钟） 1800GB<a href="/game/garden/user/dz_buy/80.html">[购买]</a><br/>
</div>
<div class="row">
12.
大魔力营养液（缩短成长1小时） 900GB<a href="/game/garden/user/dz_buy/79.html">[购买]</a><br/>
</div>
<div class="row">
13.
魔力营养液（缩短成长40分钟）） 550GB<a href="/game/garden/user/dz_buy/78.html">[购买]</a><br/>
</div>
<div class="row">
14.
小魔力营养液（缩短成长20分钟） 250GB<a href="/game/garden/user/dz_buy/77.html">[购买]</a><br/>
</div>
</div>
<div class="tool-box">
<form action="/game/garden/user/dz_list.html" method="get">
<a href="/game/garden/user/dz_list-2.html">下页</a>
第<input type="text" name="page" size="2" emptyok="true" format="*N" value="1" maxlength="10" />页
<input type="submit" value="前往" />
</form>
(第<b>1</b>/2页/共14条记录)<br/>
</div>
----------<br/>
<a href="/game/garden/user/my_garden.html">花园</a>>商店<br/>
        
    </div>
</template>

<script>
// import {getGameByID} from '@/api/request.js'
// import IndexBaseHeader from '@/components/headers/IndexBaseHeader.vue'

export default {
    name: 'GardenToolsProps',
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