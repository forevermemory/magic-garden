<template>
    <div>
        <IndexBaseHeader />
        {{error}}

        <router-view></router-view>

    </div>
</template>

<script>
// import {getGameByID} from '@/api/request.js'
import IndexBaseHeader from '@/components/headers/IndexBaseHeader.vue'

export default {
    name: 'GardenIndex',
    data() {
        return {
            // games:[],
            error:'',
            game:{},
        }
    },
    components: {
        IndexBaseHeader,
    },
    created(){
    },
    mounted(){
        
        // 查询游戏列表
        // this.getGamesList()
        // this.$router.push({ name: 'AddGame'})
        // this.$router.push({ path: res.data.data[0].g_url})
        console.log('enter garden manager')
        this.$router.push({path:'/game/garden/index'})
        

    },
    methods: {
        
    }
}
</script>

<style >



/* // <style lang="scss"> */

</style>