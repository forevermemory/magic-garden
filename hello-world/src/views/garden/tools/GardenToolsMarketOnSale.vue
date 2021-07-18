<template>
    <div>
        {{error}}
        <div class="market">
            <router-link :to="{path:'/game/garden/market'}">花种交易市场</router-link>
            |
            <router-link :to="{path:'/game/garden/market-onsale'}">我的橱窗</router-link>
        </div>

        <!-- 【我的橱窗】<br/> -->

        花种名称 | 花种价格 | 数量<br/>

        <div class="sale-item">
            1.野花 100元宝 1颗
            <router-link :to="{name:'GardenToolsMarketOnSale2'}">[下架]</router-link>
            <br/>
        </div>


----------<br/>
<a href="/game/garden/user/my_garden.html">花园</a>>市场<br/>

    </div>
</template>

<script>
// import {getGameByID} from '@/api/request.js'
// import IndexBaseHeader from '@/components/headers/IndexBaseHeader.vue'

export default {
    name: 'GardenToolsMarketOnSale',
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