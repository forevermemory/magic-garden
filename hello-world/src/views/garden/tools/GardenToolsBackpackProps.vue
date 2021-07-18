<template>
    <div>
        {{error}}

        【我的道具】<br/>
花种|<a href="/game/garden/user/my_dz.html?">道具</a><br/>
<div class="list">
<div class="row">
1.君子兰 (1) <a href="/game/garden/user/sowing/35978.html">播种</a><br/>
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
    name: 'GardenToolsBackpackProps',
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