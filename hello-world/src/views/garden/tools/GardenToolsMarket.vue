<template>
    <div>
        {{error}}

        <div class="market">
            <router-link :to="{path:'/game/garden/market'}">花种交易市场</router-link>
            |
            <router-link :to="{path:'/game/garden/market-onsale'}">我的橱窗</router-link>
        </div>

        花种名称 | 花种价格 | 数量<br/>


1.银湛光宝石花 100元宝 1颗<a href="ok.asp?id=19588&amp;Sid=zppicxtvR2kc0qgpMIujOy">[购买]</a><br/>2.银色幽灵 100元宝 1颗<a href="ok.asp?id=19587&amp;Sid=zppicxtvR2kc0qgpMIujOy">[购买]</a><br/>3.银色烈焰焚情 100元宝 1颗<a href="ok.asp?id=19586&amp;Sid=zppicxtvR2kc0qgpMIujOy">[购买]</a><br/>4.银色岁岁芳华 100元宝 1颗<a href="ok.asp?id=19585&amp;Sid=zppicxtvR2kc0qgpMIujOy">[购买]</a><br/>5.银色亚运风情 100元宝 1颗<a href="ok.asp?id=19584&amp;Sid=zppicxtvR2kc0qgpMIujOy">[购买]</a><br/>6.银色清辉月韵 100元宝 1颗<a href="ok.asp?id=19583&amp;Sid=zppicxtvR2kc0qgpMIujOy">[购买]</a><br/>7.银色国庆牡丹 100元宝 1颗<a href="ok.asp?id=19582&amp;Sid=zppicxtvR2kc0qgpMIujOy">[购买]</a><br/>8.银色寒冰幽兰 100元宝 1颗<a href="ok.asp?id=19581&amp;Sid=zppicxtvR2kc0qgpMIujOy">[购买]</a><br/>9.荣誉银仙女 100元宝 1颗<a href="ok.asp?id=19580&amp;Sid=zppicxtvR2kc0qgpMIujOy">[购买]</a><br/>10.银色祈福-心系雅安 100元宝 1颗<a href="ok.asp?id=19579&amp;Sid=zppicxtvR2kc0qgpMIujOy">[购买]</a><br/><a href="index.asp?page=2&amp;Sid=zppicxtvR2kc0qgpMIujOy">下页</a><br/>(第<b>1</b>/73页/共727条记录)<br/><form method="post" action="index.asp?Sid=zppicxtvR2kc0qgpMIujOy">第<input name="page" value="1" size="2" format="*n"/>页<input name="submit" type="submit" value="前往"/></form>

<router-link :to="{name:'GardenToolsMarketSale1'}">出售闲置花种</router-link>
<br/>
----------<br/>
<a href="/game/garden/user/my_garden.html">花园</a>>市场<br/>


    </div>
</template>

<script>
// import {getGameByID} from '@/api/request.js'
// import IndexBaseHeader from '@/components/headers/IndexBaseHeader.vue'

export default {
    name: 'GardenToolsMarket',
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