<template>
    <div>
        {{error}}
        <GardenNavHeader/>
<div class="list">
<div class="row">
1.<a href="/game/garden/user/room/110.html">银色菊花</a><br/>
</div>
<div class="row">
2.<a href="/game/garden/user/room/111.html">银野花</a><br/>
</div>
<div class="row">
3.<a href="/game/garden/user/room/112.html">银色祈福-心系雅安</a><br/>
</div>
<div class="row">
4.<a href="/game/garden/user/room/113.html">端阳花</a><br/>
</div>
<div class="row">
5.<a href="/game/garden/user/room/114.html">银色烈焰焚情</a><br/>
</div>
<div class="row">
6.<a href="/game/garden/user/room/115.html">金色烈焰焚情</a><br/>
</div>
<div class="row">
7.<a href="/game/garden/user/room/116.html">银色寒冰幽兰</a><br/>
</div>
<div class="row">
8.<a href="/game/garden/user/room/117.html">金色寒冰幽兰</a><br/>
</div>
<div class="row">
9.<a href="/game/garden/user/room/118.html">银友谊花</a><br/>
</div>
<div class="row">
10.<a href="/game/garden/user/room/119.html">荣誉银仙女</a><br/>
</div>
<div class="row">
<a href="/game/garden/user/room_list-2.html">下页</a>
<form action="/game/garden/user/room_list.html" method="get">
第<input type="text" name="page" size="2" emptyok="true" format="*N" value="1" maxlength="10" />页
<input type="submit" value="前往" />
</form>
</div>
(第<b>1</b>/22页/共217条记录)<br/>
</div>
----------<br/>
<a href="/game/garden/user/my_garden.html">花园</a>>魔法屋<br/>
        
    </div>
</template>

<script>
// import {getGameByID} from '@/api/request.js'
import GardenNavHeader from '@/components/headers/GardenNavHeader.vue'

export default {
    name: 'GardenMagicHouse',
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