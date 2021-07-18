<template>
    <div>
        {{error}}

        <GardenAtlasHeader/>
        


<div class="list">
<div class="row">
<img src="/data/game/garden/m_s.gif" alt="."/><a href="/game/garden/user/map/116.html">白色寒冰幽兰</a><br/>
</div>
<div class="row">
<img src="/data/game/garden/m_s.gif" alt="."/><a href="/game/garden/user/map/117.html">红色寒冰幽兰</a><br/>
</div>
<div class="row">
<img src="/data/game/garden/m_s.gif" alt="."/><a href="/game/garden/user/map/118.html">黄色寒冰幽兰</a><br/>
</div>
<div class="row">
<img src="/data/game/garden/m_s.gif" alt="."/><a href="/game/garden/user/map/119.html">粉色寒冰幽兰</a><br/>
</div>
<div class="row">
<img src="/data/game/garden/m_s.gif" alt="."/><a href="/game/garden/user/map/120.html">白色烈焰焚情</a><br/>
</div>
<div class="row">
<img src="/data/game/garden/m_s.gif" alt="."/><a href="/game/garden/user/map/121.html">红色烈焰焚情</a><br/>
</div>
<div class="row">
<img src="/data/game/garden/m_s.gif" alt="."/><a href="/game/garden/user/map/122.html">黄色烈焰焚情</a><br/>
</div>
<div class="row">
<img src="/data/game/garden/m_s.gif" alt="."/><a href="/game/garden/user/map/123.html">粉色烈焰焚情</a><br/>
</div>
<div class="row">
<img src="/data/game/garden/m_s.gif" alt="."/><a href="/game/garden/user/map/124.html">白色清辉月韵</a><br/>
</div>
<div class="row">
<img src="/data/game/garden/m_s.gif" alt="."/><a href="/game/garden/user/map/125.html">红色清辉月韵</a><br/>
</div>
<div class="row">
<a href="/game/garden/user/map_list.html?ty=1&amp;page=2">下页</a>
<form action="/game/garden/user/map_list.html" method="get">
<input type="hidden" name="ty" value="1"/>
第<input type="text" name="page" size="2" emptyok="true" format="*N" value="1" maxlength="10" />页
<input type="submit" value="前往" />
</form>
</div>
(第<b>1</b>/21页/共203条记录)<br/>
</div>
----------<br/>
<a href="/game/garden/user/my_garden.html">花园</a>>花种图谱<br/>
    </div>
</template>

<script>
// import {getGameByID} from '@/api/request.js'
import GardenAtlasHeader from '@/components/headers/GardenAtlasHeader.vue'

export default {
    name: 'GardenAtlasUnique',
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