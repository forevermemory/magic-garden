<template>
    <div>
        {{error}}
        <GardenNavHeader/>
        


<div class="list">
<div class="row">
1.<a href="/game/garden/garden.html?gid=35792031">111</a><br/>
</div>
<div class="row">
2.<a href="/game/garden/garden.html?gid=35798849">倪永炜</a><br/>
</div>
<div class="row">
</div>
(第<b>1</b>/1页/共2条记录)<br/>
</div>
----------<br/>
<a href="/game/garden/user/my_garden.html">花园</a>>好友<br/>


    </div>
</template>

<script>
// import {getGameByID} from '@/api/request.js'
import GardenNavHeader from '@/components/headers/GardenNavHeader.vue'

export default {
    name: 'GardenMyFriend',
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