<template>
    <div>
        {{error}}

        <GardenAtlasHeader/>
        
<div class="list">
<div class="row">
<img src="/data/game/garden/m_s_1.gif" alt="."/><a href="/game/garden/user/map/1.html">白色菊花</a><br/>
</div>
<div class="row">
<img src="/data/game/garden/m_s.gif" alt="."/><a href="/game/garden/user/map/2.html">红色菊花</a><br/>
</div>
<div class="row">
<img src="/data/game/garden/m_s.gif" alt="."/><a href="/game/garden/user/map/3.html">黄色菊花</a><br/>
</div>
<div class="row">
<img src="/data/game/garden/m_s.gif" alt="."/><a href="/game/garden/user/map/4.html">白色野花</a><br/>
</div>
<div class="row">
<img src="/data/game/garden/m_s_5.gif" alt="."/><a href="/game/garden/user/map/5.html">红色野花</a><br/>
</div>
<div class="row">
<img src="/data/game/garden/m_s_6.gif" alt="."/><a href="/game/garden/user/map/6.html">黄色野花</a><br/>
</div>
<div class="row">
<img src="/data/game/garden/m_s_7.gif" alt="."/><a href="/game/garden/user/map/7.html">白梅花</a><br/>
</div>
<div class="row">
<img src="/data/game/garden/m_s_8.gif" alt="."/><a href="/game/garden/user/map/8.html">红梅花</a><br/>
</div>
<div class="row">
<img src="/data/game/garden/m_s_9.gif" alt="."/><a href="/game/garden/user/map/9.html">黄梅花</a><br/>
</div>
<div class="row">
<img src="/data/game/garden/m_s.gif" alt="."/><a href="/game/garden/user/map/10.html">白桃花</a><br/>
</div>
<div class="row">
<a href="/game/garden/user/map_list.html?ty=0&amp;page=2">下页</a>
<form action="/game/garden/user/map_list.html" method="get">
<input type="hidden" name="ty" value="0"/>
第<input type="text" name="page" size="2" emptyok="true" format="*N" value="1" maxlength="10" />页
<input type="submit" value="前往" />
</form>
</div>
(第<b>1</b>/12页/共115条记录)<br/>
</div>
----------<br/>
<a href="/game/garden/user/my_garden.html">花园</a>>花种图谱<br/>
    </div>
</template>

<script>
// import {getGameByID} from '@/api/request.js'
import GardenAtlasHeader from '@/components/headers/GardenAtlasHeader.vue'

export default {
    name: 'GardenAtlasCommon',
    data() {
        return {
            // games:[],
            error:'',
            game:{},
        }
    },
    components: {
        GardenAtlasHeader,
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