<template>
    <div>
        {{error}}
        <div class="market">
            <router-link :to="{path:'/game/garden/market'}">花种交易市场</router-link>
            |
            <router-link :to="{path:'/game/garden/market-onsale'}">我的橱窗</router-link>
        </div>

        【出售花种】<br/>
        <div class="sale-item">
            1.君子兰 (1) 
            <router-link :to="{name:'GardenToolsMarketSale2'}">出售</router-link>
            <br/>
        </div>
        <div class="sale-item">
            2.桂花 (1) 
            <router-link :to="{name:'GardenToolsMarketSale2'}">出售</router-link>
            <br/>
        </div>

        <a href="add.asp?page=2&amp;Sid=zppicxtvR2kc0qgpMIujOy">下页</a><br/>(第<b>1</b>/2页/共6条记录)<br/>
        <form method="post" action="">
            第
            <input name="page" value="1" size="2" format="*n"/>页
            <input name="submit" type="submit" value="前往"/>
        </form>


        ----------<br/>
        <a href="/game/garden/user/my_garden.html">花园</a>><a href="/game/garden/shop/index.asp">市场</a>>出售花种<br/>

    </div>
</template>

<script>
// import {getGameByID} from '@/api/request.js'
// import IndexBaseHeader from '@/components/headers/IndexBaseHeader.vue'

export default {
    name: 'GardenToolsMarketSale1',
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