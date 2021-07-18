<template>
    <div>
        {{error}}
        <GardenAtlasHeader/>
        
        
<div class="list">
<div class="row">
<img src="/data/game/garden/m_s.gif" alt="."/><a href="/game/garden/user/map/315.html">银色菊花</a><br/>
</div>
<div class="row">
<img src="/data/game/garden/m_s.gif" alt="."/><a href="/game/garden/user/map/316.html">银野花</a><br/>
</div>
<div class="row">
<img src="/data/game/garden/m_s.gif" alt="."/><a href="/game/garden/user/map/317.html">银色祈福-心系雅</a><br/>
</div>
<div class="row">
<img src="/data/game/garden/m_s.gif" alt="."/><a href="/game/garden/user/map/318.html">端阳花</a><br/>
</div>
<div class="row">
<img src="/data/game/garden/m_s.gif" alt="."/><a href="/game/garden/user/map/319.html">银色烈焰焚情</a><br/>
</div>
<div class="row">
<img src="/data/game/garden/m_s.gif" alt="."/><a href="/game/garden/user/map/320.html">金色烈焰焚情</a><br/>
</div>
<div class="row">
<img src="/data/game/garden/m_s.gif" alt="."/><a href="/game/garden/user/map/321.html">银色寒冰幽兰</a><br/>
</div>
<div class="row">
<img src="/data/game/garden/m_s.gif" alt="."/><a href="/game/garden/user/map/322.html">金色寒冰幽兰</a><br/>
</div>
<div class="row">
<img src="/data/game/garden/m_s.gif" alt="."/><a href="/game/garden/user/map/323.html">银友谊花</a><br/>
</div>
<div class="row">
<img src="/data/game/garden/m_s.gif" alt="."/><a href="/game/garden/user/map/324.html">荣誉银仙女</a><br/>
</div>
<div class="row">
<a href="/game/garden/user/map_list.html?ty=2&amp;page=2">下页</a>
<form action="/game/garden/user/map_list.html" method="get">
<input type="hidden" name="ty" value="2"/>
第<input type="text" name="page" size="2" emptyok="true" format="*N" value="1" maxlength="10" />页
<input type="submit" value="前往" />
</form>
</div>
(第<b>1</b>/22页/共217条记录)<br/>
</div>
----------<br/>
<a href="/game/garden/user/my_garden.html">花园</a>>花种图谱<br/>
    </div>
</template>

<script>
// import {getGameByID} from '@/api/request.js'
import GardenAtlasHeader from '@/components/headers/GardenAtlasHeader.vue'

export default {
    name: 'GardenAtlasRare',
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