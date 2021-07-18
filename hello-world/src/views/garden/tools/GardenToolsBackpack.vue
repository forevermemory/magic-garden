<template>
    <div>
        {{error}}

        【花园背包】<br/>
花种|<a href="/game/garden/user/my_dz.html?">道具</a><br/>
<div class="list">
<div class="row">
1.君子兰 (1) <a href="/game/garden/user/sowing/35978.html">播种</a><br/>
</div>
<div class="row">
2.桂花 (1) <a href="/game/garden/user/sowing/35891.html">播种</a><br/>
</div>
<div class="row">
3.开心小丑 (1) <a href="/game/garden/user/sowing/32089.html">播种</a><br/>
</div>
<div class="row">
4.菊花 (2) <a href="/game/garden/user/sowing/31945.html">播种</a><br/>
</div>
<div class="row">
5.圣诞之吻 (0) <a href="/game/garden/user/sowing/19219.html">播种</a><br/>
</div>
<div class="row">
6.水之精灵 (0) <a href="/game/garden/user/sowing/6685.html">播种</a><br/>
</div>
<div class="row">
7.孤芳自赏 (0) <a href="/game/garden/user/sowing/3226.html">播种</a><br/>
</div>
<div class="row">
8.奥运之花 (0) <a href="/game/garden/user/sowing/3010.html">播种</a><br/>
</div>
<div class="row">
9.梦炫圣诞 (1) <a href="/game/garden/user/sowing/2590.html">播种</a><br/>
</div>
<div class="row">
10.花样年华 (0) <a href="/game/garden/user/sowing/1802.html">播种</a><br/>
</div>
<div class="row">
<a href="/game/garden/user/my_bag-2.html">下页</a>
<form action="/game/garden/user/my_bag.html" method="get">
第<input type="text" name="page" size="2" emptyok="true" format="*N" value="1" maxlength="10" />页
<input type="submit" value="前往" />
</form>
</div>
(第<b>1</b>/2页/共12条记录)<br/>
</div>
----------<br/>
<a href="/game/garden/user/my_garden.html">我的花园</a>>背包<br/>
    </div>
</template>

<script>
// import {getGameByID} from '@/api/request.js'
// import IndexBaseHeader from '@/components/headers/IndexBaseHeader.vue'

export default {
    name: 'GardenToolsBackpack',
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